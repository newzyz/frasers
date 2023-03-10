package nocnoc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetCategoriesDetail struct {
	Data Categories `json:"data"`
}

type Categories struct {
	ObjectID    string       `json:"objectId"`
	Title       string       `json:"title"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	ImageURL    *string      `json:"imageUrl,omitempty"`
	WebLinkURL  *string      `json:"webLinkUrl,omitempty"`
	IconURL     *string      `json:"iconUrl,omitempty"`
	IsHighlight *bool        `json:"isHighlight,omitempty"`
	ActionType  *string      `json:"actionType,omitempty"`
	LinkURL     *string      `json:"linkUrl,omitempty"`
	Categories  []Categories `json:"categories,omitempty"`
}

func (nn *nocNocClient) GetCategories(ctx context.Context, categoryId string, deepLevel int) (GetCategoriesDetail, error) {

	var result GetCategoriesDetail

	queryParams := fmt.Sprintf("?deepLevel=%d", deepLevel)

	url := fmt.Sprintf("%s%s%s%s", nn.config.NocNocBaseURL(), "/installer/v1/client/categories/", categoryId, queryParams)
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return result, err
	}

	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if res.StatusCode != 200 {

		return result, fmt.Errorf("nocnoc error : %s", string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
