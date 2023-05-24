package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type MakePurchaseHttpSyncService struct {
	BasePath string
}

func (p MakePurchaseHttpSyncService) Execute(purchase model.CreatePurchaseRequest) (model.Purchase, error) {
	client := &http.Client{}

	data, _ := json.Marshal(purchase)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/create", p.BasePath), bytes.NewBuffer(data))
	log.Println(p.BasePath)
	if err != nil {
		return model.Purchase{}, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return model.Purchase{}, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		var badReqErr model2.BadRequestError
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.Purchase{}, err
		}

		err = json.Unmarshal(msg, &badReqErr)
		if err != nil {
			return model.Purchase{}, err
		}

		return model.Purchase{}, badReqErr
	}

	defer resp.Body.Close()

	var newPurchase model.Purchase
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Purchase{}, err
	}

	err = json.Unmarshal(all, &newPurchase)
	if err != nil {
		return model.Purchase{}, err
	}

	return newPurchase, nil
}
