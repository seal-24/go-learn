package main

import "fmt"

func device(a int) int{
	c := a/(a-1)
	return c
}

/*
sleep 100 &
killall -SIGSEGV sleep
https://zhuanlan.zhihu.com/p/62825675
export GOBACTRACE=crash ./debug
 */
func main()  {
	a := 1
	b := device(a)
	fmt.Println(b)


}

