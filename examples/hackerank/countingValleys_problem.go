package main

import (
	"fmt"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	countStep := 0
	var countValley int32 = 0
	firstStep := ""

	for _, step := range s {
		stepS := string(step)

		// Count steps
		if stepS == "D" {
			countStep += 1
		} else {
			countStep -= 1
		}

		if countStep == 0 {

			if firstStep == "D" {
				countValley += 1
			}

			firstStep = ""
			fmt.Println("countValley: ", countValley)

		} else if countStep == 1 && firstStep == "" {
			firstStep = stepS
			fmt.Println("firstStep: ", firstStep)
		}

	}

	return countValley

}

func main() {

	var n int32 = 8
	ar := "UDDDUDUU"

	result := countingValleys(n, ar)

	fmt.Println("%d\n", result)

}
