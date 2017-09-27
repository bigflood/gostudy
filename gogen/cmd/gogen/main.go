package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Interface struct {
	Name    string
	Methods []Func
}

type Func struct {
	Name    string
	Params  []Field
	Results []Field
}

type Field struct {
	Name string
	Type string
}

func getInterfacesFrom(f *ast.File) []Interface {
	var list []Interface

	for _, d := range f.Decls {
		if g, ok := d.(*ast.GenDecl); ok {
			for _, s := range g.Specs {
				if ts, ok := s.(*ast.TypeSpec); ok {
					if i, ok := ts.Type.(*ast.InterfaceType); ok {
						list = append(list, getInterface(ts.Name.Name, i))
					}
				}
			}
		}
	}

	return list
}

func getInterface(name string, i *ast.InterfaceType) Interface {
	return Interface{
		Name:    name,
		Methods: getMethods(i.Methods),
	}
}

func getMethods(methods *ast.FieldList) []Func {
	var list []Func

	for _, m := range methods.List {
		f := Func{}

		for _, n := range m.Names {
			f.Name = n.Name
			list = append(list, f)
		}
	}

	return list
}

func main() {
	fset := token.NewFileSet()

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "../../service/service.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range getInterfacesFrom(f) {
		fmt.Println(i)
	}

	//ast.Print(fset, f)

}
