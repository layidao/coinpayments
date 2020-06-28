package coinpayments

import "net/url"

type GetWithdrawalInfoRequest struct {
	ID string
}

func (r *GetWithdrawalInfoRequest) command() string {
	return "get_withdrawal_info"
}

func (r *GetWithdrawalInfoRequest) values() url.Values {
	values := url.Values{}
	if r.ID != "" {
		values.Set("id", r.ID)
	}

	return values
}

type GetWithdrawalInfoResponse struct {
	TimeCreated int    `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin        string `json:"coin"`
	Amount      int    `json:"amount"`
	Amountf     string `json:"amountf"`
	Note        string `json:"note"`
	SendAddress string `json:"send_address"`
	SendTXID    string `json:"send_txid"`
}

type getWithdrawalInfoResult struct {
	errResponse
	Result *GetWithdrawalInfoResponse `json:"result"`
}

func (c *Client) GetWithdrawalInfo(request *GetWithdrawalInfoRequest) (*GetWithdrawalInfoResponse, error) {
	var resp getWithdrawalInfoResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
