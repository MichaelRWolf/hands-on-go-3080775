// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}


type deepThoughtCounter struct { question string}
func (pan_dimensional_beings deepThoughtCounter) name() string { 
	return pan_dimensional_beings.question
}
func (pan_dimensional_beings deepThoughtCounter) count(s string) int { 
	return 42
}

type letterCounter struct{ identifier string }
func (lc letterCounter) name() string {
	return lc.identifier
}
func (lc letterCounter) count(s string) int {
	return 4
}

type numberCounter struct{ designation string }
func (nc numberCounter) name() string {
	return nc.designation
}
func (lc numberCounter) count(s string) int {
	return 44
}


type symbolCounter struct{ label string }
func (sc symbolCounter) name() string {
	return sc.label
}
func (sc symbolCounter) count(s string) int {
	return 55
}

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	for i, c := range counters {
		key := strconv.Itoa(i) + "_" + c.name()
		value := c.count(data)
		analysis[key] = value
	}

	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	input_file_name := "/dev/null"
	if (len(os.Args) == 2) {
   	input_file_name = os.Args[1]
	}

	data_bytes, err := os.ReadFile(input_file_name)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	
	data_string := string(data_bytes)
	spew.Dump(data_string)

	analysis :=	doAnalysis(
		data_string, 
		letterCounter      { identifier: "letters"},
		deepThoughtCounter { question: "What is the answer to life, the Universe, and everything?"},
		numberCounter      { designation:  "One, two, buckle my shoe" },
		symbolCounter      { label: "Lapel"},
	)

	spew.Dump(analysis)
}
