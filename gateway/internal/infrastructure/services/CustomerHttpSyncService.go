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

type CustomerHttpSyncService struct {
	BasePath string
}

func (s CustomerHttpSyncService) Register(customer model.CreateCustomerRequest) (uuid.UUID, error) {
	client := &http.Client{}

	data, _ := json.Marshal(customer)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/customers", s.BasePath), bytes.NewBuffer(data))
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return uuid.Nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var id model.CustomerID

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
		return uuid.Nil, model2.CustomerAlreadyExistError{}
	}

	return uuid.Nil, errors.New("unhandled response from service")
}

func (s CustomerHttpSyncService) Update(ID uuid.UUID, name string, surname string, email string) error {
	client := &http.Client{}

	data, _ := json.Marshal(model.UpdateCustomerRequest{
		Name:    name,
		Surname: surname,
		Email:   email,
	})

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/customers/%s", s.BasePath, ID.String()), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode == 404 {
		return model2.CustomerNotFoundError{Id: ID.String()}
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
		return model2.CustomerAlreadyExistError{}
	}

	return nil
}

func (s CustomerHttpSyncService) Delete(ID uuid.UUID) error {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/customers/%s", s.BasePath, ID.String()), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil
	}

	if resp.StatusCode == 404 {
		return model2.CustomerNotFoundError{Id: ID.String()}
	}

	return model2.ServiceUnavailable{}
}

func (s CustomerHttpSyncService) Find(emailPattern string) ([]model.Customer, error) {
	url := fmt.Sprintf("%s/customers?email=%s", s.BasePath, emailPattern)
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

	var arr []model.Customer
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
