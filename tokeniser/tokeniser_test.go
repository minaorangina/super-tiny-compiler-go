package tokeniser

import (
	"reflect"
	"testing"
)

func TestTokeniser(t *testing.T) {
	sampleInput := "(add 2 (subtract 4 2))"
	expected := []Token{
		Token{"paren", "("},
		Token{"name", "add"},
		Token{"number", "2"},
		Token{"paren", "("},
		Token{"name", "subtract"},
		Token{"number", "4"},
		Token{"number", "2"},
		Token{"paren", ")"},
		Token{"paren", ")"},
	}

	result, err := Tokeniser(sampleInput)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Token slices not equal.\nExpected:\n%+v\nActual:\n%+v\n", expected, result)
	}
}
