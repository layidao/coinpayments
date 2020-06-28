package coinpayments

import "net/url"

type CreateTransferRequest struct {
	Amount      string
	Currency    string
	Merchant    string
	PBNTag      string
	AutoConfirm string
	Note        string
}

func (r *CreateTransferRequest) command() string {
	return "create_transfer"
}

func (r *CreateTransferRequest) values() url.Values {
	values := url.Values{}
	if r.Amount != "" {
		values.Set("amount", r.Amount)
	}
	if r.Currency != "" {
		values.Set("currency", r.Currency)
	}
	if r.Merchant != "" {
		values.Set("merchant", r.Merchant)
	}
	if r.PBNTag != "" {
		values.Set("pbntag", r.PBNTag)
	}
	if r.AutoConfirm != "" {
		values.Set("auto_confirm", r.AutoConfirm)
	}
	if r.Note != "" {
		values.Set("note", r.Note)
	}

	return values
}

type CreateTransferResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type createTransferResult struct {
	errResponse
	Result *CreateTransferResponse `json:"result"`
}

func (c *Client) CreateTransfer(request *CreateTransferRequest) (*CreateTransferResponse, error) {
	var resp createTransferResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
