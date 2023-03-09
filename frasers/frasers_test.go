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
func (cfg *mockFrasersClientConfig) FrasersPropertyBaseURL() string {
	return os.Getenv("TEST_FRASERS_PROPERTY_BASE_URL")
}
func (cfg *mockFrasersClientConfig) FrasersPropertyGrantTypeClient() string {
	return os.Getenv("TEST_FRASERS_PROPERTY_GRANT_TYPE_CLIENT")
}
func (cfg *mockFrasersClientConfig) FrasersPropertyClientID() string {
	return os.Getenv("TEST_FRASERS_PROPERTY_CLIENT_ID")
}
func (cfg *mockFrasersClientConfig) FrasersPropertyClientSecretKey() string {
	return os.Getenv("TEST_FRASERS_PROPERTY_CLIENT_SECRET_KEY")
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config = &mockFrasersClientConfig{}
	fc = NewFrasersClient(config)
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

func TestAccessToken(t *testing.T) {

	r, err := fc.AccessToken(context.Background())
	if err != nil {
		t.Errorf("AccessToken() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestAccessToken result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "bearer", r.TokenType)
}

func TestCustomer(t *testing.T) {

	cityid := os.Getenv("TEST_FRASERS_PROPERTY_CITIZEN_ID")
	phone := os.Getenv("TEST_FRASERS_PROPERTY_CITIZEN_PHONE")
	consumer := os.Getenv("TEST_FRASERS_PROPERTY_CONSUMER_USERNAME")

	r, err := fc.Customer(context.Background(), consumer, cityid, phone, "or")
	if err != nil {
		t.Errorf("Customer() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestCustomer result =>", string(j))
	fmt.Println("")
	assert.NotEmpty(t, r.Data)
}

func TestProject(t *testing.T) {

	r, err := fc.Project(context.Background())
	if err != nil {
		t.Errorf("Project() failed: %s", err)
	}

	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestProject result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "02021", r.Data[0].ID)
	assert.Equal(t, "โกลเด้น ทาวน์ ๒ ลาดพร้าว-เกษตรนวมินทร์ (ไม่ใช้แล้ว)", r.Data[0].NameData.NameTH)
	assert.Equal(t, "Golden Town ๒ Ladphrao-Kasetnawami (No Use)", r.Data[0].NameData.NameEN)
	assert.Equal(t, "GT๒-LPKN", r.Data[0].NameData.Nickname)
	assert.Equal(t, " ตำบล/แขวง  อำเภอ/เขต  จังหวัด  ", r.Data[0].AddressData.Zone)
	assert.Equal(t, 13.7991430, r.Data[0].AddressData.Latitude)
	assert.Equal(t, 100.6636390, r.Data[0].AddressData.Longtitude)
	assert.Empty(t, r.Data[0].TranferStatus)
}

func TestPlot(t *testing.T) {
	r, err := fc.Plot(context.Background())
	if err != nil {
		t.Errorf("Plot() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestPlot result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "30025", r.Data[0].ProjectData.ID)
	assert.Equal(t, "A01", r.Data[0].PlotData.ID)
	assert.Equal(t, "เซนต์เจมส์", r.Data[0].PlotData.TypeData.TypeTh)
	assert.Equal(t, "Saint James", r.Data[0].PlotData.TypeData.TypeEn)
	assert.Equal(t, 23.00, r.Data[0].PlotData.Sizing)
	assert.Equal(t, "ตรว", r.Data[0].PlotData.Unit)
	assert.Equal(t, "23/39หมู่ -ซอย ท่านผู้หญิงพหลถนน -ตำบล/แขวง ลาดยาวอำเภอ/เขต จตุจักรจังหวัด กรุงเทพมหานคร", r.Data[0].PlotData.Location)
	assert.Equal(t, 3793294.58, r.Data[0].PlotData.Price)
	assert.Equal(t, "T", r.Data[0].PlotData.Status)
}

func TestPlotWarranty(t *testing.T) {

	r, err := fc.PlotWarranty(context.Background())
	if err != nil {
		t.Errorf("TestPlotWarranty() failed: %s", err)
	}
	//Check Result
	j, _ := json.MarshalIndent(r, "", " ")
	fmt.Println("")
	fmt.Println("TestPlotWarranty result =>", string(j))
	fmt.Println("")
	assert.Equal(t, "38026", r.Data[0].ProjectData.ID)
	assert.Equal(t, "00001", r.Data[0].PlotData.ID)
	assert.Equal(t, "2010-08-24T00:00:00Z", r.Data[0].PlotData.TransferDate)
	assert.Equal(t, "2011-08-24T00:00:00Z", r.Data[0].PlotData.WarrantyDate)
	assert.Equal(t, 6339360.78, r.Data[0].PlotData.Price)
}
