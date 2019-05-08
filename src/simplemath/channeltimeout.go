package main

import "fmt"

func gongzhen(n float64) float64 {
	var res float64 = 1.08722 * (70000*20 - n*n - (float64)(n/3)*(n/3) - (float64)(n/5)*(n/5) -
		(n/8)*(n/8) - (250000-(n-499)*(n-499))*2.2)
	return res
}

func main() {
	var sum float64 = 0
	for i := 1; i <= 100; i++ {
		res := gongzhen(float64(i))
		sum += res
		fmt.Println(res, "  ", sum)
	}

}
