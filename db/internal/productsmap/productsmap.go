// # Implementation Constraints
//
// - Do not use generics
// - Use specific types for both input and output so that we only allow expected types in the map
//
// I was considering to write a template of this file just to experiment with code-generation, but I'll leave that for another time.
package productsmap

import (
	"sync"

	"github.com/applied-concurrency-in-go/models"
)

type valueType = models.Product

var zeroValue = valueType{}

// If I wrote `type Products sync.Map`, would the underlying type be exposed?
// If so, should I consider a struct with a single private field an encapsulation mechanism?
// Is there an added cost in using a struct with a single private field instead of creating a distinct name for an existing type?
type Map struct {
	m sync.Map
}

func (I Map) Load(key string) (valueType, bool) {
	if p, found := I.m.Load(key); found {
		return p.(valueType), found
	} else {
		return zeroValue, found
	}
}

func (I Map) Range(f func(key string, value valueType) bool) {
	I.m.Range(func(k, v interface{}) bool {
		key := k.(string)
		val := v.(valueType)
		return f(key, val)
	})
}

func (I Map) CompareAndSwap(key string, old valueType, new valueType) bool {
	return I.m.CompareAndSwap(key, old, new)
}

func (I Map) Delete(key string) {
	I.m.Delete(key)
}

func (I Map) LoadAndDelete(key string) (value valueType, loaded bool) {
	p, loaded := I.m.LoadAndDelete(key)
	return p.(valueType), loaded
}

func (I Map) LoadOrStore(key string, value valueType) (valueType, bool) {
	p, loaded := I.m.LoadOrStore(key, value)
	return p.(valueType), loaded
}

func (I Map) Store(key string, value valueType) {
	I.m.Store(key, value)
}

func (I Map) Swap(key string, value valueType) (valueType, bool) {
	p, loaded := I.m.Swap(key, value)
	return p.(valueType), loaded
}
