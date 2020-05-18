package main

import (
	"fmt"
)

func exponent(n int, e int) float64 {
	var s float64
	for i := 0; i < e; i++ {
		if s == 0 {
			s = float64(n)
		} else {
			s = s * float64(n)
		}
	}
	return s
}

func getDigits(n int) ([]int, error) {
	var nArray []int

	var thousands int = n / 1000
	if (n / 1000) > 0 {
		nArray = append(nArray, thousands)
	}
	n = n - (thousands * 1000)

	var hundreds int = n / 100
	if hundreds > 0 {
		nArray = append(nArray, hundreds)
	}
	n = n - (hundreds * 100)

	var tens int = n / 10
	if tens > 0 {
		nArray = append(nArray, tens)
	}
	n = n - (tens * 10)

	var ones int = n / 1
	if ones > 0 {
		nArray = append(nArray, ones)
	}
	n = n - ones

	if n != 0 {
		return nArray, fmt.Errorf("Something went wrong separating digits")
	}

	return nArray, nil
}

func isArmstrong(n int) (bool, error) {
	dig, err := getDigits(n)
	if err != nil {
		return false, err
	}
	if len(dig) > 3 {
		return false, fmt.Errorf("Number is too long with %d digits", len(dig))
	}
	fmt.Println("Number: ", n)
	fmt.Println("Digits: ", dig)
	var sum int
	for _, i := range dig {
		exp := int(exponent(i, 3))
		sum = sum + exp
		fmt.Println(i, exp)

	}
	if sum == n {
		return true, nil
	}

	return false, fmt.Errorf("Number %d is not Armstrong number: %d", n, sum)
}

func main() {
	fmt.Println(isArmstrong(371))
}
