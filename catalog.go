package tiktokads

import (
	"errors"
	"time"
)

type catalogConfig struct {
	RegionCode string `json:"region_code"`
	Currency   string `json:"currency"`
}

// CatalogCreateRequest Only for creation
type CatalogCreateRequest struct {
	Name             string         `json:"name,omitempty"`
	CatalogType      string         `json:"catalog_type,omitempty"`
	BusinessCenterId string         `json:"bc_id,omitempty"`
	Conf             *catalogConfig `json:"catalog_conf,omitempty"`
}

type Catalog struct {
	Id   string         `json:"catalog_id,omitempty"`
	Name string         `json:"catalog_name,omitempty"`
	Conf *catalogConfig `json:"catalog_conf,omitempty"`

	CatalogType        string `json:"catalog_type,omitempty"`
	AdCreationEligible string `json:"ad_creation_eligible,omitempty"`

	Business *BusinessCenter `json:"bc_info,omitempty"`
}

func (c *Catalog) IsAdCreationEligible() bool {
	return "AVAILABLE" == c.AdCreationEligible
}

func (c *Catalog) IsEcommerce() bool {
	return c.CatalogType == "ECOM"
}

func NewShoppingCatalog(businessId, name, currency, country string) *CatalogCreateRequest {
	return &CatalogCreateRequest{
		BusinessCenterId: businessId,
		Name:             name,
		CatalogType:      "ECOM",
		Conf: &catalogConfig{
			RegionCode: country,
			Currency:   currency,
		},
	}
}

type catalogCreateResult struct {
	CatalogId string `json:"catalog_id"`
}

func CreateCatalog(catalog *CatalogCreateRequest) (*Catalog, error) {
	req := newPostRequest(
		urlCatalogCreate,
		catalog,
	)

	res, err := fetch[catalogCreateResult](req)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second)
	catalogs, err := GetCatalogs(catalog.BusinessCenterId)

	if err != nil {
		return nil, err
	} else {
		for _, c := range catalogs {
			if c.Id == res.CatalogId {
				return c, nil
			}
		}
	}

	return nil, errors.New("catalog id not found")
}

func GetCatalogs(business string) ([]*Catalog, error) {
	req := newGetRequest(
		urlCatalogGet,
		withBusinessId(business),
	)

	return fetchAllPages[Catalog](req, 50)
}

func DeleteCatalog(businessId, catalogId string) error {
	var deleteBody struct {
		CatalogId  string `json:"catalog_id"`
		BusinessId string `json:"bc_id"`
	}
	deleteBody.CatalogId = catalogId
	deleteBody.BusinessId = businessId
	req := newPostRequest(
		urlCatalogDelete,
		&deleteBody,
	)
	_, err := fetch[catalogCreateResult](req)

	return err
}

func CatalogAddPixel(businessId, catalogId, accountId, pixelCode string) error {
	type request struct {
		BcId         string `json:"bc_id"`
		CatalogId    string `json:"catalog_id"`
		AdvertiserId string `json:"advertiser_id"`
		PixelCode    string `json:"pixel_code"`
	}
	body := &request{
		BcId:         businessId,
		CatalogId:    catalogId,
		AdvertiserId: accountId,
		PixelCode:    pixelCode,
	}

	req := newPostRequest(
		urlCatalogAddEvent,
		body,
	)

	_, err := fetch[apiResult[map[string]any]](req)

	return err
}
