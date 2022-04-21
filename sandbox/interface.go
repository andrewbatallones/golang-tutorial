package sandbox

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/andrewbatallones/golang-tutorial/sections"
	"github.com/andrewbatallones/golang-tutorial/shapes"
)

// Purpose: I want to see if I save unique objects as an interface, that later, I can get them.
var testCache map[string]sections.Geometry = map[string]sections.Geometry{}
var mu sync.RWMutex

func getCachedShape(key string) sections.Geometry {
	mu.RLock()
	defer mu.RUnlock()

	return testCache[key]
}

func cacheShape(key string, g sections.Geometry) {
	mu.Lock()
	defer mu.Unlock()

	testCache[key] = g
}

func Mutex() {
	key := "rect"
	r := shapes.Rect{
		Width: 4,
		Height: 4,
	}
	cacheShape(key, r)
	g := getCachedShape(key)
	cachedRect, ok := g.(shapes.Circle); if !ok {
		fmt.Println("Not converted...")
	}

	fmt.Println(reflect.TypeOf(g))
	fmt.Println(cachedRect)
}