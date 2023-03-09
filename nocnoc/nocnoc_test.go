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
