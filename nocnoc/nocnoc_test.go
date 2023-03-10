package nocnoc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	nn     NocNocClient
	config NocNocClientConfig
)

type mockNocNocClientConfig struct{}

func (cfg *mockNocNocClientConfig) NocNocBaseURL() string {
	return os.Getenv("TEST_NOCNOC_BASE_URL")
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config = &mockNocNocClientConfig{}
	nn = NewNocNocClient(config)
}

func TestSettingsCategories(t *testing.T) {

	isHighlight := true
	isRecommend := true
	// included := "categories,categories"

	r, err := nn.SettingsCategories(context.Background(), &isHighlight, &isRecommend, nil)
	if err != nil {
		t.Errorf("SettingsCategories() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestSettingsCategoriesresult =>", string(j))
	fmt.Println("")
	assert.Equal(t, "CAT104", *r.Data[0].ObjectId)
}

func TestDeleteSettingsCategories(t *testing.T) {

	err := nn.DeleteSettingsCategories(context.Background(), "CAT66")
	if err != nil {
		t.Errorf("DeleteSettingsCategories() failed: %s", err)
	}

}

func TestAddSettingsCategories(t *testing.T) {

	r, err := nn.AddSettingsCategories(context.Background(), "CAT104", true, true)
	if err != nil {
		t.Errorf("AddSettingsCategories() failed: %s", err)
	}

	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestAddSettingsCategoriesresult =>", string(j))
	fmt.Println("")
	assert.Equal(t, "CAT104", *r.Data.ObjectId)
}

func TestGetAllCategories(t *testing.T) {
	startLevel := 0
	downLevel := 2
	var isAvailable *bool = nil
	var isHighlight *bool = nil
	var isRecommend *bool = nil
	query := &CategoriesQuery{
		StartLevel:  &startLevel,
		DownLevel:   &downLevel,
		IsAvailable: isAvailable,
		IsHighlight: isHighlight,
		IsRecommend: isRecommend,
	}

	r, err := nn.GetAllCategories(context.Background(), query)
	if err != nil {
		t.Errorf("TestGetAllCategories() failed: %s", err)
	}

	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	
	fmt.Println("")
	fmt.Println("TestGetAllCategoriesResult =>", string(j))
	fmt.Println("")
	assert.Equal(t, 10, r.Meta.TotalCount)
	assert.Equal(t, "CAT104", r.Data[0].ObjectID)
	assert.Equal(t, "บริการแนะนำ", r.Data[0].Title)
	assert.Equal(t, "https://d1led8yawtaysu.cloudfront.net/assets/icons/Thumb-up-line.svg", r.Data[0].IconURL)
	assert.Equal(t, "2023-01-11T07:50:54.068Z", r.Data[0].CreatedAt)
	assert.Equal(t, "2023-03-09T06:32:14.000Z", r.Data[0].UpdatedAt)
}
