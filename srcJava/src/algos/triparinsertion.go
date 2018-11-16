package algos

import (
	"fmt"
	"os"
)

/**
 * le tri par insertion
 *
 * Ex : go run triparinsertion.go 6 2 5 7 0 1 4
 */
func main() {
	var tabs []string

	tabs = os.Args[1:]

	fmt.Print("tableau initial (")
	fmt.Print(len(tabs))
	fmt.Print(")")
	fmt.Println(tabs)
	permut := true

	for i := 1; i < len(tabs); i ++ {
		permut = true
		//fmt.Printf("%s %s : ", strconv.Itoa(i), tabs[i])

		for permut {
			permut = false
			for j := 0; j < i; j++ {
				//fmt.Printf("j : %s \n", strconv.Itoa(j))

				if tabs[j] > tabs[j+1] {
					tmp := tabs[j]
					tabs[j] = tabs[j+1]
					tabs[j+1] = tmp

					permut = true
				}
			}
		}

		printTab2(tabs)

	}

}

func printTab2(tab []string) {

	for i := 0; i < len(tab); i++ {
		fmt.Print(tab[i])

		if i != len(tab)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println("")
		}
	}
}
