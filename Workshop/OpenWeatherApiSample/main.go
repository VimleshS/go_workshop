//http://json2struct.mervine.net/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyJsonName struct {
	Base   string `json:"base"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Cod   int `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int `json:"dt"`
	ID   int `json:"id"`
	Main struct {
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Rain struct {
		ThreeH int `json:"3h"`
	} `json:"rain"`
	Sys struct {
		Country string  `json:"country"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
		Type    int     `json:"type"`
	} `json:"sys"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int    `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   int     `json:"deg"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	//User your API_ID
	getstr := "http://history.openweathermap.org/data/2.5/history/city?q=London,UK&appid=1d748450ce1db5383ad66a792ee069b2"
	res, err := http.Get(getstr)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	openWeatherMapResponse := &OpenWeatherMap{}
	err = json.NewDecoder(res.Body).Decode(openWeatherMapResponse)

	/*
		openWeatherMapResponse := &MyJsonName{}
		err = json.NewDecoder(res.Body).Decode(openWeatherMapResponse)
	*/

	if err != nil {
		fmt.Println("Error in decoding..")
	}
	fmt.Printf("%v \n", openWeatherMapResponse)
	fmt.Printf("%v \n", openWeatherMapResponse.Name)
}
