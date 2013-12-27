package libcoin

// define the format that will be sent out from this system

type CoinPacket struct {
  Exchange      string
  Last          float64
  CurrentVolume float64
  Currency      string
}

func PanicError(err error) {
        if err != nil {
                panic(err.Error())
        }
}


