package goric

import (
	"log"
	"regexp"
)

//CpfIsValid validates the cpf standard for Brazilian documents
func CpfIsValid(d string) bool {
	if !CpfSizeIsValid(d) {
		return false
	}
	// a := RemoveEspecChar(d)
	return false
}

//RemoveEspecChar remove special characters
func RemoveEspecChar(d string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(d, "")
}

//CpfSizeIsValid verifies that the cpf size is 11
func CpfSizeIsValid(d string) bool {
	size := 11

	if len(d) != size {
		return false
	}
	return true
}

//CpfFormatIsValid validates cpf format with brazilian algorithm
func CpfFormatIsValid(d string) bool {
	return false
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
