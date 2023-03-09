package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CustomerResponse struct {
	Data         string `json:"data"`
	EncryptedKey string `json:"encryptedKey"`
}

func (c *frasersClient) Customer(ctx context.Context, consumerUsername, citizenId, phoneNumber, ctype string) (CustomerResponse, error) {

	var result CustomerResponse
	urlStr := fmt.Sprintf("%s%s", c.config.FrasersPropertyBaseURL(), "/v1.0/Customers/Customer")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, urlStr, nil)
	if err != nil {
		return CustomerResponse{}, err
	}

	//Remark
	r, err := c.AccessToken(context.Background())
	if err != nil {
		return CustomerResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))
	req.Header.Set("X-Consumer-Username", consumerUsername)
	req.Header.Set("CitizenId", citizenId)
	req.Header.Set("PhoneNumber", phoneNumber)
	req.Header.Set("Type", ctype)

	res, err := client.Do(req)
	if err != nil {
		return CustomerResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CustomerResponse{}, err
	}

	if res.StatusCode != 200 {

		return CustomerResponse{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return CustomerResponse{}, err
	}

	return result, nil
}
