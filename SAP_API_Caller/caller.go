package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-service-category-catalog-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetServiceCategoryCatalog(iD, versionID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ServiceCategoryCatalogCollection":
			func() {
				c.ServiceCategoryCatalogCollection(iD, versionID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ServiceCategoryCatalogCollection(iD, versionID string) {
	data, err := c.callServiceCategoryCatalogSrvAPIRequirementServiceCategoryCatalogCollection("ServiceCategoryCatalogueCollection", iD, versionID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callServiceCategoryCatalogSrvAPIRequirementServiceCategoryCatalogCollection(api, iD, versionID string) ([]sap_api_output_formatter.ServiceCategoryCatalogCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithServiceCategoryCatalogCollection(req, iD, versionID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToServiceCategoryCatalogCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithServiceCategoryCatalogCollection(req *http.Request, iD, versionID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s' and VersionID eq '%s'", iD, versionID))
	req.URL.RawQuery = params.Encode()
}
