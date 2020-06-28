package coinpayments

import "net/url"

type GetBasicInfoRequest struct{}

func (r *GetBasicInfoRequest) command() string {
	return "get_basic_info"
}

func (r *GetBasicInfoRequest) values() url.Values {
	values := url.Values{}

	return values
}

type GetBasicInfoResponse struct {
	Username   string `json:"username"`
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
	PublicName string `json:"public_name"`
}

type getBasicInfoResult struct {
	errResponse
	Result *GetBasicInfoResponse `json:"result"`
}

func (c *Client) GetBasicInfo(request *GetBasicInfoRequest) (*GetBasicInfoResponse, error) {
	var resp getBasicInfoResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
