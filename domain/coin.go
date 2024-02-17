package domain

const SatoshiInBtc = 100_000_000

func SatoshiToBtc(satoshi float64) float64 {
	return satoshi / SatoshiInBtc
}

func BtcToSatoshi(coin float64) float64 {
	return coin * SatoshiInBtc
}

const (
	AddressFromBtcByte = 148
	AddressToBtcByte   = 34
	TransactionBtcByte = 10
)

var MinTransactionByte = AddressFromBtcByte + AddressToBtcByte + TransactionBtcByte
