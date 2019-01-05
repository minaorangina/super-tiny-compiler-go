package tinycompiler

import "github.com/minaorangina/super-tiny-compiler/types"

// Parser converts a slice of tokens into its Abstract Syntax Tree (AST) representation
func Parser(tokens []types.Token) (types.Ast, error) {
	current := 0
	var walk func() types.Node
	walk = func() types.Node {
		token := tokens[current]

		switch token.TokenType {
		case "number":
			current++
			return types.Node{
				NodeType: types.NumberLiteral,
				Value:    token.Value,
				Params:   nil,
			}
		case "string":
			current++
			return types.Node{
				NodeType: types.StringLiteral,
				Value:    token.Value,
				Params:   nil,
			}
		case "paren":
			if token.Value == "(" {
				// skip parenthesis
				current++
				token = tokens[current]

				node := types.Node{
					NodeType: types.CallExpression,
					Value:    token.Value,
					Params:   []types.Node{},
				}

				// first parameter
				current++
				token = tokens[current]

				for (token.TokenType != "paren") ||
					(token.TokenType == "paren" && token.Value != ")") {
					node.Params = append(node.Params, walk())
					token = tokens[current]
				}
				// skip final parenthesis
				current++
				return node
			}
		default:
			panic("unknown token type")
		}

		return types.Node{}
	}

	ast := types.Ast{
		Body: []types.Node{},
	}

	for current < len(tokens) {
		ast.Body = append(ast.Body, walk())
	}

	return ast, nil
}
