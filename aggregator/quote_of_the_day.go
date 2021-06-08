package aggregator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QUTDResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote    string `json:"quote"`
			Author   string `json:"author"`
			Language string `json:"language"`
		} `json:"quotes"`
	} `json:"contents"`
}

func FetchQuoteOfTheDay() (Module, error) {
	var module Module

	url := "https://quotes.rest/qod.json"
	resp, err := http.Get(url)
	if err != nil {
		return module, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return module, err
	}

	var r QUTDResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return module, err
	}

	module.Caption = "Quote of the day"
	module.Text = r.Contents.Quotes[0].Quote + "\n\n- " + r.Contents.Quotes[0].Author

	return module, nil
}
