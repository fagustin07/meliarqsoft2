package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type SellerHttpSyncService struct {
	BasePath string
}

func (s SellerHttpSyncService) Register(seller model.CreateSellerRequest) (uuid.UUID, error) {
	client := &http.Client{}

	data, err := json.Marshal(seller)
	if err != nil {
		return uuid.Nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/sellers", s.BasePath), bytes.NewBuffer(data))
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := client.Do(request)
	if err != nil || resp.StatusCode >= 500 {
		return uuid.Nil, model2.ServiceUnavailable{}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var id model.SellerID

		uuidResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return uuid.Nil, err
		}

		err = json.Unmarshal(uuidResp, &id)
		if err != nil {
			return uuid.Nil, err
		}

		return id.ID, err
	}

	if resp.StatusCode == 400 {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return uuid.Nil, err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return uuid.Nil, err
		}

		return uuid.Nil, badReqErr
	}

	if resp.StatusCode == 406 {
		var notAcceptableErr model2.NotAcceptableError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return uuid.Nil, err
		}

		err = json.Unmarshal(msg, &notAcceptableErr)
		if err != nil {
			return uuid.Nil, err
		}

		return uuid.Nil, notAcceptableErr
	}

	if resp.StatusCode == 409 {
		return uuid.Nil, model2.SellerAlreadyExist{}
	}

	return uuid.Nil, errors.New("unhandled response from service")
}

func (s SellerHttpSyncService) Update(id uuid.UUID, businessName string, email string) error {
	client := &http.Client{}

	data, _ := json.Marshal(model.UpdateSellerRequest{
		BusinessName: businessName,
		Email:        email,
	})

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/sellers/%s", s.BasePath, id.String()), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil || resp.StatusCode >= 500 {
		return model2.ServiceUnavailable{}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode == 404 {
		return model2.SellerNotFoundError{Id: id.String()}
	}

	if resp.StatusCode == 400 {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return err
		}

		return badReqErr
	}

	if resp.StatusCode == 406 {
		var notAcceptableErr model2.NotAcceptableError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(msg, &notAcceptableErr)
		if err != nil {
			return err
		}

		return notAcceptableErr
	}

	if resp.StatusCode == 409 {
		return model2.SellerAlreadyExist{}
	}

	return nil
}

func (s SellerHttpSyncService) Find(businessName string) ([]model.Seller, error) {
	url := fmt.Sprintf("%s/sellers?business_name=%s", s.BasePath, businessName)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 500 {
		return nil, model2.ServiceUnavailable{}
	}

	if resp.StatusCode == 400 {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return nil, err
		}

		return nil, badReqErr
	}

	defer resp.Body.Close()

	var arr []model.Seller
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(all, &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

func (s SellerHttpSyncService) Delete(ID uuid.UUID) error {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/sellers/%s", s.BasePath, ID.String()), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return model2.ServiceUnavailable{Message: "seller service currently unavailable"}
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil
	} else {
		var badReqErr model2.BadRequestError

		if resp.StatusCode == 400 {
			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			_ = json.Unmarshal(msg, &badReqErr)
			return badReqErr
		}

		if resp.StatusCode == 404 {
			return model2.SellerNotFoundError{Id: ID.String()}
		}

		if resp.StatusCode == 503 {
			var unavailableErr model2.ServiceUnavailable
			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			_ = json.Unmarshal(msg, &unavailableErr)

			if err != nil {
				return err
			}
			return unavailableErr
		}

		if resp.StatusCode >= 500 {
			if err != nil {
				return err
			}
			return errors.New("internal Server Error")
		}
	}

	return model2.ServiceUnavailable{}
}
