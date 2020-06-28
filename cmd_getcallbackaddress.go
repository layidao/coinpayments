package coinpayments

import "net/url"

type GetCallbackAddressRequest struct {
	Currency string
	IPNURL   string
	Label    string
}

func (r *GetCallbackAddressRequest) command() string {
	return "get_callback_address"
}

func (r *GetCallbackAddressRequest) values() url.Values {
	values := url.Values{}

	if r.Currency != "" {
		values.Set("currency", r.Currency)
	}
	if r.IPNURL != "" {
		values.Set("ipn_url", r.IPNURL)
	}
	if r.Label != "" {
		values.Set("label", r.Label)
	}
	return values
}

type GetCallbackAddressResponse struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	DestTag string `json:"dest_tag"`
}

type getCallbackAddressResult struct {
	errResponse
	Result *GetCallbackAddressResponse `json:"result"`
}

func (c *Client) GetCallbackAddress(request *GetCallbackAddressRequest) (*GetCallbackAddressResponse, error) {
	var resp getCallbackAddressResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
