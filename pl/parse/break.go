package parse

import (
	"shanhu.io/smlvm/pl/ast"
)

func parseBreakStmt(p *parser, withSemi bool) *ast.BreakStmt {
	ret := new(ast.BreakStmt)
	ret.Kw = p.ExpectKeyword("break")
	if p.See(Ident) {
		ret.Label = p.Expect(Ident)
	}
	if withSemi {
		ret.Semi = p.ExpectSemi()
	}
	return ret
}
