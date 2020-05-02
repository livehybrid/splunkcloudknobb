package service

import (
	"fmt"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"net/http"
	"net/http/cookiejar"
	"time"

	"gopkg.in/matryer/try.v1"
)

// Service holds a transaport and log
type Service struct {
	Log       *log.Logger
	Transport *http.Transport
}

//NewService creates Service instance
func NewService(log *log.Logger) (s *Service) {
	s = new(Service)
	s.Log = log
	s.Transport = &http.Transport{
		MaxIdleConns:        20,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return
}

func (s *Service) decorateUniversalHeaders(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	return
}

func (s *Service) authGetRequest(url string) (client *http.Client, req *http.Request, err error) {
	client = &http.Client{
		Transport: s.Transport,
		Timeout:   30 * time.Second,
	}

	client.Jar, _ = cookiejar.New(nil)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	s.decorateUniversalHeaders(req)

	return client, req, err
}

func (s *Service) authPostRequest(url string, data io.Reader) (client *http.Client, req *http.Request, err error) {
	client = &http.Client{
		Transport: s.Transport,
		Timeout:   30 * time.Second,
	}

	client.Jar, _ = cookiejar.New(nil)

	req, err = http.NewRequest("POST", url, data)
	if err != nil {
		return nil, nil, err
	}

	s.decorateUniversalHeaders(req)

	return client, req, err
}

var retryIntervals = []int{0, 500, 500, 500, 500, 1000, 1000, 1000, 1000, 1000, 3000}

func (s *Service) sleepBeforeRetry(attempt int) (shouldReRun bool) {
	if attempt < len(retryIntervals) {
		time.Sleep(time.Duration(retryIntervals[attempt]) * time.Millisecond)
		shouldReRun = true
	}
	return
}

func (s *Service) retryRequest(label string, client *http.Client, req *http.Request) ([]byte, error) {
	var body []byte

	err := try.Do(func(attempt int) (bool, error) {
		resp, err1 := client.Do(req)
		if err1 != nil {
			log.Println(fmt.Printf("Failed here: %v, %s", err1, err1))
			return s.sleepBeforeRetry(attempt), err1
		}
		defer resp.Body.Close()

		//TODO: Make this WAY WAY more robust.
		if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
			err2 := fmt.Errorf("failed to get %s: http_resp: %d", label, resp.StatusCode)
			return s.sleepBeforeRetry(attempt), err2
		}

		respBody := resp.Body

		var err2 error
		body, err2 = ioutil.ReadAll(respBody)
		if err2 != nil {
			err2 = fmt.Errorf("failed to read body contents for %s: %v", label, err2)
			return s.sleepBeforeRetry(attempt), err2
		}

		return false, nil
	})

	return body, err
}
