package functor

import "razor-functor/src/dynamic"

//Functor type
type Functor struct {
	data interface{}
}

//Map call lambda each of functor collection data item
func (functor *Functor) Map(lambdas ...dynamic.Lambda) *Functor {
	return functor
}

//Filter a argument in collection
func (functor *Functor) Filter(lambdas ...dynamic.Lambda) *Functor {
	return functor
}

//Reduce all collection elements to result aggregated value
func (functor *Functor) Reduce(lambdas ...dynamic.Lambda) *Functor {
	return functor
}

//copy functor data
func (functor *Functor) copy() *Functor {
	return &Functor{
		data: functor.data,
	}
}
