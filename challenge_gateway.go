package challengelib

import (
	"bytes"
	"github.com/buger/jsonparser"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://challenbot.herokuapp.com/challenges?level=%s&type=%s"

var client = &http.Client{}

func FindChallenge(level, challengeType string) string {
	req := buildRequest(level, challengeType)
	res, err := client.Do(req)

	if err != nil {
		log.Println(err.Error())
		return "cannot get any challenge :("
	}

	body := res.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)
	b := buf.Bytes()

	desc, err := jsonparser.GetString(b, "description")

	if err != nil {
		log.Println(err.Error())
		return "cannot get any challenge :(, body parser error"
	}

	return desc
}

func buildRequest(level, challenge string) *http.Request {
	u, _ := url.Parse(baseURL)
	req := &http.Request{
		URL:    u,
		Method: http.MethodGet,
	}

	q := req.URL.Query()

	q.Add("level", level)
	q.Add("type", challenge)
	req.URL.RawQuery = q.Encode()
	return req
}
