package main

import (
	"fmt"
	"github.com/ivpusic/urlreg"
)

func main() {
	reg := urlreg.Pattern("some/:name/path/:other/")
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
