package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
{"high": "659.00", "last": "644.00", "timestamp": "1387944176", "bid": "643.88", "volume": "10136.54430222", "low": "628.11", "ask": "644.00"}
*/

type Exchange struct {
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Timestamp int64   `json:"timestamp,string"`
	Bid       float64 `json:"bid,string"`
	Volume    float64 `json:"volume,string"`
	Low       float64 `json:"low,string"`
	Ask       float64 `json:"ask,string"`
}

type CoinPacket struct {
	Exchange      string
	Last          float64
	CurrentVolume float64
	Currency      string
}

func ticker() {
	url := "https://www.bitstamp.net/api/ticker/"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	info := &Exchange{}
	json.Unmarshal(body, &info)

	fmt.Printf("coin info: %+v \n", info)

	cp := &CoinPacket{Exchange: "bitstamp", Last: info.Last, CurrentVolume: info.Volume, Currency: "usd"}
	fmt.Printf("cp: %+v \n", cp)
}

func main() {
	ticker()
}
