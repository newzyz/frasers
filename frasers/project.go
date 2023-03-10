package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	Data []ProjectData2 `json:"data"`
}
type ProjectData2 struct {
	ID            string      `json:"id"`
	NameData      NameData    `json:"nameData"`
	AddressData   AddressData `json:"addressData"`
	TranferStatus string      `json:"tranferStatus"`
}
type NameData struct {
	NameTH   string `json:"nameTH"`
	NameEN   string `json:"nameEN"`
	Nickname string `json:"nickname"`
}
type AddressData struct {
	Zone       string  `json:"zone"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}

func (c *frasersClient) Project(ctx context.Context, projectID string, updateDate string) (Project, error) {

	var result Project
	url := fmt.Sprintf("%s%s%s%s%s", c.config.FrasersPropertyBaseURL(), "/v1.0/Projects/", projectID, "/UpdateDate/", updateDate)
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return Project{}, err
	}

	//Remark
	r, err := c.AccessToken(context.Background())
	if err != nil {
		return Project{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		return Project{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Project{}, err
	}

	if res.StatusCode != 200 {

		return Project{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return Project{}, err
	}

	return result, nil
}
