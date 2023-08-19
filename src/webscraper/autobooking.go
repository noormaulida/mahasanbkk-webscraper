package webscraper

import (
	// "log"
	"net/http"
	"net/url"
    "strings"

	"mahasanbkk-webscraper/pkg/config"
)

func AutoBooking(tableId string) *http.Response {
    client := http.Client{}

    apiUrl := config.ConfigData.MahasanUrl + config.ConfigData.MahasanBookUrl
    
    form := url.Values{}
    form.Add("table_id", tableId)
    form.Add("fname", config.ConfigData.UserFName)
    form.Add("lname", config.ConfigData.UserLName)
    form.Add("email", config.ConfigData.UserEmail)
    form.Add("phone", config.ConfigData.UserPhone)
    form.Add("message", config.ConfigData.UserMessage)
    form.Add("submit", "")
    
    req, _ := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(form.Encode()))

    resp, _ := client.Do(req)
    return resp

}