package gohanspell

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const MaxLen = 280
const PusanUnivURL = "http://speller.cs.pusan.ac.kr/results"

type Correction struct {
	Str     string `json:"str"`
	ErrInfo []struct {
		Help          string `json:"help"`
		ErrorIdx      int    `json:"errorIdx"`
		CorrectMethod int    `json:"correctMethod"`
		Start         int    `json:"start"`
		End           int    `json:"end"`
		OrgStr        string `json:"orgStr"`
		CandWord      string `json:"candWord"`
	} `json:"errInfo"`
	Idx int `json:"idx"`
}

func countByWord(s string) int {
	return len(strings.Split(s, " "))
}

func splitByWord(s string) []string {
	ret := make([]string, 0)

	sentences := strings.Split(s, "\n")
	s = ""

	for len(sentences) > 0 {
		if countByWord(s) + countByWord(sentences[0]) > MaxLen {
			s = strings.TrimSpace(s)
			ret = append(ret, s)
			s = sentences[0]
		} else {
			s += "\t" + sentences[0]
		}

		if len(sentences) == 1 {
			sentences = nil
			break
		}
		sentences = sentences[1:]
	}
	if len(s) > 0 {
		ret = append(ret, s)
	}
	return ret
}

func mergeCorrection(c Correction) string {
	rand.Seed(time.Now().Unix())
	s := c.Str
	for _, e := range c.ErrInfo {
		candWord := e.CandWord
		if strings.Count(candWord, "|") > 0 {
			words := strings.Split(candWord, "|")
			candWord = words[rand.Intn(len(words))]
		}
		s = strings.ReplaceAll(s, e.OrgStr, candWord)
	}
	return s
}

func PostPusanUniv(s string) (string, error) {
	txt := ""
	for _, sentence := range splitByWord(s) {
		body := url.Values{}
		body.Set("text1", sentence)
		resp, err := http.PostForm(PusanUnivURL, body)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			return "", errors.New("network error")
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			raw := strings.Split(strings.Split(string(data), "data = [{")[1], "}];")[0]
			raw = "{" + raw + "}"
			correction := Correction{}
			err = json.Unmarshal([]byte(raw), &correction)
			if err != nil {
				return "", err
			}
			txt += "\n" + strings.ReplaceAll(mergeCorrection(correction), "\t", "\n")
		}

		resp.Body.Close()
	}

	txt = strings.TrimSpace(txt)
	return txt, nil
}
