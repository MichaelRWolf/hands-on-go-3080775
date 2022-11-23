// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getRandomProverb())
	// print the result of calling your getRandomProverb function
}

func getRandomProverb() string {
	return(proverbs.Random().Saying)
}
