package coinpayments

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type IPN struct {
	ipnInformation
	depositInformation            depositInformation
	withdrawalInformation         withdrawalInformation
	buyerInformation              buyerInformation
	shippingInformation           shippingInformation
	simpleButtonFields            simpleButtonFields
	advancedButtonFields          advancedButtonFields
	shoppingCartButtonFields      shoppingCartButtonFields
	donationButtonFields          donationButtonFields
	apiGeneratedTransactionFields apiGeneratedTransactionFields
}

type ipnInformation struct {
	IPNVersion string `json:"ipn_version"`
	IPNType    string `json:"ipn_type"`
	IPNMode    string `json:"ipn_mode"`
	IPNId      string `json:"ipn_id"`
	Merchant   string `json:"merchant"`
}

type simpleIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	simpleButtonFields
}

func (i *IPN) ToSimpleIPN() (*simpleIPN, error) {
	if i.IPNType != "simple" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'simple'")
	}
	return &simpleIPN{
		ipnInformation:      i.ipnInformation,
		buyerInformation:    i.buyerInformation,
		shippingInformation: i.shippingInformation,
		simpleButtonFields:  i.simpleButtonFields,
	}, nil
}

type buttonIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	advancedButtonFields
}

func (i *IPN) ToButtonIPN() (*buttonIPN, error) {
	if i.IPNType != "button" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'button'")
	}
	return &buttonIPN{
		ipnInformation:       i.ipnInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		advancedButtonFields: i.advancedButtonFields,
	}, nil
}

type cartIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	shoppingCartButtonFields
}

func (i *IPN) ToCartIPN() (*cartIPN, error) {
	if i.IPNType != "cart" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'cart'")
	}
	return &cartIPN{
		ipnInformation:           i.ipnInformation,
		buyerInformation:         i.buyerInformation,
		shippingInformation:      i.shippingInformation,
		shoppingCartButtonFields: i.shoppingCartButtonFields,
	}, nil
}

type donationIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	donationButtonFields
}

func (i *IPN) ToDonationIPN() (*donationIPN, error) {
	if i.IPNType != "donation" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'donation'")
	}
	return &donationIPN{
		ipnInformation:       i.ipnInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		donationButtonFields: i.donationButtonFields,
	}, nil
}

type depositIPN struct {
	ipnInformation
	depositInformation
}

func (i *IPN) ToDepositIPN() (*depositIPN, error) {
	if i.IPNType != "deposit" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'deposit'")
	}
	return &depositIPN{
		ipnInformation:     i.ipnInformation,
		depositInformation: i.depositInformation,
	}, nil
}

type withdrawalIPN struct {
	ipnInformation
	withdrawalInformation
}

func (i *IPN) ToWithdrawalIPN() (*withdrawalIPN, error) {
	if i.IPNType != "withdrawal" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'withdrawal'")
	}
	return &withdrawalIPN{
		ipnInformation:        i.ipnInformation,
		withdrawalInformation: i.withdrawalInformation,
	}, nil
}

type apiIPN struct {
	ipnInformation
	apiGeneratedTransactionFields
}

func (i *IPN) ToApiIPN() (*apiIPN, error) {
	if i.IPNType != "api" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'api'")
	}
	return &apiIPN{
		ipnInformation:                i.ipnInformation,
		apiGeneratedTransactionFields: i.apiGeneratedTransactionFields,
	}, nil
}

func (c *Client) ParseIPN(r *http.Request, ipnSecret string) (*IPN, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("coinpayments: error reading request body - %v", err)
	}

	if ipnSecret != "" {
		hmac := r.Header.Get("HMAC")

		genHMAC, err := c.makeIPNHMAC(string(data), ipnSecret)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error generating ipn HMAC - %v", err)
		}

		if hmac != genHMAC {
			return nil, fmt.Errorf("coinpayments: could not validate server HMAC")
		}
	}

	values, err := url.ParseQuery(string(data))
	if err != nil {
		return nil, err
	}

	ipn := &IPN{
		ipnInformation: ipnInformation{
			IPNVersion: values.Get("ipn_version"),
			IPNType:    values.Get("ipn_type"),
			IPNMode:    values.Get("ipn_mode"),
			IPNId:      values.Get("ipn_id"),
			Merchant:   values.Get("merchant"),
		},
	}

	switch ipn.IPNType {
	case "simple":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.simpleButtonFields = simpleButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemAmount:       values.Get("item_amount"),
			ItemName:         values.Get("item_name"),
			ItemDescription:  values.Get("item_desc"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "button":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.advancedButtonFields = advancedButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemAmount:       values.Get("item_amount"),
			ItemName:         values.Get("item_name"),
			Quantity:         values.Get("quantity"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "cart":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.shoppingCartButtonFields = shoppingCartButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			ItemName:         values.Get("item_name_#"),
			ItemAmount:       values.Get("item_amount_#"),
			ItemQuantity:     values.Get("item_quantity_#"),
			ItemNumber:       values.Get("item_number_#"),
			Option1Name:      values.Get("item_on1_#"),
			Option1Value:     values.Get("item_ov1_#"),
			Option2Name:      values.Get("item_on2_#"),
			Option2Value:     values.Get("item_ov2_#"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "donation":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.donationButtonFields = donationButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemName:         values.Get("item_name"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "deposit":
		ipn.depositInformation = depositInformation{
			TransactionID: values.Get("txn_id"),
			Address:       values.Get("address"),
			DestTag:       values.Get("dest_tag"),
			Status:        values.Get("status"),
			StatusText:    values.Get("status_text"),
			Currency:      values.Get("currency"),
			Confirms:      values.Get("confirms"),
			Amount:        values.Get("amount"),
			Amounti:       values.Get("amounti"),
			Fee:           values.Get("fee"),
			Feei:          values.Get("feei"),
			FiatCoin:      values.Get("fiat_coin"),
			FiatAmount:    values.Get("fiat_amount"),
			FiatAmounti:   values.Get("fiat_amounti"),
			FiatFee:       values.Get("fiat_fee"),
			FiatFeei:      values.Get("fiat_feei"),
		}
	case "withdrawal":
		ipn.withdrawalInformation = withdrawalInformation{
			ID:            values.Get("id"),
			Status:        values.Get("status"),
			StatusText:    values.Get("status_text"),
			Address:       values.Get("address"),
			TransactionID: values.Get("txn_id"),
			Currency:      values.Get("currency"),
			Amount:        values.Get("amount"),
			Amounti:       values.Get("amounti"),
		}
	case "api":
		ipn.apiGeneratedTransactionFields = apiGeneratedTransactionFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Fee:              values.Get("fee"),
			BuyerName:        values.Get("buyer_name"),
			Email:            values.Get("email"),
			ItemName:         values.Get("item_name"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	}

	return ipn, nil
}

type depositInformation struct {
	TransactionID string `json:"txn_id"`
	Address       string `json:"address"`
	DestTag       string `json:"dest_tag"`
	Status        string `json:"status"`
	StatusText    string `json:"status_text"`
	Currency      string `json:"currency"`
	Confirms      string `json:"confirms"`
	Amount        string `json:"amount"`
	Amounti       string `json:"amounti"`
	Fee           string `json:"fee"`
	Feei          string `json:"feei"`
	FiatCoin      string `json:"fiat_coin"`
	FiatAmount    string `json:"fiat_amount"`
	FiatAmounti   string `json:"fiat_amounti"`
	FiatFee       string `json:"fiat_fee"`
	FiatFeei      string `json:"fiat_feei"`
}

type withdrawalInformation struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	StatusText    string `json:"status_text"`
	Address       string `json:"address"`
	TransactionID string `json:"txn_id"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	Amounti       string `json:"amounti"`
}

type buyerInformation struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Email     string `json:"email"`
}

type shippingInformation struct {
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip"`
	Country     string `json:"country"`
	CountryName string `json:"country_name"`
	Phone       string `json:"phone"`
}

type simpleButtonFields struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TransactionID    string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Subtotal         string `json:"subtotal"`
	Shipping         string `json:"shipping"`
	Tax              string `json:"tax"`
	Fee              string `json:"fee"`
	Net              string `json:"net"`
	ItemAmount       string `json:"item_amount"`
	ItemName         string `json:"item_name"`
	ItemDescription  string `json:"item_desc"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	Option1Name      string `json:"on1"`
	Option1Value     string `json:"ov1"`
	Option2Name      string `json:"on2"`
	Option2Value     string `json:"ov2"`
	SendTransaction  string `json:"send_tx"`
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

type advancedButtonFields struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TransactionID    string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Subtotal         string `json:"subtotal"`
	Shipping         string `json:"shipping"`
	Tax              string `json:"tax"`
	Fee              string `json:"fee"`
	Net              string `json:"net"`
	ItemAmount       string `json:"item_amount"`
	ItemName         string `json:"item_name"`
	Quantity         string `json:"quantity"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	Option1Name      string `json:"on1"`
	Option1Value     string `json:"ov1"`
	Option2Name      string `json:"on2"`
	Option2Value     string `json:"ov2"`
	Extra            string `json:"extra"`
	SendTransaction  string `json:"send_tx"`
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

type shoppingCartButtonFields struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TransactionID    string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Subtotal         string `json:"subtotal"`
	Shipping         string `json:"shipping"`
	Tax              string `json:"tax"`
	Fee              string `json:"fee"`
	ItemName         string `json:"item_name_#"`
	ItemAmount       string `json:"item_amount_#"`
	ItemQuantity     string `json:"item_quantity_#"`
	ItemNumber       string `json:"item_number_#"`
	Option1Name      string `json:"item_on1_#"`
	Option1Value     string `json:"item_ov1_#"`
	Option2Name      string `json:"item_on2_#"`
	Option2Value     string `json:"item_ov2_#"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	Extra            string `json:"extra"`
	SendTransaction  string `json:"send_tx"`
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

type donationButtonFields struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TransactionID    string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Subtotal         string `json:"subtotal"`
	Shipping         string `json:"shipping"`
	Tax              string `json:"tax"`
	Fee              string `json:"fee"`
	Net              string `json:"net"`
	ItemName         string `json:"item_name"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	Option1Name      string `json:"on1"`
	Option1Value     string `json:"ov1"`
	Option2Name      string `json:"on2"`
	Option2Value     string `json:"ov2"`
	Extra            string `json:"extra"`
	SendTransaction  string `json:"send_tx"`
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

type apiGeneratedTransactionFields struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TransactionID    string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Fee              string `json:"fee"`
	BuyerName        string `json:"buyer_name"`
	Email            string `json:"email"`
	ItemName         string `json:"item_name"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	SendTransaction  string `json:"send_tx"`
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}
