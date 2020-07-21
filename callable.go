package coinpayments

import "net/url"

type callable interface {
	command() string
	values() url.Values
}
