package vircurex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//  "strconv"
	"github.com/blooberr/coin-exchange/libcoin"
)

// https://vircurex.com/welcome/api?locale=en

type exchange struct {
	Base       string  `json:"base"`
	Alt        string  `json:"alt"`
	LowestAsk  float64 `json:"lowest_ask,string"`
	HighestBid float64 `json:"highest_bid,string"`
	LastTrade  float64 `json:"last_trade,string"`
	Volume     float64 `json:"volume,string"`
}

func GetTicker() []byte {
	url := "https://vircurex.com/api/get_info_for_1_currency.json?base=BTC&alt=USD"
	res, err := http.Get(url)
	libcoin.PanicError(err)

	body, err := ioutil.ReadAll(res.Body)
	libcoin.PanicError(err)

	data := &exchange{}
	err = json.Unmarshal(body, &data)
	libcoin.PanicError(err)

	fmt.Printf("data: %s\n", data)

	cp := &libcoin.CoinPacket{Exchange: "vircurex", Last: data.LastTrade, CurrentVolume: data.Volume, Currency: "usd"}
	b, _ := json.Marshal(cp)
	return b
}

func Loop(interval int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))

	for t := range ticker.C {
		fmt.Printf("[%s] [vircurex]: %s \n", t, GetTicker())
	}

}
