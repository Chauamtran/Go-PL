package main

import (
	"fmt"
)

// Complete the sockMerchant function below.
func sockMerchant(n int, ar []int32) int {
	// index_group := 0
	group_dict := make(map[int32]int)
	var count int = 0

	fmt.Println("AA ", ar)
	for i := 0; i < n; i++ {
		group_dict[ar[i]] += 1
	}

	for _, value := range group_dict {
		fmt.Println("count: ", count)
		count += (value / 2)
	}

	fmt.Println("BB ", group_dict, count)
	return count

}

func main() {

	n := 15
	ar := []int32{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
	result := sockMerchant(n, ar)

	fmt.Println("%d\n", result)

}
