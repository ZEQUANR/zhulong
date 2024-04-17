package main

import (
	"context"
	"fmt"
)

func main() {

	ctx1, cancel := context.WithCancel(context.Background())

	// ctx2 := context.Background()

	cancel()

	fmt.Println(<-ctx1.Done())
	// fmt.Println(<-ctx2.Done())
}
