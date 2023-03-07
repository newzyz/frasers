package frasers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProjectListQuery struct {
	PageNo        *int     `json:"pageNo"`
	PageSize      *int     `json:"pageSize"`
	SearchQuery   *string  `json:"searchQuery"`
	ProjectType   *int     `json:"projectType"`
	BrandId       *int     `json:"brandId"`
	SortType      *int     `json:"sortType"`
	PriceMin      *float64 `json:"priceMin"`
	PriceMax      *float64 `json:"priceMax"`
	RecommendFlag *int     `json:"recommendFlag"`
	Lat           *float64 `json:"lat"`
	Long          *float64 `json:"long"`
}

type ProjectList struct {
	Success   bool          `json:"success"`
	PageNo    int           `json:"pageNo"`
	PageSize  int           `json:"pageSize"`
	PageCount int           `json:"pageCount"`
	Data      []ProjectData `json:"data"`
}

type ProjectData struct {
	ProjectID       int     `json:"projectId"`
	NameTH          string  `json:"nameTH"`
	NameEN          string  `json:"nameEN"`
	PriceStart      float64 `json:"priceStart"`
	PriceStartText  string  `json:"priceStartText"`
	CoverPhoto      string  `json:"coverPhoto"`
	RecommendFlag   int     `json:"recommendFlag"`
	ReadyToMoveFlag int     `json:"readyToMoveFlag"`
	URLDetailTH     string  `json:"urlDetailTH"`
	URLDetailEN     string  `json:"urlDetailEN"`
}

func (c *frasersClient) ProjectList(ctx context.Context, plq ProjectListQuery) (ProjectList, error) {

	var result ProjectList
	url := fmt.Sprintf("%s%s", c.config.FrasersBaseURL(), "/projectlist")
	method := "POST"

	jsonBytes, err := json.Marshal(plq)
	if err != nil {
		return ProjectList{}, err
	}
	jsonStr := string(jsonBytes)

	// payload := strings.NewReader(fmt.Sprintf(`{
	//     "pageNo":%d,
	// 	"pageSize":%d,
	// 	"searchQuery":%s,
	// 	"projectType":%d,
	// 	"brandId":%d,
	// 	"sortType":%d,
	// 	"priceMin":%f,
	// 	"priceMax":%f,
	// 	"recommendFlag":%d,
	// 	"lat":%f,
	// 	"long":%f
	// }`, *plq.PageNo, *plq.PageSize, *plq.SearchQuery, *plq.ProjectType,
	// 	plq.BrandId, *plq.SortType, *plq.PriceMin,
	// 	*plq.PriceMax, *plq.RecommendFlag, *plq.Lat, *plq.Long))

	payload := strings.NewReader(jsonStr)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return ProjectList{}, err
	}
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Secret-Key", c.config.FrasersAPIKey())

	res, err := client.Do(req)
	if err != nil {
		return ProjectList{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ProjectList{}, err
	}

	if res.StatusCode != 200 {

		return ProjectList{}, fmt.Errorf("client error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return ProjectList{}, err
	}

	return result, nil
}
