package coinpayments

import "net/url"

type GetTxInfoMultiRequest struct {
	TXID string
}

func (r *GetTxInfoMultiRequest) command() string {
	return "get_tx_info_multi"
}

func (r *GetTxInfoMultiRequest) values() url.Values {
	values := url.Values{}
	if r.TXID != "" {
		values.Set("txid", r.TXID)
	}
	return values
}

type GetTxInfoMultiResponse map[string]struct {
	Error            string `json:"error"`
	TimeCreated      int    `json:"time_created"`
	TimeExpires      int    `json:"time_expires"`
	Status           int    `json:"status"`
	StatusText       string `json:"status_text"`
	Type             string `json:"type"`
	Coin             string `json:"coin"`
	Amount           int    `json:"amount"`
	Amountf          string `json:"amountf"`
	Received         int    `json:"received"`
	Recievedf        string `json:"recievedf"`
	RecievedConfirms int    `json:"recv_confirms"`
	PaymentAddress   string `json:"payment_address"`
}

type getTxInfoMultiResult struct {
	errResponse
	Result *GetTxInfoMultiResponse `json:"result"`
}

func (c *Client) GetTxInfoMulti(request *GetTxInfoMultiRequest) (*GetTxInfoMultiResponse, error) {
	var resp getTxInfoMultiResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
