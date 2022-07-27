package responses

type ServiceCategoryCatalogCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
