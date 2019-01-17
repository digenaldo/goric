package goric

import "testing"

func TestCpfIsValid(t *testing.T) {
	got := CpfIsValid("111.111.111-11")
	want := true

	if got != want {
		t.Errorf("got: %t is different from want: %t", got, want)
	}
}

func TestRemoveEspecChar(t *testing.T) {
	got := RemoveEspecChar("111.111.111-11")
	want := "11111111111"

	if got != want {
		t.Errorf("got: %s is different from want: %s", got, want)
	}

}

func TestCpfSizeIsValid(t *testing.T) {
	got := CpfSizeIsValid("11111111111")
	want := true

	if got != want {
		t.Errorf("got: %t is different from want: %t", got, want)
	}
}

func TestInvalidCpfIsKnown(t *testing.T) {
	got := InvalidCpfIsKnown("22222222229")
	want := false

	if got != want {
		t.Errorf("got: %t is different from want: %t", got, want)
	}
}
