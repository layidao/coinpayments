package coinpayments

import (
	"net/url"
)

type RatesRequest struct {
	Short    string
	Accepted string
}

func (r *RatesRequest) command() string {
	return "rates"
}

func (r *RatesRequest) values() url.Values {
	values := url.Values{}

	if r.Short != "" {
		values.Set("short", r.Short)
	}
	if r.Accepted != "" {
		values.Set("accepted", r.Accepted)
	}
	return values
}

type RatesResponse map[string]struct {
	IsFiat       int      `json:"is_fiat"`
	RateBTC      string   `json:"rate_btc"`
	LastUpdate   string   `json:"last_update"`
	TxFee        string   `json:"tx_fee"`
	Status       string   `json:"status"`
	Name         string   `json:"name"`
	Confirms     string   `json:"confirms"`
	Capabilities []string `json:"capabilities"`
	Accepted     int      `json:"accepted"`
}

type ratesResult struct {
	errResponse
	Result *RatesResponse `json:"result"`
}

func (c *Client) Rates(request *RatesRequest) (*RatesResponse, error) {
	var resp ratesResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
