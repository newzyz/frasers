package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PlotWarranty struct {
	Data []PlotWarranties `json:"data"`
}
type PlotWarranties struct {
	ProjectData PlotWarrantyProjectData `json:"projectData"`
	PlotData    PlotWarrantyData        `json:"plotData"`
}
type PlotWarrantyProjectData struct {
	ID string `json:"id"`
}
type PlotWarrantyData struct {
	ID           string  `json:"id"`
	TransferDate string  `json:"transferDate"`
	WarrantyDate string  `json:"warrantyDate"`
	Price        float64 `json:"price"`
}

func (c *frasersClient) PlotWarranty(ctx context.Context, projectID string, plotID string) (PlotWarranty, error) {

	var result PlotWarranty
	url := fmt.Sprintf("%s%s%s%s%s%s", c.config.FrasersPropertyBaseURL(), "/v1.0/Projects/", projectID, "/Plots/", plotID, "/PlotWarranty")
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return PlotWarranty{}, err
	}

	//Remark
	r, err := c.AccessToken(context.Background())
	if err != nil {
		return PlotWarranty{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		return PlotWarranty{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return PlotWarranty{}, err
	}

	if res.StatusCode != 200 {

		return PlotWarranty{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return PlotWarranty{}, err
	}

	return result, nil
}
