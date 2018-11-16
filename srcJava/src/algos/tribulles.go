package algos

import (
	"fmt"
	"os"
	time2 "time"
)

/**
 * le tri à bulles
 *
 * Ex : go run tribulles.go 6 2 5 7 0 1 4
 */
func main() {
	var tabs []string
	tabs = os.Args[1:]

	fmt.Print("tableau initial (")
	fmt.Print(len(tabs))
	fmt.Print(")")
	fmt.Println(tabs)

	permut := true
	passage := 1

	for permut {
		permut = false

		for index := 0; index < len(tabs)-passage; index++ {

			if tabs[index] > tabs [index+1] {
				tmp := tabs[index]
				tabs[index] = tabs[index+1]
				tabs[index+1] = tmp
				permut = true
			}
			printTab(tabs, permut)
		}
		passage ++
	}
}

func printTab(tab []string, permuted bool) {

	for i := 0; i < len(tab); i++ {
		fmt.Print(tab[i])

		if i != len(tab)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println("")
		}
	}
}
