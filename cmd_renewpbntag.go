package coinpayments

import "net/url"

type RenewPBNTagRequest struct {
	TagID string
	Coin  string
	Years string
}

func (r *RenewPBNTagRequest) command() string {
	return "renew_pbn_tag"
}

func (r *RenewPBNTagRequest) values() url.Values {
	values := url.Values{}
	if r.TagID != "" {
		values.Set("tagid", r.TagID)
	}
	if r.Coin != "" {
		values.Set("coin", r.Coin)
	}
	if r.Years != "" {
		values.Set("years", r.Years)
	}

	return values
}

type RenewPBNTagResponse []interface{}

type renewPBNTagResult struct {
	errResponse
	Result *RenewPBNTagResponse `json:"result"`
}

func (c *Client) RenewPBNTag(request *RenewPBNTagRequest) (*RenewPBNTagResponse, error) {
	var resp renewPBNTagResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
