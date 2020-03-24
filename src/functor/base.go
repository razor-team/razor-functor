package functor

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
