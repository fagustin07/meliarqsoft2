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

type ProductHttpSyncService struct {
	BasePath string
}

func (p ProductHttpSyncService) Create(createReq model.CreateProductRequest) (model.ProductID, error) {
	client := &http.Client{}

	data, _ := json.Marshal(createReq)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/products", p.BasePath), bytes.NewBuffer(data))
	if err != nil {
		return model.ProductID{}, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return model.ProductID{}, model2.ServiceUnavailable{Message: "product service currently unavailable"}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var newProdId model.ProductID
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.ProductID{}, err
		}

		err = json.Unmarshal(msg, &newProdId)
		if err != nil {
			return model.ProductID{}, err
		}

		return newProdId, nil
	}

	if resp.StatusCode == 400 {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.ProductID{}, err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return model.ProductID{}, err
		}

		return model.ProductID{}, badReqErr
	}

	if resp.StatusCode == 404 {
		return model.ProductID{}, model2.SellerNotFoundError{Id: createReq.IDSeller.String()}
	}

	if resp.StatusCode == 406 {
		var notAcceptableErr model2.NotAcceptableError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.ProductID{}, err
		}

		err = json.Unmarshal(msg, &notAcceptableErr)
		if err != nil {
			return model.ProductID{}, err
		}

		return model.ProductID{}, notAcceptableErr
	}

	if resp.StatusCode == 409 {
		return model.ProductID{}, model2.ProductAlreadyExist{Name: createReq.Name, SellerId: createReq.IDSeller.String()}
	}

	if resp.StatusCode == 503 {
		var unavailableErr model2.ServiceUnavailable
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.ProductID{}, err
		}
		_ = json.Unmarshal(msg, &unavailableErr)

		if err != nil {
			return model.ProductID{}, err
		}
		return model.ProductID{}, unavailableErr
	}
	
	if resp.StatusCode >= 500 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.ProductID{}, err
		}

		return model.ProductID{}, errors.New(string(msg))
	}

	return model.ProductID{}, nil
}

func (p ProductHttpSyncService) Update(ID uuid.UUID, prodReq model.UpdateProductRequest) error {
	client := &http.Client{}

	data, _ := json.Marshal(prodReq)

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/products/%s", p.BasePath, ID.String()), bytes.NewBuffer(data))
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
	url := fmt.Sprintf("%s/products?name=%s&category=%s", p.BasePath, name, category)
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
	url := fmt.Sprintf("%s/products/prices?min_price=%f&max_price=%f", p.BasePath, minPrice, maxPrice)
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

func (p ProductHttpSyncService) Delete(ID uuid.UUID) error {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/products/%s", p.BasePath, ID.String()), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return model2.ServiceUnavailable{Message: "products service currently unavailable"}
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
			return model2.ProductNotFound{Id: ID.String()}
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
			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return errors.New(string(msg))
		}
	}

	return model2.ServiceUnavailable{}
}
