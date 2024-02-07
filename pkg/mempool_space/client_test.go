package mempool_space

import (
	"context"
	"fmt"
	"testing"
)

func Test_Price(t *testing.T) {
	c := NewClient("https://mempool.space")
	res, err := c.Price(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
func Test_RecommendedFees(t *testing.T) {
	c := NewClient("https://mempool.space")
	res, err := c.RecommendedFees(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
