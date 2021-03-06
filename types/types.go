package types

type NodeType int

const (
	NumberLiteral NodeType = iota
	StringLiteral
	CallExpression
)

// Token represents a unit of lisp-syntax code
type Token struct {
	TokenType string
	Value     string
}

// Ast is a representation of the structure and grammar of lisp-syntax code
type Ast struct {
	Body []Node
}

// Node is a node in an AST structure
type Node struct {
	NodeType NodeType // switch over this rather than the type
	Value    string
	Params   []Node
}
