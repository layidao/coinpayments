package coinpayments

import "net/url"

type GetWithdrawalHistoryRequest struct {
	Limit string
	Start string
	Newer string
}

func (r *GetWithdrawalHistoryRequest) command() string {
	return "get_withdrawal_history"
}

func (r *GetWithdrawalHistoryRequest) values() url.Values {
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

	return values
}

type GetWithdrawalHistoryResponse []struct {
	ID          string `json:"id"`
	TimeCreated int    `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin        string `json:"coin"`
	Amount      int    `json:"amount"`
	Amountf     string `json:"amountf"`
	Note        string `json:"note"`
	SendAddress string `json:"send_address"`
	SendDestTag string `json:"send_dest_tag"`
	SendTXID    string `json:"send_txid"`
}

type getWithdrawalHistoryResult struct {
	errResponse
	Result *GetWithdrawalHistoryResponse `json:"result"`
}

func (c *Client) GetWithdrawalHistory(request *GetWithdrawalHistoryRequest) (*GetWithdrawalHistoryResponse, error) {
	var resp getWithdrawalHistoryResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
