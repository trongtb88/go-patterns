package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func getData(key string) (string, error) {
	fmt.Println("invoked getData func")
	time.Sleep(2 * time.Second)
	return "data for " + key, nil
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, _, _ := group.Do("key", func() (interface{}, error) {
				return getData("key")
			})
			fmt.Println(data)
		}()

	}
	wg.Wait()
}
