package frasers

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
	fc     FrasersClient
	config FrasersClientConfig
)

type mockFrasersClientConfig struct{}

func (cfg *mockFrasersClientConfig) FrasersBaseURL() string {
	return os.Getenv("TEST_FRASERS_BASE_URL")
}
func (cfg *mockFrasersClientConfig) FrasersAPIKey() string {
	return os.Getenv("TEST_FRASERS_API_KEY")
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config = &mockFrasersClientConfig{}
	fc = NewFrasersClient(config)
	fmt.Println(os.Getenv("TEST_FRASERS_BASE_URL"))
}

func TestProjectList(t *testing.T) {

	pageN := 1
	pageS := 10
	NameTH := "เพรสทีจ ฟิวเจอร์-รังสิต"
	query := ProjectListQuery{
		PageNo:        &pageN,
		PageSize:      &pageS,
		SearchQuery:   &NameTH,
		ProjectType:   nil,
		PriceMin:      nil,
		PriceMax:      nil,
		BrandId:       nil,
		SortType:      nil,
		RecommendFlag: nil,
		Lat:           nil,
		Long:          nil,
	}
	r, err := fc.ProjectList(context.Background(), query)
	if err != nil {
		t.Errorf("ProjectList() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestProjectList result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "Prestige Future-Rangsit", r.Data[0].NameEN)
}

func TestZoneList(t *testing.T) {

	r, err := fc.ZoneList(context.Background())
	if err != nil {
		t.Errorf("ZoneList() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestZoneList result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "งามวงศ์วาน-แคราย-วงศ์สว่าง-ประชาชื่น", r.Data[0].NameTH)
}
