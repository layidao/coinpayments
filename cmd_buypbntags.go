package coinpayments

import "net/url"

type BuyPBNTagsRequest struct {
	Coin   string
	Number string
}

func (r *BuyPBNTagsRequest) command() string {
	return "buy_pbn_tags"
}

func (r *BuyPBNTagsRequest) values() url.Values {
	values := url.Values{}
	if r.Coin != "" {
		values.Set("coin", r.Coin)
	}
	if r.Number != "" {
		values.Set("num", r.Number)
	}

	return values
}

type BuyPBNTagsResponse []interface{}

type buyPBNTagsResult struct {
	errResponse
	Result *BuyPBNTagsResponse `json:"result"`
}

func (c *Client) BuyPBNTags(request *BuyPBNTagsRequest) (*BuyPBNTagsResponse, error) {
	var resp buyPBNTagsResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
