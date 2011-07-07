package main

import (
	"fmt"
	"http"
	"json"
	"os"
)

var (
	bitjson   = "http://api.bitly.com/v3/shorten?login=%s&apiKey=%s&longUrl=%s&format=json"
	bitapikey string
	bituser   string
)

type BitlyData struct {
	Url        string
	Hash       string
	GlobalHash string
	LongUrl    string
	NewHash    int
}

type BitlyStatus struct {
	StatusCode int
	Data       BitlyData
}

func SetUser(user string) {
	bituser = user
}

func SetKey(key string) {
	bitapikey = key
}

func Shorten(longurl string) (string, os.Error) {
	if bituser == "" || bitapikey == "" {
		return "", fmt.Errorf("use SetUser and Setkey to set your bitly user and apikey")
	}
	longurl = http.URLEscape(longurl)
	apiurl := fmt.Sprintf(bitjson, bituser, bitapikey, longurl)
	res, err := client.Get(apiurl)
	rerr := checkResponse(res, err)
	bitly := new(BitlyStatus)

	if rerr != nil {
		return "", rerr
	}

	err = json.NewDecoder(res.Body).Decode(bitly)
	if err != nil {
		return "", nil
	}

	if bitly.StatusCode != 0 {
		return "", fmt.Errorf("Bitly Error %#v", bitly)
	}
	return bitly.Data.Url, nil
}

func checkResponse(res *http.Response, err os.Error) os.Error {
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP status %s", res.Status)
	}
	return nil
}
