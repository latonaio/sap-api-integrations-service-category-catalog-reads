package sap_api_output_formatter

type ServiceCategoryCatalog struct {
	ConnectionKey              string `json:"connection_key"`
	Result                     bool   `json:"result"`
	RedisKey                   string `json:"redis_key"`
	Filepath                   string `json:"filepath"`
	APISchema                  string `json:"api_schema"`
	ServiceCategoryCatalogCode string `json:"service_category_catalog_code"`
	Deleted                    bool   `json:"deleted"`
}

type ServiceCategoryCatalogCollection struct {
	ObjectID                                     string `json:"ObjectID"`
	ETag                                         string `json:"ETag"`
	ID                                           string `json:"ID"`
	VersionID                                    string `json:"VersionID"`
	LifeCycleStatusCode                          string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText                      string `json:"LifeCycleStatusCodeText"`
	EndDateTime                                  string `json:"EndDateTime"`
	StartDateTime                                string `json:"StartDateTime"`
	ServiceCategoryCatalogueName                 string `json:"ServiceCategoryCatalogueName"`
	ServiceCategoryCatalogueNamelanguageCode     string `json:"ServiceCategoryCatalogueNamelanguageCode"`
	ServiceCategoryCatalogueNamelanguageCodeText string `json:"ServiceCategoryCatalogueNamelanguageCodeText"`
	EntityLastChangedOn                          string `json:"EntityLastChangedOn"`
}
