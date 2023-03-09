package nocnoc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (nn *nocNocClient) DeleteSettingsCategories(ctx context.Context, categoryId string) error {

	url := fmt.Sprintf("%s%s%s", nn.config.NocNocBaseURL(), "/installer/v1/admin/settings/categories/", categoryId)
	method := "DELETE"

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {

		return fmt.Errorf("nocnoc error : %s", string(body))
	}

	return nil
}
