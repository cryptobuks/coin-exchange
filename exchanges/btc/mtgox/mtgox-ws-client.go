package main

import "code.google.com/p/go.net/websocket"
import (
	"encoding/json"
	"fmt"
)


// interesting note with parsing - have to define the top fields before we can 
/  get to vol.  Is this just a golang quirk?
type ticker struct {
	High high `json:"high"`
  Low low `json:"low"`
  Avg avg `json:"avg"`
  Vwap vwap `json:"vwap"`
	Vol  vol  `json:"vol"`
}

type high struct {
	Value string `json:"value"`
}

type low struct {
  Value float64 `json:"value,string"`
}

type vwap struct {
  Value float64 `json:"value,string"`
}

type avg struct {
  Value float64 `json:"value,string"`
}

type vol struct {
	Value    float64 `json:"value,string"`
	ValueInt int64   `json:"value_int,string"`
	Currency string  `json:"currency"`
}

type Exchange struct {
	Channel     string `json:"channel"`
	ChannelName string `json:"channel_name"`
	Op          string `json:"op"`
	Origin      string `json:"origin"`
	Private     string `json:"private"`
	Ticker      ticker `json:"ticker"`
}

func panic_error(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// reference - https://en.bitcoin.it/wiki/MtGox/API/Streaming

func main() {
	origin := "http://localhost"
	ws, err := websocket.Dial("ws://websocket.mtgox.com:80/mtgox?Currency=USD", "", origin)
	panic_error(err)

	var resp = make([]byte, 4096)
	info := &Exchange{}

	for {
		n, err := ws.Read(resp)
		panic_error(err)

		// fmt.Printf("Received: %s\n", string(resp[0:n]))
		json.Unmarshal(resp[0:n], &info)

		if info.ChannelName == "ticker.BTCUSD" {
			fmt.Printf("Received: %s \n", string(resp[0:n]))
			fmt.Printf("%+v \n", info)
		}
	}
}
