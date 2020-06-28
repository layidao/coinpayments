package coinpayments

import "net/url"

type GetTxInfoRequest struct {
	TXID string
	Full string
}

func (r *GetTxInfoRequest) command() string {
	return "get_tx_info"
}

func (r *GetTxInfoRequest) values() url.Values {
	values := url.Values{}

	if r.TXID != "" {
		values.Set("txid", r.TXID)
	}
	if r.Full != "" {
		values.Set("full", r.Full)
	}
	return values
}

type GetTxInfoResponse struct {
	TimeCreated      int    `json:"time_created"`
	TimeExpires      int    `json:"time_expires"`
	Status           int    `json:"status"`
	StatusText       string `json:"status_text"`
	Type             string `json:"type"`
	Coin             string `json:"coin"`
	Amount           int    `json:"amount"`
	Amountf          string `json:"amountf"`
	Received         int    `json:"received"`
	Receivedf        string `json:"receivedf"`
	ReceivedConfirms int    `json:"recv_confirms"`
	PaymentAddress   string `json:"payment_address"`
	Checkout         struct {
		Currency   string        `json:"currency"`
		Amount     int           `json:"amount"`
		Test       int           `json:"test"`
		ItemNumber string        `json:"item_number"`
		ItemName   string        `json:"item_name"`
		Details    []interface{} `json:"details"`
		Invoice    string        `json:"invoice"`
		Custom     string        `json:"custom"`
		IPNURL     string        `json:"ipn_url"`
		Amountf    int           `json:"amountf"`
	} `json:"checkout,omitempty"`
	Shipping []interface{} `json:"shipping,omitempty"`
}

type getTxInfoResult struct {
	errResponse
	Result *GetTxInfoResponse `json:"result"`
}

func (c *Client) GetTxInfo(request *GetTxInfoRequest) (*GetTxInfoResponse, error) {
	var resp getTxInfoResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
