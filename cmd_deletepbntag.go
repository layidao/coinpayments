package coinpayments

import "net/url"

type DeletePBNTagRequest struct {
	TagID string
}

func (r *DeletePBNTagRequest) command() string {
	return "delete_pbn_tag"
}

func (r *DeletePBNTagRequest) values() url.Values {
	values := url.Values{}
	if r.TagID != "" {
		values.Set("tagid", r.TagID)
	}

	return values
}

type DeletePBNTagResponse []interface{}

type deletePBNTagResult struct {
	errResponse
	Result *DeletePBNTagResponse `json:"result"`
}

func (c *Client) DeletePBNTag(request *DeletePBNTagRequest) (*DeletePBNTagResponse, error) {
	var resp deletePBNTagResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
