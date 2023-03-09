package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Plot struct {
	Data []PlotDatas `json:"data"`
}
type PlotDatas struct {
	ProjectData PlotProjectData `json:"projectData"`
	PlotData    PlotData        `json:"plotData"`
}
type PlotProjectData struct {
	ID string `json:"id"`
}
type PlotData struct {
	ID       string   `json:"id"`
	TypeData TypeData `json:"typeData"`
	Sizing   float64  `json:"sizing"`
	Unit     string   `json:"unit"`
	UnitType string   `json:"unitType"`
	Location string   `json:"location"`
	Price    float64  `json:"price"`
	Status   string   `json:"status"`
}
type TypeData struct {
	TypeTh string `json:"typeTh"`
	TypeEn string `json:"typeEn"`
}

func (c *frasersClient) Plot(ctx context.Context) (Plot, error) {

	var result Plot
	url := fmt.Sprintf("%s%s", c.config.FrasersPropertyBaseURL(), "/v1.0/Projects/30025/Plots/All/UpdateDate/2010-01-01")
	method := "GET"

	fmt.Println("URL", url)
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return Plot{}, err
	}

	//Remark
	r, err := c.AccessToken(context.Background())
	if err != nil {
		return Plot{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		return Plot{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Plot{}, err
	}

	if res.StatusCode != 200 {

		return Plot{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return Plot{}, err
	}

	return result, nil
}
