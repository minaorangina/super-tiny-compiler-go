package tinycompiler

import (
	"reflect"
	"testing"

	"github.com/minaorangina/super-tiny-compiler/types"
)

func TestTokeniser(t *testing.T) {
	sampleInput := "(add 2 (subtract 4 2))"
	expected := []types.Token{
		types.Token{"paren", "("},
		types.Token{"name", "add"},
		types.Token{"number", "2"},
		types.Token{"paren", "("},
		types.Token{"name", "subtract"},
		types.Token{"number", "4"},
		types.Token{"number", "2"},
		types.Token{"paren", ")"},
		types.Token{"paren", ")"},
	}

	result, err := Tokeniser(sampleInput)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Token slices not equal.\nExpected:\n%+v\nActual:\n%+v\n", expected, result)
	}
}
