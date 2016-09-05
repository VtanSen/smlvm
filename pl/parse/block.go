package parse

import (
	"e8vm.io/e8vm/lexing"
	"e8vm.io/e8vm/pl/ast"
)

func parseBlock(p *parser) *ast.Block {
	ret := new(ast.Block)
	ret.Lbrace = p.ExpectOp("{")
	if ret.Lbrace == nil {
		return ret
	}

	for !(p.SeeOp("}") || p.See(lexing.EOF)) {
		if stmt := p.parseStmt(); stmt != nil {
			ret.Stmts = append(ret.Stmts, stmt)
		}
		p.skipErrStmt()
	}

	ret.Rbrace = p.ExpectOp("}")
	return ret
}