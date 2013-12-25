package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type Exchange struct {
	Ticker ticker `json:"ticker"`
}

type ticker struct {
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Avg        float64 `json:"avg"`
	Vol        float64 `json:"vol"`
	VolCur     float64 `json:"vol_cur"`
	Last       float64 `json:"last"`
	Buy        float64 `json:"buy"`
	Sell       float64 `json:"sell"`
	Updated    int64   `json:"updated"`
	ServerTime int64   `json:"server_time"`
}

type CoinPacket struct {
	Exchange      string
	Last          float64
	CurrentVolume float64
	Currency      string
}

func get_ticker() {
	url := "https://btc-e.com/api/2/btc_usd/ticker"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)

	info := &Exchange{}
	json.Unmarshal(body, &info)

	fmt.Printf("coin info: %+v \n", info)

	cp := &CoinPacket{Exchange: "btc-e", Last: info.Ticker.Last, CurrentVolume: info.Ticker.VolCur, Currency: "usd"}
  fmt.Printf("cp: %+v \n", cp)
}

func main() {
	get_ticker()
}
