package nocnoc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SettingsCategoriesQuery struct {
	IsHighlight *bool   `json:"isHighlight"`
	IsRecommend *bool   `json:"isRecommend"`
	Includes    *string `json:"includes,omitempty"`
}

type SettingsCategoriesList struct {
	Meta     Meta                `json:"meta"`
	Data     []CategoriesSetting `json:"data"`
	Included *Included           `json:"included"`
}

type Meta struct {
	TotalCount *int `json:"totalCount"`
}

type Included struct {
	Categories []Categories `json:"categories"`
}

type CategoriesSetting struct {
	ObjectId    string `json:"objectId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	IsHighlight *bool  `json:"isHighlight,omitempty"`
	IsRecommend *bool  `json:"isRecommend,omitempty"`
}

func (nn *nocNocClient) SettingsCategories(ctx context.Context, isHighlight, isRecommend *bool, includes *string) (SettingsCategoriesList, error) {

	var result SettingsCategoriesList

	queryParams := "?"
	if isHighlight != nil {
		queryParams += fmt.Sprintf("isHighlight=%t&", *isHighlight)
	}
	if isRecommend != nil {
		queryParams += fmt.Sprintf("isRecommend=%t&", *isRecommend)
	}
	if includes != nil {
		queryParams += fmt.Sprintf("includes=%s&", *includes)
	}

	url := fmt.Sprintf("%s%s%s", nn.config.NocNocBaseURL(), "/installer/v1/admin/settings/categories", queryParams)
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return SettingsCategoriesList{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return SettingsCategoriesList{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SettingsCategoriesList{}, err
	}

	if res.StatusCode != 200 {

		return SettingsCategoriesList{}, fmt.Errorf("nocnoc error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return SettingsCategoriesList{}, err
	}

	return result, nil
}
