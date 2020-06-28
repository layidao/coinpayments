package coinpayments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPN struct {
	IPNInformation
	depositInformation            *depositInformation            `json:"-"`
	withdrawalInformation         *withdrawalInformation         `json:"-"`
	buyerInformation              *buyerInformation              `json:"-"`
	shippingInformation           *shippingInformation           `json:"-"`
	simpleButtonFields            *simpleButtonFields            `json:"-"`
	advancedButtonFields          *advancedButtonFields          `json:"-"`
	shoppingCartButtonFields      *shoppingCartButtonFields      `json:"-"`
	donationButtonFields          *donationButtonFields          `json:"-"`
	apiGeneratedTransactionFields *apiGeneratedTransactionFields `json:"-"`
}

type IPNInformation struct {
	IPNVersion string `json:"ipn_version"`
	IPNType    string `json:"ipn_type"`
	IPNMode    string `json:"ipn_mode"`
	IPNID      string `json:"ipn_id"`
	Merchant   string `json:"merchant"`
}

type simpleIPN struct {
	*IPNInformation
	*buyerInformation
	*shippingInformation
	*simpleButtonFields
}

func (i *IPN) ToSimpleIPN() (*simpleIPN, error) {
	if i.IPNType != "simple" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'simple'")
	}
	return &simpleIPN{
		IPNInformation:      &i.IPNInformation,
		buyerInformation:    i.buyerInformation,
		shippingInformation: i.shippingInformation,
		simpleButtonFields:  i.simpleButtonFields,
	}, nil
}

type buttonIPN struct {
	*IPNInformation
	*buyerInformation
	*shippingInformation
	*advancedButtonFields
}

func (i *IPN) ToButtonIPN() (*buttonIPN, error) {
	if i.IPNType != "button" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'button'")
	}
	return &buttonIPN{
		IPNInformation:       &i.IPNInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		advancedButtonFields: i.advancedButtonFields,
	}, nil
}

type cartIPN struct {
	*IPNInformation
	*buyerInformation
	*shippingInformation
	*shoppingCartButtonFields
}

func (i *IPN) ToCartIPN() (*cartIPN, error) {
	if i.IPNType != "cart" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'cart'")
	}
	return &cartIPN{
		IPNInformation:           &i.IPNInformation,
		buyerInformation:         i.buyerInformation,
		shippingInformation:      i.shippingInformation,
		shoppingCartButtonFields: i.shoppingCartButtonFields,
	}, nil
}

type donationIPN struct {
	*IPNInformation
	*buyerInformation
	*shippingInformation
	*donationButtonFields
}

func (i *IPN) ToDonationIPN() (*donationIPN, error) {
	if i.IPNType != "donation" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'donation'")
	}
	return &donationIPN{
		IPNInformation:       &i.IPNInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		donationButtonFields: i.donationButtonFields,
	}, nil
}

type depositIPN struct {
	*IPNInformation
	*depositInformation
}

func (i *IPN) ToDepositIPN() (*depositIPN, error) {
	if i.IPNType != "deposit" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'deposit'")
	}
	return &depositIPN{
		IPNInformation:     &i.IPNInformation,
		depositInformation: i.depositInformation,
	}, nil
}

type withdrawalIPN struct {
	*IPNInformation
	*withdrawalInformation
}

func (i *IPN) ToWithdrawalIPN() (*withdrawalIPN, error) {
	if i.IPNType != "withdrawal" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'withdrawal'")
	}
	return &withdrawalIPN{
		IPNInformation:        &i.IPNInformation,
		withdrawalInformation: i.withdrawalInformation,
	}, nil
}

type apiIPN struct {
	*IPNInformation
	*apiGeneratedTransactionFields
}

func (i *IPN) ToApiIPN() (*apiIPN, error) {
	if i.IPNType != "api" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'api'")
	}
	return &apiIPN{
		IPNInformation:                &i.IPNInformation,
		apiGeneratedTransactionFields: i.apiGeneratedTransactionFields,
	}, nil
}

func (c *Client) ParseIPN(r *http.Request) (*IPN, error) {
	hmac := r.Header.Get("HMAC")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("coinpayments: error reading request body - %v", err)
	}

	genHMAC, err := c.makeHMAC(string(data))
	if err != nil {
		return nil, fmt.Errorf("coinpayments: error generating ipn HMAC - %v", err)
	}

	if hmac != genHMAC {
		return nil, fmt.Errorf("coinpayments: could not validate server HMAC")
	}

	ipn := &IPN{}
	err = json.Unmarshal(data, ipn)
	if err != nil {
		return nil, fmt.Errorf("coinpayments: error unmarshaling ipn json - %v", err)
	}
	switch ipn.IPNType {
	case "simple":
		bi := &buyerInformation{}
		err := json.Unmarshal(data, bi)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling buyer information - %v", err)
		}
		ipn.buyerInformation = bi

		si := &shippingInformation{}
		err = json.Unmarshal(data, si)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling shipping information - %v", err)
		}
		ipn.shippingInformation = si

		sb := &simpleButtonFields{}
		err = json.Unmarshal(data, sb)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling simple button fields - %v", err)
		}
		ipn.simpleButtonFields = sb
	case "button":
		bi := &buyerInformation{}
		err := json.Unmarshal(data, bi)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling buyer information - %v", err)
		}
		ipn.buyerInformation = bi

		si := &shippingInformation{}
		err = json.Unmarshal(data, si)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling shipping information - %v", err)
		}
		ipn.shippingInformation = si

		ab := &advancedButtonFields{}
		err = json.Unmarshal(data, ab)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling advanced button fields - %v", err)
		}
		ipn.advancedButtonFields = ab
	case "cart":
		bi := &buyerInformation{}
		err := json.Unmarshal(data, bi)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling buyer information - %v", err)
		}
		ipn.buyerInformation = bi

		si := &shippingInformation{}
		err = json.Unmarshal(data, si)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling shipping information - %v", err)
		}
		ipn.shippingInformation = si

		sc := &shoppingCartButtonFields{}
		err = json.Unmarshal(data, sc)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling shopping cart button fields - %v", err)
		}
		ipn.shoppingCartButtonFields = sc
	case "donation":
		bi := &buyerInformation{}
		err := json.Unmarshal(data, bi)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling buyer information - %v", err)
		}
		ipn.buyerInformation = bi

		si := &shippingInformation{}
		err = json.Unmarshal(data, si)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling shipping information - %v", err)
		}
		ipn.shippingInformation = si

		db := &donationButtonFields{}
		err = json.Unmarshal(data, db)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling donation button fields - %v", err)
		}
		ipn.donationButtonFields = db
	case "deposit":
		di := &depositInformation{}
		err = json.Unmarshal(data, di)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling deposit information - %v", err)
		}
		ipn.depositInformation = di
	case "withdrawal":
		wi := &withdrawalInformation{}
		err = json.Unmarshal(data, wi)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling withdrawal information - %v", err)
		}
		ipn.withdrawalInformation = wi
	case "api":
		api := &apiGeneratedTransactionFields{}
		err = json.Unmarshal(data, api)
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error unmarshaling api generated transaction fields - %v", err)
		}
		ipn.apiGeneratedTransactionFields = api
	}
	return ipn, nil
}

type depositInformation struct {
	TransactionID string `json:"txn_id"`
	Address       string `json:"address"`
	DestTag       string `json:"dest_tag"`
	Status        int    `json:"status"`
	StatusText    string `json:"status_text"`
	Currency      string `json:"currency"`
	Confirms      int    `json:"confirms"`
	Amount        string `json:"amount"`
	Amounti       int    `json:"amounti"`
	Fee           string `json:"fee"`
	Feei          int    `json:"feei"`
	FiatCoin      string `json:"fiat_coin"`
	FiatAmount    string `json:"fiat_amount"`
	FiatAmounti   int    `json:"fiat_amounti"`
	FiatFee       string `json:"fiat_fee"`
	FiatFeei      int    `json:"fiat_feei"`
}

type withdrawalInformation struct {
	ID            string `json:"id"`
	Status        int    `json:"status"`
	StatusText    string `json:"status_text"`
	Address       string `json:"address"`
	TransactionID string `json:"txn_id"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	Amounti       int    `json:"amounti"`
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
	Status           int    `json:"status"`
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
	ReceivedConfirms int    `json:"received_confirms"`
}

type advancedButtonFields struct {
	Status           int    `json:"status"`
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
	ReceivedConfirms int    `json:"received_confirms"`
}

type shoppingCartButtonFields struct {
	Status           int    `json:"status"`
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
	ReceivedConfirms int    `json:"received_confirms"`
}

type donationButtonFields struct {
	Status           int    `json:"status"`
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
	ReceivedConfirms int    `json:"received_confirms"`
}

type apiGeneratedTransactionFields struct {
	Status           int    `json:"status"`
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
	ReceivedConfirms int    `json:"received_confirms"`
}
