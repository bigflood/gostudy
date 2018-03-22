package gocalc

import (
	"strconv"
	"fmt"
	"reflect"
	"go/parser"
	"go/ast"
	"go/token"
)

func CalcExpr(src string) (float64, error) {
	expr, err := parser.ParseExpr(src)
	if err != nil {
		panic(err)
	}

	return calcExpr(expr)
}


func calcExpr(expr ast.Expr) (float64, error) {
	switch p := expr.(type) {
	case *ast.BasicLit:
		return strconv.ParseFloat(p.Value, 64)
	case *ast.BinaryExpr:
		x, err := calcExpr(p.X)
		if err != nil {
			return 0, err
		}
		y, err := calcExpr(p.Y)
		if err != nil {
			return 0, err
		}

		switch p.Op {
		case token.MUL:
			return x * y, nil
		case token.QUO:
			return x / y, nil
		case token.ADD:
			return x + y, nil
		case token.SUB:
			return x - y, nil
		default:
			return 0, fmt.Errorf("%v:unknown operator", p.OpPos)
		}
	case *ast.ParenExpr:
		return calcExpr(p.X)
	default:
		return 0, fmt.Errorf("%v:%v:unknown expr (%v)", expr.Pos(), expr.End(), reflect.TypeOf(expr))
	}
}
