package coinpayments

import "net/url"

type ClaimPBNTagRequest struct {
	TagID string
	Name  string
}

func (r *ClaimPBNTagRequest) command() string {
	return "claim_pbn_tag"
}

func (r *ClaimPBNTagRequest) values() url.Values {
	values := url.Values{}
	if r.TagID != "" {
		values.Set("tagid", r.TagID)
	}
	if r.Name != "" {
		values.Set("name", r.Name)
	}

	return values
}

type ClaimPBNTagResponse []interface{}

type claimPBNTagResult struct {
	errResponse
	Result *ClaimPBNTagResponse `json:"result"`
}

func (c *Client) ClaimPBNTag(request *ClaimPBNTagRequest) (*ClaimPBNTagResponse, error) {
	var resp claimPBNTagResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
