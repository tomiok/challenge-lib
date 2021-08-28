package challengelib

import (
	"bytes"
	"github.com/buger/jsonparser"
	"log"
	"net/http"
	"net/url"
	"strings"
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

	var sb strings.Builder
	_, err = jsonparser.ArrayEach(buf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ct, _ := jsonparser.GetString(value, "challenge_type")
		desc, _ := jsonparser.GetString(value, "description")
		resLevel, _ := jsonparser.GetString(value, "level")

		sb.WriteString(ct)
		sb.WriteString(", ")
		sb.WriteString(desc)
		sb.WriteString(", ")
		sb.WriteString(resLevel)
	})

	if err != nil {
		log.Println(err.Error())
		return "cannot get any challenge :(, body perser error"
	}

	return sb.String()
}

func buildRequest(level, challenge string) *http.Request {
	u, _ := url.Parse(baseURL)
	req := &http.Request{
		URL:    u,
		Method: http.MethodGet,
	}

	q := req.URL.Query()

	q.Add(level, challenge)
	req.URL.RawQuery = q.Encode()
	return req
}
