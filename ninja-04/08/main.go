package main

import "fmt"

func main() {
	mxs := make(map[string][]string)
	mxs[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mxs[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	mxs[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}
	for k, xs := range mxs {
		fmt.Println(k)
		for _, s := range xs {
			fmt.Printf("\t%q", s)
		}
		fmt.Print("\n")
	}
}
