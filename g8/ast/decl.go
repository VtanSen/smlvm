package ast

import (
	"e8vm.io/e8vm/lex8"
)

// IdentList is a list of identifiers
type IdentList struct {
	Idents []*lex8.Token
	Commas []*lex8.Token
}

// VarDecl declares a set of variables. It is both a top level declaration
// and a statement.
type VarDecl struct {
	Idents *IdentList
	Type   Expr
	Eq     *lex8.Token
	Exprs  *ExprList
	Semi   *lex8.Token
}

// VarDecls is a variable declaration with a leading keyword
// It could be a single decl or a decl block.
type VarDecls struct {
	Kw     *lex8.Token
	Lparen *lex8.Token // optional
	Decls  []*VarDecl
	Rparen *lex8.Token // optional
	Semi   *lex8.Token
}

// ConstDecl declares a set of constants.
type ConstDecl struct {
	Idents *IdentList
	Type   Expr
	Eq     *lex8.Token
	Exprs  *ExprList
	Semi   *lex8.Token
}

// ConstDecls is a const declaration with a leading keyword
// It could be a single decl or a decl block.
type ConstDecls struct {
	Kw     *lex8.Token
	Lparen *lex8.Token // optional
	Decls  []*ConstDecl
	Rparen *lex8.Token // optional
	Semi   *lex8.Token
}
