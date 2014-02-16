package main

import (
	"fmt"
	"github.com/blooberr/coin-exchange/libcoin/btc/bitstamp"
	"github.com/blooberr/coin-exchange/libcoin/btc/btc-e"
	"github.com/blooberr/coin-exchange/libcoin/btc/campbx"
	"github.com/blooberr/coin-exchange/libcoin/btc/coinbase"
	"github.com/blooberr/coin-exchange/libcoin/btc/localbitcoins"
	"github.com/blooberr/coin-exchange/libcoin/btc/mtgox"
	"github.com/blooberr/coin-exchange/libcoin/btc/vircurex"
	//  "time"
)

func main() {
	fmt.Printf("connecting to exchanges.. \n")
	fmt.Printf("feed running on: \n")

	//v := btce.GetTicker()
	//v := libcoin.CoinPacket{ Exchange: "hello" }
	//fmt.Printf("btce -> %s\n", string(v))

	go func() {
		btce.Loop(5000)
	}()

	go func() {
		bitstamp.Loop(10000)
	}()

	go func() {
		coinbase.Loop(3000)
	}()

	go func() {
		vircurex.Loop(8000)
	}()

	go func() {
		campbx.Loop(2500)
	}()

	go func() {
		localbitcoins.Loop(10000)
	}()

	//go func() {
	mtgox.Loop()
	//}()

}
