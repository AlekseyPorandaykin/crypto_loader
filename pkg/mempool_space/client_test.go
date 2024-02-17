package mempool_space

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"testing"
)

func Test_Price(t *testing.T) {
	c := NewClient("https://mempool.space")
	res, err := c.Price(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	//44241
	fmt.Println(res)
}
func Test_RecommendedFees(t *testing.T) {
	c := NewClient("https://mempool.space")
	res, err := c.RecommendedFees(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	satoshiByteFee := res.FastestFee * float64(domain.MinTransactionByte)
	btcByteFee := domain.SatoshiToBtc(satoshiByteFee)
	usdtFee := btcByteFee * 44241
	fmt.Println(usdtFee)
	fmt.Println(res)
}
