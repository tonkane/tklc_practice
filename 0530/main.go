package main

import (
	// "context"
	"fmt"
	"sync"
)

func maximumLength(s string) int {
	r := [27][2]int{}
	n := len(s)
	left := 0
	right := 0
	ans := -1
	size := 0
	for right < n {
		for right < n && s[left] == s[right] {
			size = right - left + 1
			right++
		}
		ans = max(ans, size-2)
		t := s[left] - 'a'
		// 比首个大
		if size > r[t][0] {
			if r[t][0] != 0 {
				ans = max(ans, r[t][0])
			}
			r[t][1] = r[t][0]
			r[t][0] = size
		} else if size > r[t][1] {
			if r[t][0]-size <= 1 {
				ans = max(ans, r[t][0]-1)
			}
			r[t][1] = size
		} else if size == r[t][1] {
			ans = max(ans, size)
		}

		left = right
	}

	if ans == 0 {
		return -1
	}
	return ans
}

func abc(a []int) {
	a[0] = 3
}

type MyStruct struct {
	Field1 int
	Field2 string
	// 可以添加更多字段...
}


// 创建一个返回闭包的函数
func newClosure() func() int {
    x := 0
    return func() int {
        x++ // 闭包记住了 x 的状态
        return x
    }
}


func pa (a string, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- true
		fmt.Println(a)
		<-ch
	}
}


func printLetter(letter byte, chIn, chOut chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// 等待接收信号开始打印
		<-chIn
		fmt.Println(string(letter))
		// 发送信号给下一个 goroutine
		chOut <- true
	}
	if letter == 'a' {
		<-chIn
	}
}

func main() {
	var wg sync.WaitGroup
	ch1, ch2, ch3 := make(chan bool), make(chan bool), make(chan bool)

	// 初始化通道，ch1 用于触发第一个 goroutine，ch3 用于接收最后一个 goroutine 的信号
	
	// go func() {
		defer close(ch1) // 确保在最后关闭通道
		wg.Add(3)
		go printLetter('a', ch1, ch2, &wg)
		go printLetter('b', ch2, ch3, &wg)
		go printLetter('c', ch3, ch1, &wg) // 形成循环

		ch1 <- true
	// }()
	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("打印完成")

}
