package coinpayments

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type errResponse struct {
	Error string `json:"error"`
}

type clientOptions func(client *Client)

type Client struct {
	client     *http.Client
	privateKey string
	publicKey  string
	ipnSecret  string
}

func NewClient(publicKey, privateKey string, options ...clientOptions) *Client {
	client := &Client{
		privateKey: privateKey,
		publicKey:  publicKey,
		client:     http.DefaultClient,
	}

	for _, o := range options {
		o(client)
	}
	return client
}

func (c *Client) call(callable Callable, response interface{}) error {
	data := callable.values()

	data.Add("key", c.publicKey)
	data.Add("version", apiVersion)
	data.Add("cmd", callable.command())
	data.Add("format", apiFormat)

	sData := data.Encode()

	dataHMAC, err := c.makeHMAC(sData)
	if err != nil {
		return fmt.Errorf("coinpayments: error making HMAC - %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(sData))
	if err != nil {
		return fmt.Errorf("coinpayments: error making api request - %v", err)
	}

	req.Header.Add("HMAC", dataHMAC)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(sData)))

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("coinpayments: error doing api request - %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("coinpayments: api call returned unexpected status: %v", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("coinpayments: error reading api response body - %v", err)
	}

	errResp := &errResponse{}
	if err := json.Unmarshal(body, &errResp); err != nil {
		return fmt.Errorf("coinpayments: error unmarshaling api error response - %v", err)
	}

	if errResp.Error != apiSuccess {
		return fmt.Errorf("coinpayments: api error - %v", errResp.Error)
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("coinpayments: error unmarshaling response json - %v", err)
	}

	return nil
}

func (c *Client) makeHMAC(data string) (string, error) {
	hash := hmac.New(sha512.New, []byte(c.privateKey))
	if _, err := hash.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
