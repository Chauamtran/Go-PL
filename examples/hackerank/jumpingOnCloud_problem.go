package main

import (
	"fmt"
)


// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
    
    var count int32 = 0

    for i := 0; i < len(c); i++ {
        
        // If next value is OK
        if i + 2 < len(c) && c[i + 2] == 0 {
            if i + 2 < len(c) {
                count += 1
                i += 1
            } else {
                count += 1
                break
            }
            
        } else if i + 2 == len(c){
            count += 1
            break
        } else {
            if i + 1 < len(c) {
                count += 1
                continue
            } else {
                break
            }
        }
    }

    return count

}

func main() {

	var c []string{0,0,1,0,0,1,0}
	result := jumpingOnClouds(c)

	fmt.Fprintf(writer, "%d\n", result)

}
