package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type ProductHttpSyncService struct {
	BasePath string
}

func (p ProductHttpSyncService) Update(ID uuid.UUID, prodReq model.UpdateProductRequest) error {
	client := &http.Client{}

	data, _ := json.Marshal(prodReq)

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", p.BasePath, ID.String()), bytes.NewBuffer(data))
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
		return model2.ProductNotFound{Id: ID.String()}
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
		return model2.ProductAlreadyExist{}
	}

	return nil
}

func (p ProductHttpSyncService) Find(name string, category string) ([]model.Product, error) {
	url := fmt.Sprintf("%s?name=%s&category=%s", p.BasePath, name, category)
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

	var arr []model.Product
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

func (p ProductHttpSyncService) Filter(minPrice float32, maxPrice float32) ([]model.Product, error) {
	url := fmt.Sprintf("%s/prices?min_price=%f&max_price=%f", p.BasePath, minPrice, maxPrice)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 500 {
		return nil, model2.ServiceUnavailable{}
	}

	if resp.StatusCode == 406 {
		var notAcceptableErr model2.NotAcceptableError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(msg, &notAcceptableErr)
		if err != nil {
			return nil, err
		}

		return nil, notAcceptableErr
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

	var arr []model.Product
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
