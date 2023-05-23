package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type SyncHttpFindSellerById struct {
	BasePath string
}

func NewSyncHttpFindSellerById(basePath string) *SyncHttpFindSellerById {
	return &SyncHttpFindSellerById{BasePath: basePath}
}

func (service SyncHttpFindSellerById) Execute(id uuid.UUID) (model.SellerDTO, error) {
	url := fmt.Sprintf("%s/sellers/%s", service.BasePath, id)
	resp, err := http.Get(url)

	if err != nil {
		return model.SellerDTO{}, model2.ServiceUnavailable{Message: "seller service currently unavailable"}
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var seller model.SellerDTO
		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.SellerDTO{}, err
		}

		err = json.Unmarshal(bodyResp, &seller)
		if err != nil {
			return model.SellerDTO{}, err
		}

		return seller, nil

	}

	if resp.StatusCode == 400 {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.SellerDTO{}, err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return model.SellerDTO{}, err
		}

		return model.SellerDTO{}, badReqErr
	}

	if resp.StatusCode == 404 {
		return model.SellerDTO{}, model2.SellerNotFoundError{Id: id.String()}
	}

	if resp.StatusCode >= 500 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.SellerDTO{}, err
		}

		return model.SellerDTO{}, errors.New(string(msg))
	}

	return model.SellerDTO{}, errors.New("unhandled status code")
}
