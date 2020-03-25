package functor

import "razor-functor/src/dynamic"

func mapper(lambda dynamic.Lambda, list []interface{}) (result []interface{}) {
	for _, item := range list {
		if i, ok := item.([]interface{}); ok {
			result = mapper(lambda, i)
		} else {
			
		}
	}
	return
}

func filter(lambda dynamic.Lambda, list []interface{}) (result []interface{}) {
	return
}

func reduce(lambda dynamic.Lambda, list []interface{}) (result []interface{}) {
	return
}

func flatten(list []interface{}) (result []interface{}) {
	for _, item := range list {
		if i, ok := item.([]interface{}); ok {
			result = append(result, flatten(i)...)
		} else {
			result = append(result, item)
		}
	}
	return
}
