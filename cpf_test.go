package goric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpfIsValid(t *testing.T) {
	got, _ := CpfIsValid("111.111.111-11")
	want := false

	got2, _ := CpfIsValid("11111111111")
	want2 := false

	got3, _ := CpfIsValid("432.484.831-99")
	want3 := true

	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)
	assert.Equal(t, want3, got3)
}
func TestRemoveEspecChar(t *testing.T) {
	got, err := RemoveEspecChar("111.111.111-11")
	want := "11111111111"

	assert.Equal(t, want, got)
	assert.Equal(t, nil, err)

}
func TestCpfSizeIsValid(t *testing.T) {
	got := CpfSizeIsValid("11111111111")
	want := true

	got2 := CpfSizeIsValid("11111111")
	want2 := false

	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)
}

func TestInvalidCpfIsKnown(t *testing.T) {
	got := InvalidCpfIsKnown("22222222229")
	want := false

	got2 := InvalidCpfIsKnown("33333333333")
	want2 := false

	assert.Equal(t, want, got)
	assert.NotEqual(t, want2, got2)
}
func TestGenerateCpfFirstDigit(t *testing.T) {
	got := GenerateCpfFirstDigit("44504356724")
	want := 2

	got2 := GenerateCpfFirstDigit("82607236207")
	want2 := 0

	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)

}
func TestGenerateCpfSecondDigit(t *testing.T) {
	got := GenerateCpfSecondDigit("44504356724")
	want := 4

	got2 := GenerateCpfSecondDigit("56043299160")
	want2 := 0

	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)

}
func TestCpfDigitsValid(t *testing.T) {
	got := CpfDigitsValid("44504356724")
	want := true

	got2 := CpfDigitsValid("82607236209")
	want2 := false

	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)

}
