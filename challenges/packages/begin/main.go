// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
)

// import the proverbs package

// getRandomProverb returns a random proverb from the proverbs package

func main() {
	fmt.Println(getRandomProverb())
	// print the result of calling your getRandomProverb function
}

func getRandomProverb() string {
	return(proverbs.Random().Saying)
}