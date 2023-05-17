package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type SellerHttpSyncService struct {
	BasePath string
}

func (s SellerHttpSyncService) Register(seller *model.Seller) (uuid.UUID, error) {
	client := &http.Client{}

	data, err := json.Marshal(dto.CreateSellerRequest{Email: seller.Email.Address, BusinessName: seller.BusinessName})
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
		var id dto.SellerID

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

	if resp.StatusCode == 409 {
		return uuid.Nil, model2.SellerAlreadyExist{}
	}

	return uuid.Nil, errors.New("unhandled response from service")
}

func (s SellerHttpSyncService) Update(id uuid.UUID, businessName string, email string) error {
	client := &http.Client{}

	data, _ := json.Marshal(dto.UpdateSellerRequest{
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

	if resp.StatusCode == 409 {
		return model2.SellerAlreadyExist{}
	}

	return nil
}

func (s SellerHttpSyncService) Find(businessName string) ([]model.SellerJSON, error) {
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

	var arr []model.SellerJSON
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
