package tinycompiler

import (
	"reflect"
	"testing"

	"github.com/minaorangina/super-tiny-compiler/types"
)

func TestParser(t *testing.T) {
	tokens := []types.Token{
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

	expected := types.Ast{
		Body: []types.Node{
			types.Node{
				NodeType: types.CallExpression,
				Value:    "add",
				Params: []types.Node{
					types.Node{types.NumberLiteral, "2", nil},
					types.Node{types.CallExpression, "subtract", []types.Node{
						types.Node{types.NumberLiteral, "4", nil},
						types.Node{types.NumberLiteral, "2", nil},
					}},
				},
			},
		},
	}

	result, err := Parser(tokens)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("ASTs not equal.\nExpected:\n%+v\nActual:\n%+v\n", expected, result)
	}

	tokens = []types.Token{
		types.Token{"paren", "("},
		types.Token{"name", "add"},
		types.Token{"notanodetype", "2"},
		types.Token{"paren", "("},
		types.Token{"name", "subtract"},
		types.Token{"number", "4"},
		types.Token{"number", "2"},
		types.Token{"alsonotanodetype", ")"},
		types.Token{"paren", ")"},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic when unknown NodeType encountered")
		}
	}()

	Parser(tokens)
}
