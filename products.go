package paypal

import (
	"fmt"
	"net/http"
)

type (
	ProductType     string
	ProductCategory string

	// Product struct
	Product struct {
		ID          string          `json:"id,omitempty"`
		Name        string          `json:"name,omitempty"`
		Description string          `json:"description"`
		Category    ProductCategory `json:"category,omitempty"`
		Type        ProductType     `json:"type"`
		ImageUrl    string          `json:"image_url"`
		HomeUrl     string          `json:"home_url"`
	}

	CreateProductResponse struct {
		Product
		SharedResponse
	}

	ListProductsResponse struct {
		Products []Product `json:"products"`
		SharedListResponse
	}

	ProductListParameters struct {
		ListParams
	}
)

func (self *Product) GetUpdatePatch() []Patch {
	return []Patch{
		{
			Operation: "replace",
			Path:      "/description",
			Value:     self.Description,
		},
		{
			Operation: "replace",
			Path:      "/category",
			Value:     self.Category,
		},
		{
			Operation: "replace",
			Path:      "/image_url",
			Value:     self.ImageUrl,
		},
		{
			Operation: "replace",
			Path:      "/home_url",
			Value:     self.HomeUrl,
		},
	}
}

const (
	PRODUCT_TYPE_PHYSICAL ProductType = "PHYSICAL"
	PRODUCT_TYPE_DIGITAL  ProductType = "DIGITAL"
	PRODUCT_TYPE_SERVICE  ProductType = "SERVICE"

	PRODUCT_CATEGORY_SOFTWARE ProductCategory = "software"
)

// CreateProduct creates a product
// Doc: https://developer.paypal.com/docs/api/catalog-products/v1/#products_create
// Endpoint: POST /v1/catalogs/products
func (c *Client) CreateProduct(product Product) (*CreateProductResponse, error) {
	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.APIBase, "/v1/catalogs/products"), product)
	response := &CreateProductResponse{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}

// UpdateProduct. updates a product information
// Doc: https://developer.paypal.com/docs/api/catalog-products/v1/#products_patch
// Endpoint: PATCH /v1/catalogs/products/:product_id
func (c *Client) UpdateProduct(product Product) error {
	req, err := c.NewRequest(http.MethodPatch, fmt.Sprintf("%s%s%s", c.APIBase, "/v1/catalogs/products/", product.ID), product.GetUpdatePatch())
	if err != nil {
		return err
	}
	err = c.SendWithAuth(req, nil)
	return err
}

// Get product details
// Doc: https://developer.paypal.com/docs/api/catalog-products/v1/#products_get
// Endpoint: GET /v1/catalogs/products/:product_id
func (c *Client) GetProduct(productId string) (*Product, error) {
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s%s", c.APIBase, "/v1/catalogs/products/", productId), nil)
	response := &Product{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}

// List all products
// Doc: https://developer.paypal.com/docs/api/catalog-products/v1/#products_list
// Endpoint: GET /v1/catalogs/products
func (c *Client) ListProducts(params *ProductListParameters) (*ListProductsResponse, error) {
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.APIBase, "/v1/catalogs/products"), nil)
	response := &ListProductsResponse{}
	if err != nil {
		return response, err
	}

	if params != nil {
		q := req.URL.Query()
		q.Add("page", params.Page)
		q.Add("page_size", params.PageSize)
		q.Add("total_required", params.TotalRequired)
		req.URL.RawQuery = q.Encode()
	}

	err = c.SendWithAuth(req, response)
	return response, err
}
