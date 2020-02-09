package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
)

func init()  {
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
}

const BASE_URL = "https://atcoder.jp"

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

	_, err = ac.client.Post(BASE_URL+"/login?"+values.Encode(), "", nil)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AtCoder) getCSRFToken() (string, error) {
	get, err := ac.client.Get(BASE_URL+"/login")
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

func (ac *AtCoder) FetchContest(contest string) (*Contest, error) {
	contestURL := BASE_URL + "/contests/" + contest
	res, err := ac.client.Get(contestURL + "/tasks")
	if err != nil {
		return nil, err
	}

	paths, err := ParseTasksPage(res.Body)
	if err != nil {
		return nil, err
	}

	tasks := make([]*Task, len(paths))
	for i, p := range paths {
		tres, err := ac.client.Get(BASE_URL + p)
		if err != nil {
			return nil, err
		}

		task, err := ParseTaskPage(tres.Body)
		if err != nil {
			return nil, err
		}
		task.ID = path.Base(p)
		tasks[i] = task
	}

	return &Contest{
		ID:    contest,
		URL:   contestURL,
		Tasks: tasks,
	}, nil
}
