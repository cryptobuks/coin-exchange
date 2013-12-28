package btce

import (
	"encoding/json"
	"fmt"
	"github.com/blooberr/coin-exchange/libcoin"
	"io/ioutil"
	"net/http"
	"time"
)

type exchange struct {
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

func GetTicker() []byte {
	url := "https://btc-e.com/api/2/btc_usd/ticker"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)

	info := &exchange{}
	json.Unmarshal(body, &info)

	//fmt.Printf("coin info: %+v \n", info)

	cp := &libcoin.CoinPacket{Exchange: "btc-e", Last: info.Ticker.Last, CurrentVolume: info.Ticker.VolCur, Currency: "usd"}
	//fmt.Printf("cp: %+v \n", cp)

	b, _ := json.Marshal(cp)
	return b
}

func Loop(interval int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))

	for t := range ticker.C {
		fmt.Printf("[%s] [btc-e]: %s \n", t, GetTicker())
	}

}
