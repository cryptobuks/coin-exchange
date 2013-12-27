package main

import(
  "fmt"
  "github.com/blooberr/coin-exchange/libcoin/btc/btc-e"
  "github.com/blooberr/coin-exchange/libcoin/btc/bitstamp"
  "github.com/blooberr/coin-exchange/libcoin/btc/mtgox"
)

func main() {
  fmt.Printf("connecting to exchanges.. \n")
  fmt.Printf("feed running on: \n")

  //v := btce.GetTicker()
  //v := libcoin.CoinPacket{ Exchange: "hello" }
  //fmt.Printf("btce -> %s\n", string(v))

  go func() {
    btce.Loop()
  }()

  go func() {
    bitstamp.Loop()
  }()

  mtgox.Loop()
}

