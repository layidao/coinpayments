package coinpayments

import "net/url"

type GetPBNInfoRequest struct {
	PBNTag string
}

func (r *GetPBNInfoRequest) command() string {
	return "get_pbn_info"
}

func (r *GetPBNInfoRequest) values() url.Values {
	values := url.Values{}
	if r.PBNTag != "" {
		values.Set("pbntag", r.PBNTag)
	}

	return values
}

type GetPBNInfoResponse struct {
	PBNTag       string `json:"pbntag"`
	Merchant     string `json:"merchant"`
	ProfileName  string `json:"profile_name"`
	ProfileURL   string `json:"profile_url"`
	ProfileEmail string `json:"profile_email"`
	ProfileImage string `json:"profile_image"`
	MemberSince  int    `json:"member_since"`
	Feedback     struct {
		Positive int    `json:"pos"`
		Negative int    `json:"neg"`
		Neutral  string `json:"neut"`
		Total    int    `json:"total"`
		Percent  string `json:"percent"`
	} `json:"feedback"`
}

type getPBNInfoResult struct {
	errResponse
	Result *GetPBNInfoResponse `json:"result"`
}

func (c *Client) GetPBNInfo(request *GetPBNInfoRequest) (*GetPBNInfoResponse, error) {
	var resp getPBNInfoResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
