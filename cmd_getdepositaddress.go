package coinpayments

import "net/url"

type GetDepositAddressRequest struct {
	Currency string
}

func (r *GetDepositAddressRequest) command() string {
	return "get_deposit_address"
}

func (r *GetDepositAddressRequest) values() url.Values {
	values := url.Values{}
	if r.Currency != "" {
		values.Set("currency", r.Currency)
	}

	return values
}

type GetDepositAddressResponse struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	DestTag int    `json:"dest_tag"`
}

type getDepositAddressResult struct {
	errResponse
	Result *GetDepositAddressResponse `json:"result"`
}

func (c *Client) GetDepositAddress(request *GetDepositAddressRequest) (*GetDepositAddressResponse, error) {
	var resp getDepositAddressResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
