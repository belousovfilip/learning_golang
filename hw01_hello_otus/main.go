package main

import (
	"fmt"
	"regexp"
)

func main() {
	var words = map[string]int{}
	group := []struct {
		name string
		count int
	}{}
	reg := regexp.MustCompile(`[ ]+`)
	res := reg.Split("sad sd as d d d asd sa ds dasd sad sadfgvfdgfsdg fg sdf dsf sdf sf sd fsd f", -1)
	for _, v := range res {
		fmt.Println(words["sd"])
		words[v]++
	}
	fmt.Println(words.)
}
