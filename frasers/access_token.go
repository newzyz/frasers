package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AccessTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *frasersClient) AccessToken(ctx context.Context) (AccessTokenResponse, error) {

	var result AccessTokenResponse
	urlStr := fmt.Sprintf("%s%s", c.config.FrasersPropertyBaseURL(), "/oauth2/token")
	method := "POST"

	form := url.Values{}
	form.Add("grant_type", c.config.FrasersPropertyGrantTypeClient())
	form.Add("client_id", c.config.FrasersPropertyClientID())
	form.Add("client_secret", c.config.FrasersPropertyClientSecretKey())

	payload := strings.NewReader(form.Encode())

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, urlStr, payload)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	if res.StatusCode != 200 {

		return AccessTokenResponse{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return AccessTokenResponse{}, err
	}

	return result, nil
}
