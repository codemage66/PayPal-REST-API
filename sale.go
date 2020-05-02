package paypal

import "fmt"

// GetSale returns a sale by ID
// Use this call to get details about a sale transaction.
// Note: This call returns only the sales that were created via the REST API.
// Endpoint: GET /v1/payments/sale/ID
func (c *Client) GetSale(saleID string) (*Sale, error) {
	sale := &Sale{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/sale/"+saleID), nil)
	if err != nil {
		return sale, err
	}

	if err = c.SendWithAuth(req, sale); err != nil {
		return sale, err
	}

	return sale, nil
}

// RefundSale refunds a completed payment.
// Use this call to refund a completed payment. Provide the sale_id in the URI and an empty JSON payload for a full refund. For partial refunds, you can include an amount.
func (c *Client) RefundSale(saleID string, r RefundRequest) (*Refund, error) {
	refund := &Refund{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/sale/"+saleID+"/refund"), &r)
	if err != nil {
		return refund, err
	}

	if err = c.SendWithAuth(req, refund); err != nil {
		return refund, err
	}

	return refund, nil
}

// GetRefund by ID
// Use it to look up details of a specific refund on direct and captured payments.
// Endpoint: GET /v2/payments/refund/ID
func (c *Client) GetRefund(refundID string) (*Refund, error) {
	refund := &Refund{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v2/payments/refund/"+refundID), nil)
	if err != nil {
		return refund, err
	}

	if err = c.SendWithAuth(req, refund); err != nil {
		return refund, err
	}

	return refund, nil
}
