package main

import (
	"log"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)


func getWeather() (string, error) {
	url := "https://weixin.jirengu.com/weather/future24h?cityid=WX4FBXXFKE4F"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func queryWeatherToUser(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm();
	// fmt.Println(r.Form)
	
	data, err := getWeather()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, data)
}

func main() {
	http.HandleFunc("/getweather",queryWeatherToUser)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}