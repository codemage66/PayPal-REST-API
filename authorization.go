package paypalsdk

import (
	"bytes"
	"fmt"
	"net/http"
)

// GetAuthorization returns an authorization by ID
// Endpoint: GET /v1/payments/authorization/ID
func (c *Client) GetAuthorization(authID string) (*Authorization, error) {
	buf := bytes.NewBuffer([]byte(""))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/payments/authorization/", authID), buf)
	if err != nil {
		return &Authorization{}, err
	}

	auth := &Authorization{}

	err = c.SendWithAuth(req, auth)
	if err != nil {
		return auth, err
	}

	return auth, nil
}

// CaptureAuthorization captures and process an existing authorization.
// To use this method, the original payment must have Intent set to "authorize"
// Endpoint: POST /v1/payments/authorization/ID/capture
func (c *Client) CaptureAuthorization(authID string, a *Amount, isFinalCapture bool) (*Capture, error) {
	isFinalStr := "false"
	if isFinalCapture {
		isFinalStr = "true"
	}
	buf := bytes.NewBuffer([]byte("{\"amount\":{\"currency\":\"" + a.Currency + "\",\"total\":\"" + a.Total + "\"},\"is_final_capture\":" + isFinalStr + "}"))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/authorization/"+authID+"/capture"), buf)
	if err != nil {
		return &Capture{}, err
	}

	capture := &Capture{}

	err = c.SendWithAuth(req, capture)
	if err != nil {
		return capture, err
	}

	return capture, nil
}

// VoidAuthorization voids a previously authorized payment
// Endpoint: POST /v1/payments/authorization/ID/void
func (c *Client) VoidAuthorization(authID string) (*Authorization, error) {
	buf := bytes.NewBuffer([]byte(""))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/authorization/"+authID+"/void"), buf)
	if err != nil {
		return &Authorization{}, err
	}

	auth := &Authorization{}

	err = c.SendWithAuth(req, auth)
	if err != nil {
		return auth, err
	}

	return auth, nil
}

// ReauthorizeAuthorization reauthorize a Paypal account payment.
// PayPal recommends to reauthorize payment after ~3 days
// Endpoint: POST /v1/payments/authorization/ID/reauthorize
func (c *Client) ReauthorizeAuthorization(authID string, a *Amount) (*Authorization, error) {
	buf := bytes.NewBuffer([]byte("{\"amount\":{\"currency\":\"" + a.Currency + "\",\"total\":\"" + a.Total + "\"}}"))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/authorization/"+authID+"/reauthorize"), buf)
	if err != nil {
		return &Authorization{}, err
	}

	auth := &Authorization{}

	err = c.SendWithAuth(req, auth)
	if err != nil {
		return auth, err
	}

	return auth, nil
}
