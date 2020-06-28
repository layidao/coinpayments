package coinpayments

import "net/url"

type Callable interface {
	command() string
	values() url.Values
}
