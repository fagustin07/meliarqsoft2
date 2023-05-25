package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type HttpSyncDeleteProductsBySeller struct {
	BasePath string
}

func NewHttpSyncDeleteProductsBySeller(basePath string) *HttpSyncDeleteProductsBySeller {
	return &HttpSyncDeleteProductsBySeller{BasePath: basePath}
}

func (s HttpSyncDeleteProductsBySeller) Execute(ID uuid.UUID) error {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/products/sellers/%s", s.BasePath, ID.String()), nil)
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

		if resp.StatusCode == 503 {
			var unavailableErr model2.ServiceUnavailable

			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			_ = json.Unmarshal(msg, &unavailableErr)

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

	return nil
}
