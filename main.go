package main

import "fmt"

// Client - intetface
type Client interface {
	Get() int
}

// Sample - 構造体
type Sample struct {
	// interfaceを定義する
	client Client
}

// doGet - interfaceのGet()をラップする
func (sample *Sample) doGet() int {
	return sample.client.Get()
}

// Client interface の実装をする構造体
type hogeClientImpl struct {
}

// レシーバで実装する
func (client *hogeClientImpl) Get() int {
	return 1
}

func main() {
	//
	sample := &Sample{&hogeClientImpl{}}
	res := sample.doGet()
	fmt.Printf("result: %d", res)
}
