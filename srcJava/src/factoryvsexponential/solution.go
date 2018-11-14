package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

// https://www.codingame.com/training/hard/factorial-vs-exponential

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var K int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&K)

	var arrayA = make([]float64, K)
	scanner.Scan()
	inputs := strings.Split(scanner.Text()," ")

	var arrays = make([]float64, K)

	for i := 0; i < K; i++ {
		A,_ := strconv.ParseFloat(inputs[i],10)
		_ = A
		arrayA[i] = A
	}

	for i := 0; i < K; i++ {
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Fprint(os.Stderr, "\n\n-> ")
		fmt.Fprintln(os.Stderr, arrayA[i])

		var cpt float64 = 1
		for cpt != -1 {
			var expo float64 = calcExponential(arrayA[i], cpt)
			var fact float64 = calcFactorial(cpt)

			fmt.Fprint(os.Stderr, arrayA[i])
			fmt.Fprint(os.Stderr, " | " )
			fmt.Fprint(os.Stderr, cpt )
			fmt.Fprint(os.Stderr, "\t expo : ")
			fmt.Fprint(os.Stderr, expo)

			fmt.Fprint(os.Stderr, " - fact : " )
			fmt.Fprintln(os.Stderr, fact )

			if expo < fact {
				//fmt.Print(cpt)
				//fmt.Print(" ")
				arrays[i] = cpt
				cpt = -1
				break
			}

			cpt ++
		}

	}

	fmt.Fprintln(os.Stderr, "end")

	for i := 0; i < K; i++ {
		fmt.Print(arrays[i])

		// last element
		if i != K - 1 {
			fmt.Print(" ")
		}
	}
}

func calcExponential(value float64, puissance float64) float64 {
	var result float64 = 1

	if value == 0 {
		return 0
	}

	for i := 0; i < int (puissance); i++ {
		result = result * value
	}

	return result
}
func calcFactorial(value float64) float64 {
	if value == 0 {
		return 1
	} else {
		return value * calcFactorial(value - 1)
	}
}
