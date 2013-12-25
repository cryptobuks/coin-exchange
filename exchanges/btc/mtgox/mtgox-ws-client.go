package main

import "code.google.com/p/go.net/websocket"
import (
	"encoding/json"
	"fmt"
)

type CoinPacket struct {
  Exchange      string
  Last          float64
  CurrentVolume float64
  Currency      string
}

// interesting note with parsing - have to define the top fields before we can
//  get to vol.  Is this just a golang quirk?
type ticker struct {
	High high `json:"high"`
	Low  low  `json:"low"`
	Avg  avg  `json:"avg"`
	Vwap vwap `json:"vwap"`
	Vol  vol  `json:"vol"`
  LastLocal map[string]string `json:"last_local"`
  LastOrig map[string]string `json:"last_orig"`
  LastAll map[string]string `json:"last_all"`
  Last last `json:"last"`
  Buy map[string] string `json:"buy"`
  Sell map[string]string `json:"sell"`
  Item string `json:"item"`
  Now int64 `json:"now,string"` // timestamp in microseconds
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
	Value        float64 `json:"value,string"`
	ValueInt     int64   `json:"value_int,string"`
	Display      string  `json:"display"`
	DisplayShort string  `json:"display_short"`
	Currency     string  `json:"currency"`
}

type last struct {
  Value float64 `json:"value,string"`
  ValueInt int64 `json:"value_int,string"`
  Display      string  `json:"display"`
  DisplayShort string  `json:"display_short"`
  Currency     string  `json:"currency"`
}

type Exchange struct {
	Channel     string `json:"channel"`
	ChannelName string `json:"channel_name"`
	Op          string `json:"op"`
	Origin      string `json:"origin"`
	Private     string `json:"private"`
	Ticker      ticker `json:"ticker"`
	Stamp       int64  `json:"stamp"`
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

  cp := &CoinPacket{}

	var resp = make([]byte, 4096)
	info := &Exchange{}
	for {
		n, err := ws.Read(resp)
		panic_error(err)

		// fmt.Printf("Received: %s\n", string(resp[0:n]))
		json.Unmarshal(resp[0:n], &info)

		if (info.ChannelName == "ticker.BTCUSD") {
			//fmt.Printf("Received: %s \n", string(resp[0:n]))
			//fmt.Printf("%+v \n", info)
     //fmt.Printf("last: %f (%s) with volume: %f \n", info.Ticker.Last.Value, info.Ticker.Last.Currency, info.Ticker.Vol.Value)

      fmt.Printf("timestamp: %d \n", info.Ticker.Now)
      cp.Exchange = "mtgox"
      cp.Last = info.Ticker.Last.Value
      cp.Currency = info.Ticker.Last.Currency
      cp.CurrentVolume = info.Ticker.Vol.Value
      fmt.Printf("cp: %+v \n", cp)
		}
	}
}
