package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func Internet(Time int64) int64 {
	rand.Seed(Time)
	return getConnectionTime(getRandomUrl(rand.Int63()))
}

func getConnectionTime(url string) int64 {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}
	fmt.Println(resp)

	return time.Now().UnixMicro() - startTime.UnixMicro()
}
func getRandomUrl(time int64) string {
	rand.Seed(time)
	urlArray := []string{"facebook.com",
		"google.com",
		"sap.com",
		"bing.com",
		"cnn.com"}
	index := rand.Intn(len(urlArray) - 1)
	return urlArray[index]
}
