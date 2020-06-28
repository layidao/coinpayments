package coinpayments

import "net/url"

type ClaimPBNCouponRequest struct {
	Coupon string
}

func (r *ClaimPBNCouponRequest) command() string {
	return "claimm_pbn_coupon"
}

func (r *ClaimPBNCouponRequest) values() url.Values {
	values := url.Values{}
	if r.Coupon != "" {
		values.Set("coupon", r.Coupon)
	}

	return values
}

type ClaimPBNCouponResponse struct {
	TagID string `json:"tagid"`
}

type claimPBNCouponResult struct {
	errResponse
	Result *ClaimPBNCouponResponse `json:"result"`
}

func (c *Client) ClaimPBNCoupon(request *ClaimPBNCouponRequest) (*ClaimPBNCouponResponse, error) {
	var resp claimPBNCouponResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
