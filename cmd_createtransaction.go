package coinpayments

import (
	"net/url"
)

type CreateTransactionRequest struct {
	Amount     string
	Currency1  string
	Currency2  string
	BuyerEmail string
	Address    string
	BuyerName  string
	ItemName   string
	ItemNumber string
	Invoice    string
	Custom     string
	IPNURL     string
	SuccessURL string
	CancelURL  string
}

func (r *CreateTransactionRequest) command() string {
	return "create_transaction"
}

func (r *CreateTransactionRequest) values() url.Values {
	values := url.Values{}

	if r.Amount != "" {
		values.Set("amount", r.Amount)
	}
	if r.Currency1 != "" {
		values.Set("currency1", r.Currency1)
	}
	if r.Currency2 != "" {
		values.Set("currency2", r.Currency2)
	}
	if r.BuyerEmail != "" {
		values.Set("buyer_email", r.BuyerEmail)
	}
	if r.Address != "" {
		values.Set("address", r.Address)
	}
	if r.BuyerName != "" {
		values.Set("buyer_name", r.BuyerName)
	}
	if r.ItemName != "" {
		values.Set("item_name", r.ItemName)
	}
	if r.ItemNumber != "" {
		values.Set("item_number", r.ItemNumber)
	}
	if r.Invoice != "" {
		values.Set("invoice", r.Invoice)
	}
	if r.Custom != "" {
		values.Set("custom", r.Custom)
	}
	if r.IPNURL != "" {
		values.Set("ipn_url", r.IPNURL)
	}
	if r.SuccessURL != "" {
		values.Set("success_url", r.SuccessURL)
	}
	if r.CancelURL != "" {
		values.Set("cancel_url", r.CancelURL)
	}
	return values
}

type CreateTransactionResponse struct {
	Amount         string `json:"amount"`
	Address        string `json:"address"`
	DestTag        string `json:"dest_tag"`
	TxnId          string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        int    `json:"timeout"`
	CheckoutURL    string `json:"checkout_url"`
	StatusURL      string `json:"status_url"`
	QRCodeURL      string `json:"qrcode_url"`
}

type createTransactionResult struct {
	errResponse
	Result *CreateTransactionResponse `json:"result"`
}

func (c *Client) CreateTransaction(request *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	var resp createTransactionResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
