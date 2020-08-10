package romannumerals

import "errors"

var I = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var X = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var C = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var M = [4]string{"", "M", "MM", "MMM"}

func ToRomanNumeral(input int) (string, error) {

	if input > 3000 || input <= 0 {
		return "", errors.New("is Error")
	} else {
		thousand := input / 1000
		hundred := (input - thousand*1000) / 100
		ten := (input - thousand*1000 - hundred*100) / 10
		one := input % 10
		result := M[thousand] + C[hundred] + X[ten] + I[one]
		return result, nil
	}

}
