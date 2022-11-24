package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"Bisshwajit",
		"Samanta",
		"Payelh",
		"Sinchan",
	}
	wg.Add(3)
	for i, v := range words {
		go PrintSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}
