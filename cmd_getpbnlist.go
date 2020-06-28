package coinpayments

import "net/url"

type GetPBNListRequest struct{}

func (r *GetPBNListRequest) command() string {
	return "get_pbn_list"
}

func (r *GetPBNListRequest) values() url.Values {
	values := url.Values{}

	return values
}

type GetPBNListResponse []struct {
	TagID       string `json:"tagid"`
	PBGTag      string `json:"pbgtag"`
	TimeExpires int    `json:"time_expires"`
}

type getPBNListResult struct {
	errResponse
	Result *GetPBNListResponse `json:"result"`
}

func (c *Client) GetPBNList(request *GetPBNListRequest) (*GetPBNListResponse, error) {
	var resp getPBNListResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
