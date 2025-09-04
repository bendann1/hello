package main

import (
	"fmt"

	"github.com/bendann1/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := "Hello, World!"
	rs := morestrings.Reverse(s)
	fmt.Println(cmp.Diff(s, rs))
}
