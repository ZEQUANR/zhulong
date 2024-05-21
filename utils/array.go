package utils

import (
	"fmt"
	"math/rand"
)

func RandomElementFromSlice(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}

	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex], nil
}
