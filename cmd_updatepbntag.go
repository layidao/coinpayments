package coinpayments

import (
	"net/url"
)

type UpdatePBNTagRequest struct {
	TagID string
	Name  string
	Email string
	URL   string
	Image string
}

func (r *UpdatePBNTagRequest) command() string {
	return "update_pbn_tag"
}

func (r *UpdatePBNTagRequest) values() url.Values {
	values := url.Values{}
	if r.TagID != "" {
		values.Set("tagid", r.TagID)
	}
	if r.Name != "" {
		values.Set("name", r.Name)
	}
	if r.Email != "" {
		values.Set("email", r.Email)
	}
	if r.URL != "" {
		values.Set("url", r.URL)
	}
	if r.Image != "" {
		values.Set("image", r.Image)
	}

	return values
}

type UpdatePBNTagResponse []interface{}

type updatePBNTagResult struct {
	errResponse
	Result *UpdatePBNTagResponse `json:"result"`
}

func (c *Client) UpdatePBNTag(request *UpdatePBNTagRequest) (*UpdatePBNTagResponse, error) {
	var resp updatePBNTagResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
