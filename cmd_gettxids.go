package coinpayments

import "net/url"

type GetTxIdsRequest struct {
	Limit string
	Start string
	Newer string
	//Causes issues parsing response
	//All   string
}

func (r *GetTxIdsRequest) command() string {
	return "get_tx_ids"
}

func (r *GetTxIdsRequest) values() url.Values {
	values := url.Values{}
	if r.Limit != "" {
		values.Set("limit", r.Limit)
	}
	if r.Start != "" {
		values.Set("start", r.Start)
	}
	if r.Newer != "" {
		values.Set("newer", r.Newer)
	}
	//if r.All != "" {
	//	values.Set("all", r.All)
	//}

	return values
}

type GetTxIdsResponse []string

type getTxIdsResult struct {
	errResponse
	Result *GetTxIdsResponse `json:"result"`
}

func (c *Client) GetTxIds(request *GetTxIdsRequest) (*GetTxIdsResponse, error) {
	var resp getTxIdsResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
