// Go program to illustrate the concept
// of splitting a string
package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and Splitting a string with a separator
	// Using SplitN function
        data := string("a:b:c:d")

	lines := strings.SplitN(string(data), ":", -1)
	fmt.Println(lines)

}
