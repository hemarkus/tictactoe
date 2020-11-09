package tictactoe

import (
	"strings"
	"testing"
)

func Test_genericDialog_String(t *testing.T) {
	var stringResult string

	const input = "just a phrase\n"

	if err := genericDialog("Read string", "default", &stringResult, strings.NewReader(input)); err != nil {
		t.FailNow()
	}
	if stringResult != "just a phrase" {
		t.Errorf("Expected %s got %s", input, stringResult)
	}
}

func Test_genericDialog_DefaultString(t *testing.T) {
	var stringResult string

	const input = "\n"

	if err := genericDialog("Read string", "default", &stringResult, strings.NewReader(input)); err != nil {
		t.FailNow()
	}
	if stringResult != "default" {
		t.Errorf("Expected %s got %s", input, stringResult)
	}
}

func Test_genericDialog_DefaultUint(t *testing.T) {
	var uintResult uint

	const input = "\n"
	const defaultResult = uint(9)

	if err := genericDialog("Read positive number", defaultResult, &uintResult, strings.NewReader(input)); err != nil {
		t.FailNow()
	}
	if uintResult != defaultResult {
		t.Errorf("Expected %d got %d", defaultResult, uintResult)
	}
}

func Test_genericDialog_Uint(t *testing.T) {
	var uintResult uint

	const input = "5\n"
	const defaultResult = uint(9)

	if err := genericDialog("Read positive number", defaultResult, &uintResult, strings.NewReader(input)); err != nil {
		t.FailNow()
	}
	if uintResult != uint(5) {
		t.Errorf("Expected %d got %d", defaultResult, uintResult)
	}
}

func Test_genericDialog_Type(t *testing.T) {
	var floatResult float64

	const input = "5\n"
	const defaultResult = uint(9)

	if err := genericDialog("Read uint to float", defaultResult, &floatResult, strings.NewReader(input)); err != InconsistantTypesProvidedErr {
		t.FailNow()
	}
}

func Test_genericDialog_CallByValue(t *testing.T) {
	var target uint
	const input = "5\n"
	const defaultResult = uint(9)

	if err := genericDialog("Read uint to float", defaultResult, target, strings.NewReader(input)); err != InvalidCallByValue {
		t.FailNow()
	}
}
