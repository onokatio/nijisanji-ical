package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/arran4/golang-ical"
	"time"
	"strconv"
)

type Response struct {
	Status string `json:"status"`
	Data struct {
		Events []struct {
			Id int `json:"id"`
			Name string `json:"name"`
			Description string `json:"description"`
			Public float64 `json:"public"`
			Url string `json:"url"`
			Thumbnail string `json:"thunmbnail"`
			StartDate string `json:"start_date"`
			EndDate string `json:"end_date"`
			Recommend bool `json:"recommend"`
			Genre struct {
				Id float64 `json:"id"`
				Name string `json:"name"`
			}`json:"genre"`
			Livers struct {
				Id float64 `json:"id"`
				Name string `json:"name"`
				Avatar string `json:"avatar"`
				Color string `json:"color"`
			}`json:"genre"`
		}`json:"events"`
	}`json:"data"`
}

func main(){
	resp, err := http.Get("https://api.itsukaralink.jp/v1.2/events.json")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	schedule := Response{}
	err = json.NewDecoder(resp.Body).Decode(&schedule)
	if err != nil {
		fmt.Println(err)
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetProductId("-//onokatio//nijisanji-ics")
	cal.SetXWRCalName("にじさんじ 配信スケジュール")

	for _, value := range schedule.Data.Events {
		StartDate, err := time.Parse(time.RFC3339Nano,value.StartDate)
		EndDate, err := time.Parse(time.RFC3339Nano,value.EndDate)
		Id := strconv.Itoa(value.Id)
		if err != nil {
			fmt.Println(err)
		}

		event := cal.AddEvent(Id)
		event.SetSummary(value.Name)
		event.SetStartAt(StartDate)
		event.SetEndAt(EndDate)
		event.SetDescription(value.Url + "\n\n" + value.Description)
		event.SetURL(value.Url)
	}

	fmt.Println(cal.Serialize())
}
