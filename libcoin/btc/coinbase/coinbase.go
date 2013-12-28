package coinbase

import (
	"encoding/json"
	"fmt"
	"github.com/blooberr/coin-exchange/libcoin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetTicker() []byte {
	// tells you the current price of a bitcoin.  Usually between buy and sell price.
	url := "https://coinbase.com/api/v1/prices/spot_rate"
	res, err := http.Get(url)
	libcoin.PanicError(err)

	body, err := ioutil.ReadAll(res.Body)
	libcoin.PanicError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	libcoin.PanicError(err)

	//fmt.Printf("spot_price is %s \n", data["amount"])
	amount, _ := strconv.ParseFloat(data["amount"].(string), 64)

	cp := &libcoin.CoinPacket{Exchange: "coinbase", Last: amount, CurrentVolume: -1, Currency: "usd"}
	b, _ := json.Marshal(cp)
	return b
}

func Loop(interval int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(10000))

	for t := range ticker.C {
		fmt.Printf("[%s] [coinbase]: %s \n", t, GetTicker())
	}

}
