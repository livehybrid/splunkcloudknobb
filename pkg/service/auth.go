package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// AuthCookies holds the Splunk Cloud cookies required to interact
type AuthCookies struct {
	URL           string
	Username      string
	Password      string
	SplunkWebUUID string // experience_id
	Cvalue        string
	SessionID     string
	Splunkd       string
	SplunkWebCSRF string // token_key
	Expiry        string
	DTS           string
	CookiePort    string
}

func (a AuthCookies) String() string {
	b := a
	b.Password = "[unprotected]"
	b.SplunkWebCSRF = "[unprotected]"
	b.Splunkd = "[unprotected]"
	b.SplunkWebUUID = "[unprotected]"
	spew.Config.DisableMethods = true
	s := spew.Sdump(b)
	return s
}

func (s *Service) authCookieDecorate(auth AuthCookies, client *http.Client, req *http.Request) {

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1) //This date is ignored server side (apparently)

	cookies = append(cookies, &http.Cookie{Name: "session_id_" + auth.CookiePort, Value: auth.SessionID, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "experience_id", Value: auth.SplunkWebUUID, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_uid", Value: auth.SplunkWebUUID, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkd_" + auth.CookiePort, Value: auth.Splunkd, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_csrf_token_" + auth.CookiePort, Value: auth.SplunkWebCSRF, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "token_key", Value: auth.SplunkWebCSRF, Expires: expire})

	client.Jar.SetCookies(req.URL, cookies)
}

// Login takes a username and password, executed 4 steps to login, and returns auth cookies
func (s *Service) Login(u string, username string, password string, cookiePort string, log *log.Logger) (AuthCookies, error) {
	var authd AuthCookies

	authd.DTS = time.Now().Format("20060102T150405")

	log.Debugf("Splunk Login Attempt: %s", u)
	authd.URL = u
	authd.Username = username
	authd.Password = password
	authd.CookiePort = cookiePort

	// splunkuuid == experience_id also
	cval, splunkwebuuid, err := s.step1(u)
	if err != nil || cval == "" || splunkwebuuid == "" {
		err = fmt.Errorf("failed Authentication Step 1: invalid url, username, or password: values ('%s', '%s')", u, username)
		return AuthCookies{}, err
	}
	authd.SplunkWebUUID = splunkwebuuid
	authd.Cvalue = cval

	tsession, err := s.step2(u, authd.CookiePort)
	if err != nil || tsession == "" {
		fmt.Printf("error: %v, tempsess: %s\n", err, tsession)
		err = errors.New("failed on auth Step 2")
		return AuthCookies{}, err
	}

	sessionid, err := s.step3(u, tsession, authd.CookiePort)
	if err != nil || sessionid == "" {
		err = errors.New("failed on auth Step 3")
		fmt.Printf("error: %v, session_id_%s: %s\n", err, authd.CookiePort, sessionid)
		return AuthCookies{}, err
	}
	authd.SessionID = sessionid

	//splunkwebcsrf also known as token_key
	splunkd, splunkwebcsrf, expiry, err := s.step4(u, cval, splunkwebuuid, sessionid, username, password, authd.CookiePort)
	if err != nil || splunkd == "" || splunkwebcsrf == "" {
		err = errors.New("failed Authentication Step 4: likey wrong username or password")
		fmt.Printf("error: %v, splunkd_%s: %s, splunkweb_csrf_token_%s:%s\n", err, authd.CookiePort, splunkd, authd.CookiePort, splunkwebcsrf)
		return AuthCookies{}, err
	}
	authd.Splunkd = splunkd
	authd.SplunkWebCSRF = splunkwebcsrf
	authd.Expiry = expiry

	log.Debugf("Success: AuthorizedToken: %s", authd)

	s.SessionMap["Username"] = authd.Username
	s.SessionMap["SplunkWebCSRF"] = authd.SplunkWebCSRF
	return authd, nil
}

func (s *Service) step1(u string) (string, string, error) {
	var url = u + "account/login?return_to=/en-US/&loginType=splunk"

	client, req, err := s.authGetRequest(url)
	if err != nil {
		return "", "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", err
	}

	var cval string
	var splunkwebuid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "cval":
			cval = v.Value
		case "splunkweb_uid":
			splunkwebuid = v.Value
		}
	}
	return cval, splunkwebuid, nil
}

func (s *Service) step2(u string, cookiePort string) (string, error) {
	var url = u + `config?autoload=1`

	client, req, err := s.authGetRequest(url)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var sessionid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "session_id_" + cookiePort:
			sessionid = v.Value
		}
	}

	return sessionid, nil
}

func (s *Service) step3(u string, tsess string, cookiePort string) (string, error) {
	var url = u + `config`

	client, req, err := s.authGetRequest(url)
	if err != nil {
		return "", err
	}

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1)
	cookies = append(cookies, &http.Cookie{Name: "session_id_" + cookiePort, Value: tsess, Expires: expire})
	client.Jar.SetCookies(req.URL, cookies)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var sessionid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "session_id_" + cookiePort:
			sessionid = v.Value
		}
	}

	return sessionid, nil
}

func (s *Service) step4(u, cval, splunkuuid, sessionid, username, password, cookiePort string) (splunkd string, csrf string, expiry string, err error) {
	var url = u + `account/login`

	returnto := "/en-US/"

	body := fmt.Sprintf("cval=%s&username=%s&password=%s&return_to=%s", cval, username, password, returnto)

	client, req, err := s.authPostRequest(url, strings.NewReader(body))

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1)
	cookies = append(cookies, &http.Cookie{Name: "cval", Value: cval, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_uid", Value: splunkuuid, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "session_id_" + cookiePort, Value: sessionid, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_uid", Value: splunkuuid, Expires: expire})
	client.Jar.SetCookies(req.URL, cookies)

	if err != nil {
		return "", "", "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", err
	}

	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "splunkd_" + cookiePort:
			splunkd = v.Value
			expiry = v.Expires.String()
		case "splunkweb_csrf_token_" + cookiePort:
			csrf = v.Value
		}
	}

	return splunkd, csrf, expiry, nil
}
