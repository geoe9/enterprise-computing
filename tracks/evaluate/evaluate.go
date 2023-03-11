package evaluate

import (
	"enterprise-computing/tracks/repository"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func Evaluate(form string) int {
	if expr, err := parser.ParseExpr(form); err == nil {
		return natural(expr)
	} else {
		return 0
	}
}

func natural(expr ast.Expr) int {
	switch t := expr.(type) {
	case *ast.BasicLit:
		if t.Kind == token.INT {
			if val, err := strconv.Atoi(t.Value); err == nil {
				return val
			}
		}
	case *ast.Ident:
		if t, n := repository.Read(t.Name); n == 1 {
			return Evaluate(t.Audio)
		}
	case *ast.BinaryExpr:
		x := natural(t.X)
		y := natural(t.Y)
		switch t.Op {
		case token.ADD:
			return x + y
		case token.SUB:
			return x - y
		case token.MUL:
			return x * y
		case token.QUO:
			return x / y
		}
	}
	return 0
}
