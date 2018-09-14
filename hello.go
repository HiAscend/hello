package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 这是入口方法
func main() {
	matchers := make(map[string]string)
	if _, exists := matchers["hello"]; exists {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

	if matchers == nil {

	}

}
