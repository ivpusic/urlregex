package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/:name/path/:other/")
	fmt.Println("regex: " + reg.Regex.String())

	res, err := reg.Match("some/123/path/456/")
	if err != nil {
		fmt.Println("no matches")
		return
	}

	fmt.Println("found matches")
	for k, v := range res {
		fmt.Println(k + ": " + v)
	}
}
