package nocnoc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CategoriesQuery struct {
	StartLevel  *int  `json:"startLevel"`
	DownLevel   *int  `json:"downLevel"`
	IsAvailable *bool `json:"isAvailable"`
	IsHighlight *bool `json:"isHighlight"`
	IsRecommend *bool `json:"isRecommend"`
}

func (query *CategoriesQuery) fill() string {
	queryParams := "?"
	if query.StartLevel != nil {
		queryParams += fmt.Sprintf("startLevel=%d&", *query.StartLevel)
	}
	if query.DownLevel != nil {
		queryParams += fmt.Sprintf("downLevel=%d&", *query.DownLevel)
	}
	if query.IsAvailable != nil {
		queryParams += fmt.Sprintf("isAvailable=%t&", *query.IsAvailable)
	}
	if query.IsHighlight != nil {
		queryParams += fmt.Sprintf("isHighlight=%t&", *query.IsHighlight)
	}
	if query.IsRecommend != nil {
		queryParams += fmt.Sprintf("isRecommend=%t&", *query.IsRecommend)
	}
	return queryParams
}

type CategoriesResponse struct {
	Meta struct {
		TotalCount int `json:"totalCount"`
	} `json:"meta"`
	Data []struct {
		ObjectID   string `json:"objectId"`
		Title      string `json:"title"`
		IconURL    string `json:"iconUrl"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		Categories []struct {
			ObjectID    string `json:"objectId"`
			Title       string `json:"title"`
			ImageURL    string `json:"imageUrl"`
			IsHighlight bool   `json:"isHighlight"`
			CreatedAt   string `json:"createdAt"`
			UpdatedAt   string `json:"updatedAt"`
			Categories  []struct {
				ObjectID   string `json:"objectId"`
				Title      string `json:"title"`
				ImageURL   string `json:"imageUrl"`
				CreatedAt  string `json:"createdAt"`
				UpdatedAt  string `json:"updatedAt"`
				WebLinkURL string `json:"webLinkUrl"`
			} `json:"categories,omitempty"`
			WebLinkURL string `json:"webLinkUrl"`
		} `json:"categories,omitempty"`
		ImageURL string `json:"imageUrl"`
	} `json:"data,omitempty"`
}

func (nn *nocNocClient) GetAllCategories(ctx context.Context, query *CategoriesQuery) (CategoriesResponse, error) {

	var result CategoriesResponse
	queryParams := query.fill()
	url := fmt.Sprintf("%s%s%s", nn.config.NocNocBaseURL(), "/installer/v1/client/categories", queryParams)
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return CategoriesResponse{}, err
	}

	req.Header.Set("Accept-Language", fmt.Sprintf("th, en;q=0.9, *;q=0.3"))
	req.Header.Set("Accept", fmt.Sprintf("application/json"))

	res, err := client.Do(req)
	if err != nil {
		return CategoriesResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CategoriesResponse{}, err
	}

	if res.StatusCode != 200 {

		return CategoriesResponse{}, fmt.Errorf("nocnoc error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return CategoriesResponse{}, err
	}

	return result, nil
}
