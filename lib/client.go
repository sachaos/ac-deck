package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type AtCoder struct {
	client *http.Client
}

func NewAtCoder() (*AtCoder, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &AtCoder{client: &http.Client{Jar: jar}}, nil
}

func NewAtCoderWithClient(client *http.Client) *AtCoder {
	return &AtCoder{client: client}
}

func (ac *AtCoder) Login(name, password string) error {
	token, err := ac.getCSRFToken()
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("csrf_token", token)
	values.Add("username", name)
	values.Add("password", password)

	_, err = ac.client.Post("https://atcoder.jp/login?"+values.Encode(), "", nil)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AtCoder) getCSRFToken() (string, error) {
	get, err := ac.client.Get("https://atcoder.jp/login")
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(get.Body)
	if err != nil {
		return "", err
	}

	tokenEl := doc.Find("#main-container > div.row > div > form > input[name=csrf_token]")
	token, exists := tokenEl.Attr("value")
	if !exists {
		return "", fmt.Errorf("value not found")
	}

	return token, nil
}
