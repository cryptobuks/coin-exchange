package bitstamp

import (
	"encoding/json"
	"fmt"
	"github.com/blooberr/coin-exchange/libcoin"
	"io/ioutil"
	"net/http"
	"time"
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

func GetTicker() []byte {
	url := "https://www.bitstamp.net/api/ticker/"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	info := &Exchange{}
	json.Unmarshal(body, &info)

	//fmt.Printf("coin info: %+v \n", info)

	cp := &libcoin.CoinPacket{Exchange: "bitstamp", Last: info.Last, CurrentVolume: info.Volume, Currency: "usd"}
	//fmt.Printf("cp: %+v \n", cp)
	b, _ := json.Marshal(cp)
	return b
}

func Loop(interval int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))

	for t := range ticker.C {
		fmt.Printf("[%s] [bitstamp]: %s \n", t, GetTicker())
	}

}
