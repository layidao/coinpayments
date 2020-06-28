package coinpayments

import "net/url"

type ConvertRequest struct {
	Amount  string
	From    string
	To      string
	Address string
	DestTag string
}

func (r *ConvertRequest) command() string {
	return "convert"
}

func (r *ConvertRequest) values() url.Values {
	values := url.Values{}
	if r.Amount != "" {
		values.Set("amount", r.Amount)
	}
	if r.From != "" {
		values.Set("from", r.From)
	}
	if r.To != "" {
		values.Set("to", r.To)
	}
	if r.Address != "" {
		values.Set("address", r.Address)
	}
	if r.DestTag != "" {
		values.Set("dest_tag", r.DestTag)
	}

	return values
}

type ConvertResponse struct {
	ID string `json:"id"`
}

type convertResult struct {
	errResponse
	Result *ConvertResponse `json:"result"`
}

func (c *Client) Convert(request *ConvertRequest) (*ConvertResponse, error) {
	var resp convertResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
