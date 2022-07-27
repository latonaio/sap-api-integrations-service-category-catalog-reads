package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-service-category-catalog-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToServiceCategoryCatalogCollection(raw []byte, l *logger.Logger) ([]ServiceCategoryCatalogCollection, error) {
	pm := &responses.ServiceCategoryCatalogCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ServiceCategoryCatalogCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	serviceCategoryCatalogCollection := make([]ServiceCategoryCatalogCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		serviceCategoryCatalogCollection = append(serviceCategoryCatalogCollection, ServiceCategoryCatalogCollection{
			ObjectID:                                 			data.ObjectID,
			ETag:                                     			data.ETag,
			ID:                                       			data.ID,
			VersionID:                                			data.VersionID,
			LifeCycleStatusCode:                      			data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:                  			data.LifeCycleStatusCodeText,
			EndDateTime:                              			data.EndDateTime,
			StartDateTime:                            			data.StartDateTime,
			ServiceCategoryCatalogueName:             			data.ServiceCategoryCatalogueName,
			ServiceCategoryCatalogueNamelanguageCode: 			data.ServiceCategoryCatalogueNamelanguageCode,
			ServiceCategoryCatalogueNamelanguageCodeText:       data.ServiceCategoryCatalogueNamelanguageCodeText,
			EntityLastChangedOn:                                data.EntityLastChangedOn,
		})
	}

	return serviceCategoryCatalogCollection, nil
}