package main

import "fmt"

//直接将两个数组进行==、!=比较即可
func main() {
	var Numbers1 [5]int = [5]int{2, 6, 4, 7, 1}
	var Numbers2 [5]int = [5]int{2, 6, 44, 7, 1}

	b := Numbers1 == Numbers2
	//b:=CompareValue(Numbers1,Numbers2)

	if b {
		fmt.Println("数组一致")
	} else {
		fmt.Println("不一致！")
	}
}

/*
func CompareValue(n1[5]int, n2[5]int) (b bool) {
	if len(n1)==len(n2){
		for i:=0;i<len(n1);i++{
			if n1!=n2{
				break
			}
			b=true
		}
	}else{
		b=false
	}
	return
}
*/
