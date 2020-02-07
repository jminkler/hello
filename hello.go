package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jminkler/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello, world", "hello go"))
}
