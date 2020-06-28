package coinpayments

import "net/url"

type CreateWithdrawalRequest struct {
	Amount      string
	AddTxFee    string
	Currency    string
	Currency2   string
	Address     string
	PBNTag      string
	DestTag     string
	IPNURL      string
	AutoConfirm string
	Note        string
}

func (r *CreateWithdrawalRequest) command() string {
	return "create_withdrawal"
}

func (r *CreateWithdrawalRequest) values() url.Values {
	values := url.Values{}
	if r.Amount != "" {
		values.Set("amount", r.Amount)
	}
	if r.AddTxFee != "" {
		values.Set("add_tx_fee", r.AddTxFee)
	}
	if r.Currency != "" {
		values.Set("currency", r.Currency)
	}
	if r.Currency2 != "" {
		values.Set("currency2", r.Currency2)
	}
	if r.Address != "" {
		values.Set("address", r.Address)
	}
	if r.PBNTag != "" {
		values.Set("pbntag", r.PBNTag)
	}
	if r.DestTag != "" {
		values.Set("dest_tag", r.DestTag)
	}
	if r.IPNURL != "" {
		values.Set("ipn_url", r.IPNURL)
	}
	if r.AutoConfirm != "" {
		values.Set("auto_confirm", r.AutoConfirm)
	}
	if r.Note != "" {
		values.Set("note", r.Note)
	}

	return values
}

type CreateWithdrawalResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Amount string `json:"amount"`
}

type createWithdrawalResult struct {
	errResponse
	Result *CreateWithdrawalResponse `json:"result"`
}

func (c *Client) CreateWithdrawal(request *CreateWithdrawalRequest) (*CreateWithdrawalResponse, error) {
	var resp createWithdrawalResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
