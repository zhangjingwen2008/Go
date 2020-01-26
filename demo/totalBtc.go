package main

import "fmt"

//计算比特币总量
func main() {
	//1.每21万个区块挖出的比特币减半
	//2.最初奖励50个比特币
	//循环判断 累加

	total := 0.0          //比特币总量
	blockInterval := 21.0 //比特币奖励减半周期
	currentReward := 50.0 //最开始的奖励

	for currentReward > 0 {
		amount := blockInterval * currentReward
		currentReward *= 0.5 //除操作效率低，换用等价的乘操作
		total += amount
	}

	fmt.Println("比特币总量：", total, "万")
}
