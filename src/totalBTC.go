package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const BtcInterval = 21 //单位为万，比特币奖励衰减单位
const InitReward = 50.0 //单位个，比特币初始挖矿奖励个数

func main(){

	total := decimal.NewFromFloat(0.0)
	reward := decimal.NewFromFloat(InitReward)
	BtcInter := decimal.NewFromFloat(BtcInterval)
	Zero := decimal.NewFromFloat(0.000000000000000000000000000000000000001)
	for reward.Cmp(Zero) > 0{
		total = total.Add(reward.Mul(BtcInter))
		reward = reward.Mul(decimal.NewFromFloat(0.5))
	}

	fmt.Println(total.String())
}