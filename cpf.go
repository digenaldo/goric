package goric

import (
	"regexp"
	"strconv"
	"strings"
)

//CpfIsValid validates the cpf standard for Brazilian documents
func CpfIsValid(d string) (bool, error) {
	c, err := RemoveEspecChar(d)
	if err != nil {
		return false, err
	}

	if !CpfSizeIsValid(c) {
		return false, nil
	}

	if InvalidCpfIsKnown(c) {
		return false, nil
	}

	if !CpfDigitsValid(c) {
		return false, nil
	}

	return true, nil
}

//RemoveEspecChar remove special characters
func RemoveEspecChar(d string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(d, ""), nil
}

//CpfSizeIsValid verifies that the cpf size is 11
func CpfSizeIsValid(d string) bool {
	size := 11

	if len(d) != size {
		return false
	}
	return true
}

//CpfDigitsValid check if the cpf digits are valid
func CpfDigitsValid(d string) bool {
	s := strings.Split(d[9:], "")
	c1, _ := strconv.Atoi(s[0])
	c2, _ := strconv.Atoi(s[1])

	d1 := GenerateCpfFirstDigit(d)
	d2 := GenerateCpfSecondDigit(d)

	if d1 != c1 || d2 != c2 {
		return false
	}

	return true
}

//InvalidCpfIsKnown verifies that the cpf is in the known invalid format
func InvalidCpfIsKnown(d string) bool {
	knowCpfs := map[string]bool{"00000000000": true, "11111111111": true,
		"22222222222": true, "33333333333": true, "44444444444": true,
		"55555555555": true, "66666666666": true, "77777777777": true,
		"88888888888": true, "99999999999": true}

	if knowCpfs[d] {
		return true
	}
	return false
}

//GenerateCpfFirstDigit generate first cpf digit
func GenerateCpfFirstDigit(d string) int {
	size := 11

	c := strings.Split(d, "")
	soma := 0
	h := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	for i, v := range h {
		s, _ := strconv.Atoi(c[i])
		soma += v * s
	}
	r := soma % size

	if r < 2 {
		return 0
	}
	return size - r
}

//GenerateCpfSecondDigit generate second cpf digit
func GenerateCpfSecondDigit(d string) int {
	size := 11

	c := strings.Split(d, "")
	soma := 0
	h := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	for i, v := range h {
		s, _ := strconv.Atoi(c[i])
		soma += v * s
	}
	r := soma % size

	if r < 2 {
		return 0
	}
	return size - r
}
