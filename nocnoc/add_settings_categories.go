package nocnoc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AddSettingsCategoriesResponse struct {
	Data AddSettingsCategoriesData `json:"data"`
}

type AddSettingsCategoriesData struct {
	ObjectId    *string `json:"objectId"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	IsHighlight *bool   `json:"isHighlight"`
	IsRecommend *bool   `json:"isRecommend"`
}

func (nn *nocNocClient) AddSettingsCategories(ctx context.Context, categoryId string, isHighlight, isRecommend bool) (AddSettingsCategoriesResponse, error) {

	var result AddSettingsCategoriesResponse

	url := fmt.Sprintf("%s%s%s", nn.config.NocNocBaseURL(), "/installer/v1/admin/settings/categories/", categoryId)
	method := "PATCH"

	payload := strings.NewReader(fmt.Sprintf(`{
        "isHighlight":%t,
        "isRecommend":%t
    }`, isHighlight, isRecommend))

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return AddSettingsCategoriesResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return AddSettingsCategoriesResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return AddSettingsCategoriesResponse{}, err
	}

	if res.StatusCode != 200 {

		return AddSettingsCategoriesResponse{}, fmt.Errorf("nocnoc error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return AddSettingsCategoriesResponse{}, err
	}

	return result, nil
}
