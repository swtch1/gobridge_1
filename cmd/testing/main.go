package main

import (
	"fmt"
)

func main() {
	// addToCache("foo", "bar")
	// addToCache("bar", "baz")
	// addToCache("baz", "bam")

	cache := newNameCache()
	cache.set("foo", "bar")
	cache.set("bar", "baz")
	cache.set("baz", "bam")

	v := cache.get("foo")
	fmt.Println("foo", v)
	v = cache.get("bar")
	fmt.Println("bar", v)
	v = cache.get("baz")
	fmt.Println("baz", v)
}

// create your own type and pass the requirements, or create them
// in the constructor here.
func newNameCache() *nameCache {
	return &nameCache{
		cache: make(map[string]string),
	}
}

type nameCache struct {
	cache map[string]string
}

func (nc *nameCache) set(fName, lName string) {
	nc.cache[fName] = lName
}

func (nc *nameCache) get(fName string) string {
	lName := nc.cache[fName]
	return lName
}

// don't do this!
// var (
// 	cache = make(map[string]string)
// )
//
// func addToCache(x, y string) {
// 	cache[x] = y
// }
//
// func getFromCache(x string) (string, bool) {
// 	v, ok := cache[x]
// 	return v, ok
// }
//
// func validateType(x someType) error {
// 	if x.name == "" {
// 		return fmt.Errorf("gotta have a name")
// 	}
// 	db.GetTypeByName(x.name)
// }
