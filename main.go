package main

import (
	"razor-functor/src/dynamic"
	"razor-functor/src/logger"
)

func main() {
	logger.Debug(true)
	dynamic.Call(func(i int) int {
		return i + 1
	}, 1, 2, 3)
}
