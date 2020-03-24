package stl

//TList dyncamic array
type TList struct {
	data interface{}
}

//List make array from any data type
func List(list ...interface{}) []interface{} {
	return list
}
