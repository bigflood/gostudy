package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
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

func (f Field) PascalCaseName() string {
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

func getInterfacesFrom(f *ast.File) map[string]Interface {
	result := make(map[string]Interface)

	for _, d := range f.Decls {
		if g, ok := d.(*ast.GenDecl); ok {
			for _, s := range g.Specs {
				if ts, ok := s.(*ast.TypeSpec); ok {
					if i, ok := ts.Type.(*ast.InterfaceType); ok {
						p := getInterface(ts.Name.Name, i)
						result[p.Name] = p
					}
				}
			}
		}
	}

	return result
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

		ft, _ := m.Type.(*ast.FuncType)

		// param
		f.Params = getFields(ft.Params)

		// result
		f.Results = getFields(ft.Results)

		for _, n := range m.Names {
			f.Name = n.Name
			list = append(list, f)
		}
	}

	return list
}

func getFields(fl *ast.FieldList) []Field {
	result := make([]Field, 0, len(fl.List))

	for _, f := range fl.List {
		result = append(result, getFieldsFromField(f)...)
	}

	return result
}

func getFieldsFromField(f *ast.Field) []Field {
	result := make([]Field, len(f.Names))
	tn := f.Type.(*ast.Ident).Name

	if len(f.Names) == 0 {
		result = append(result, Field{
			Name: "ret",
			Type: tn,
		})
	} else {
		for i, n := range f.Names {
			result[i] = Field{
				Name: n.Name,
				Type: tn,
			}
		}
	}

	return result
}

func main() {
	srcFile := flag.String("s", "", ".go source file path")
	tmplFile := flag.String("t", "", "template file path")
	infName := flag.String("i", "Service", "interface name")

	flag.Parse()

	if *srcFile == "" || *tmplFile == "" || *infName == "" {
		flag.Usage()
		os.Exit(-1)
	}

	fset := token.NewFileSet()

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, *srcFile, nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ast.Print(fset, f)

	t, err := template.ParseFiles(*tmplFile)
	if err != nil {
		panic(err)
	}

	def, ok := getInterfacesFrom(f)[*infName]
	if !ok {
		panic("Service interface not found!")
	}

	if err := t.Execute(os.Stdout, def); err != nil {
		panic(err)
	}

}
