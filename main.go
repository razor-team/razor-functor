package main

import (
	"fmt"
	"razor-functor/src/functor"
	"razor-functor/src/logger"
	"razor-functor/src/stl"
)

func main() {
	logger.Debug(true)
	data := functor.F(stl.List(1, 2, stl.List(1, 23, stl.List(7, 4)), stl.List(3))).Flatten()
	fmt.Println(data.Result())
}
