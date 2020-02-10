package lib

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
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
	get, err := ac.client.Get(BASE_URL + "/login")
	if err != nil {
		return err
	}
	defer get.Body.Close()

	token, err := ParseCSRFToken(get.Body)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("csrf_token", token)
	values.Set("username", name)
	values.Set("password", password)

	post, err := ac.client.Post(BASE_URL+"/login?"+values.Encode(), "", nil)
	if err != nil {
		return err
	}
	defer post.Body.Close()

	return nil
}

func (ac *AtCoder) FetchContest(contest string) (*Contest, error) {
	contestURL := BASE_URL + "/contests/" + contest
	res, err := ac.client.Get(contestURL + "/tasks")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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
		defer tres.Body.Close()

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

func (ac *AtCoder) Submit(contestId, taskId, languageId string, code io.Reader) error {
	submitUrl := BASE_URL + "/contests/" + contestId + "/submit"
	get, err := ac.client.Get(submitUrl)
	if err != nil {
		return err
	}
	defer get.Body.Close()

	token, err := ParseCSRFToken(get.Body)
	if err != nil {
		return err
	}

	all, err := ioutil.ReadAll(code)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("data.TaskScreenName", taskId)
	values.Set("csrf_token", token)
	values.Set("data.LanguageId", languageId)
	values.Set("sourceCode", string(all))

	post, err := ac.client.Post(submitUrl+"?"+values.Encode(), "", nil)
	if err != nil {
		return err
	}
	defer post.Body.Close()

	if post.StatusCode != 200 {
		return fmt.Errorf("sumit failed")
	}

	return nil
}