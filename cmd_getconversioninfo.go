package coinpayments

import "net/url"

type GetConversionInfoRequest struct {
	ID string
}

func (r *GetConversionInfoRequest) command() string {
	return "get_conversion_info"
}

func (r *GetConversionInfoRequest) values() url.Values {
	values := url.Values{}
	if r.ID != "" {
		values.Set("id", r.ID)
	}

	return values
}

type GetConversionInfoResponse struct {
	TimeCreated string `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin1       string `json:"coin1"`
	Coin2       string `json:"coin2"`
	AmountSent  int    `json:"amount_sent"`
	AmountSentf string `json:"amount_sentf"`
	Received    int    `json:"received"`
	Receivedf   string `json:"receivedf"`
}

type getConversionInfoResult struct {
	errResponse
	Result *GetConversionInfoResponse `json:"result"`
}

func (c *Client) GetConversionInfo(request *GetConversionInfoRequest) (*GetConversionInfoResponse, error) {
	var resp getConversionInfoResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
