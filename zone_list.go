package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ZoneList struct {
	Success bool       `json:"success"`
	Data    []ZoneData `json:"data"`
}

type ZoneData struct {
	ZoneId int    `json:"zoneId"`
	NameTH string `json:"nameTH"`
	NameEN string `json:"nameEN"`
}

func (c *frasersClient) ZoneList(ctx context.Context) (ZoneList, error) {

	var result ZoneList
	url := fmt.Sprintf("%s%s", c.config.FrasersBaseURL(), "/zonelist")
	method := "POST"

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return ZoneList{}, err
	}

	req.Header.Add("Secret-Key", c.config.FrasersAPIKey())

	res, err := client.Do(req)
	if err != nil {
		return ZoneList{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ZoneList{}, err
	}

	if res.StatusCode != 200 {

		return ZoneList{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return ZoneList{}, err
	}

	return result, nil
}
