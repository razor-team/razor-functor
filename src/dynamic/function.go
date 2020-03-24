package dynamic

import (
	"reflect"
	"runtime"

	"razor-functor/src/logger"
)

//Lambda callable function
type Lambda interface{}

//Call any signature callable function
func Call(lambda Lambda, params ...interface{}) []reflect.Value {
	in := make([]reflect.Value, 0, 0)
	for _, param := range params {
		in = append(in, reflect.ValueOf(param))
	}
	v := reflect.ValueOf(lambda)
	out := make([]reflect.Value, 0, 0)
	for _, i := range in {
		ret := v.Call([]reflect.Value{i})
		out = append(out, ret...)
		logger.Log("func: %s; args: %v; return %v;", runtime.FuncForPC(v.Pointer()).Name(), i, ret[0].Interface())
	}
	return out
}
