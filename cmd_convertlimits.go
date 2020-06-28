package coinpayments

import "net/url"

type ConvertLimitsRequest struct {
	From string
	To   string
}

func (r *ConvertLimitsRequest) command() string {
	return "convert_limits"
}

func (r *ConvertLimitsRequest) values() url.Values {
	values := url.Values{}
	if r.From != "" {
		values.Set("from", r.From)
	}
	if r.To != "" {
		values.Set("to", r.To)
	}

	return values
}

type ConvertLimitsResponse struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type convertLimitsResult struct {
	errResponse
	Result *ConvertLimitsResponse `json:"result"`
}

func (c *Client) ConvertLimits(request *ConvertLimitsRequest) (*ConvertLimitsResponse, error) {
	var resp convertLimitsResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
