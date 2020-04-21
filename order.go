package paypal

import "fmt"

// GetOrder retrieves order by ID
// Endpoint: GET /v2/checkout/orders/ID
func (c *Client) GetOrder(orderID string) (*Order, error) {
	order := &Order{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v2/checkout/orders/", orderID), nil)
	if err != nil {
		return order, err
	}

	if err = c.SendWithAuth(req, order); err != nil {
		return order, err
	}

	return order, nil
}

// CreateOrder - Use this call to create an order
// Endpoint: POST /v2/checkout/orders
func (c *Client) CreateOrder(intent string, purchaseUnits []PurchaseUnitRequest, payer *CreateOrderPayer, appContext *ApplicationContext) (*Order, error) {
	type createOrderRequest struct {
		Intent             string                `json:"intent"`
		Payer              *CreateOrderPayer     `json:"payer,omitempty"`
		PurchaseUnits      []PurchaseUnitRequest `json:"purchase_units"`
		ApplicationContext *ApplicationContext   `json:"application_context,omitempty"`
	}

	order := &Order{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v2/checkout/orders"), createOrderRequest{Intent: intent, PurchaseUnits: purchaseUnits, Payer: payer, ApplicationContext: appContext})
	if err != nil {
		return order, err
	}

	if err = c.SendWithAuth(req, order); err != nil {
		return order, err
	}

	return order, nil
}

// UpdateOrder updates the order by ID
// Endpoint: PATCH /v2/checkout/orders/ID
func (c *Client) UpdateOrder(orderID string, purchaseUnits []PurchaseUnitRequest) (*Order, error) {
	order := &Order{}

	req, err := c.NewRequest("PATCH", fmt.Sprintf("%s%s%s", c.APIBase, "/v2/checkout/orders/", orderID), purchaseUnits)
	if err != nil {
		return order, err
	}

	if err = c.SendWithAuth(req, order); err != nil {
		return order, err
	}

	return order, nil
}

// AuthorizeOrder - https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
// Endpoint: POST /v2/checkout/orders/ID/authorize
func (c *Client) AuthorizeOrder(orderID string, authorizeOrderRequest AuthorizeOrderRequest) (*Authorization, error) {
	auth := &Authorization{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v2/checkout/orders/"+orderID+"/authorize"), authorizeOrderRequest)
	if err != nil {
		return auth, err
	}

	if err = c.SendWithAuth(req, auth); err != nil {
		return auth, err
	}

	return auth, nil
}

// CaptureOrder - https://developer.paypal.com/docs/api/orders/v2/#orders_capture
// Endpoint: POST /v2/checkout/orders/ID/capture
func (c *Client) CaptureOrder(orderID string, captureOrderRequest CaptureOrderRequest) (*CaptureOrderResponse, error) {
	return c.CaptureOrderWithPaypalRequestId(orderID, captureOrderRequest, "")
}

// CaptureOrder with idempotency - https://developer.paypal.com/docs/api/orders/v2/#orders_capture
// Endpoint: POST /v2/checkout/orders/ID/capture
// https://developer.paypal.com/docs/api/reference/api-requests/#http-request-headers
func (c *Client) CaptureOrderWithPaypalRequestId(
	orderID string,
	captureOrderRequest CaptureOrderRequest,
	requestID string,
) (*CaptureOrderResponse, error) {
	capture := &CaptureOrderResponse{}

	c.SetReturnRepresentation()
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v2/checkout/orders/"+orderID+"/capture"), captureOrderRequest)
	if err != nil {
		return capture, err
	}

	if requestID != "" {
		req.Header.Set("PayPal-Request-Id", requestID)
	}

	if err = c.SendWithAuth(req, capture); err != nil {
		return capture, err
	}

	return capture, nil
}

// RefundCapture - https://developer.paypal.com/docs/api/payments/v2/#captures_refund
// Endpoint: POST /v2/payments/captures/ID/refund
func (c *Client) RefundCapture(captureID string, refundCaptureRequest RefundCaptureRequest) (*RefundResponse, error) {
	return c.RefundCaptureWithPaypalRequestId(captureID, refundCaptureRequest, "")
}

// RefundCapture with idempotency - https://developer.paypal.com/docs/api/payments/v2/#captures_refund
// Endpoint: POST /v2/payments/captures/ID/refund
func (c *Client) RefundCaptureWithPaypalRequestId(
	captureID string,
	refundCaptureRequest RefundCaptureRequest,
	requestID string,
) (*RefundResponse, error) {
	refund := &RefundResponse{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v2/payments/captures/"+captureID+"/refund"), refundCaptureRequest)
	if err != nil {
		return refund, err
	}

	if requestID != "" {
		req.Header.Set("PayPal-Request-Id", requestID)
	}

	if err = c.SendWithAuth(req, refund); err != nil {
		return refund, err
	}
	return refund, nil
}
