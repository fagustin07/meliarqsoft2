package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type SyncHttpDeletePurchasesByProducts struct {
	BasePath string
}

func NewSyncHttpDeletePurchasesByProducts(basePath string) *SyncHttpDeletePurchasesByProducts {
	return &SyncHttpDeletePurchasesByProducts{BasePath: basePath}
}

func (s SyncHttpDeletePurchasesByProducts) Execute(IDs []uuid.UUID) error {
	client := &http.Client{}
	if IDs == nil {
		IDs = []uuid.UUID{}
	}

	data, err := json.Marshal(bson.M{"ids": IDs})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/purchases", s.BasePath), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return model2.ServiceUnavailable{Message: "purchase service currently unavailable"}
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
