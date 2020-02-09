package main

import "fmt"

func main() {
	mxs := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	mxs[`gump_forrest`] = []string{`Sweets`, `Running`, `Jane`}

	delete(mxs, `no_dr`)
	if _, ok := mxs["no_dr"]; !ok {
		fmt.Println("Entry deleted successfully")
	}

	for k, xs := range mxs {
		fmt.Println(k)
		for _, s := range xs {
			fmt.Printf("\t%q", s)
		}
		fmt.Print("\n")
	}
}
