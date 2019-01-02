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
				NodeType: "CallExpression",
				Value:    "add",
				Params:   []Node{},
			},
		},
	}

	result, err := Parser(tokens)
	if err != nil {
		return t.Fail()
	}
	if reflect.DeepEqual(expected, result) {
		return t.Errorf("ASTs not equal.\nExpected:\n%+v\nActual:\n%+v\n", expected, result)
	}
}
