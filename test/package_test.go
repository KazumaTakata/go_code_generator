package test

import (
	"github.com/KazumaTakata/go_code_generator"
	"io/ioutil"
	"os"
	"testing"
	"text/template"
)

func TestIf(t *testing.T) {

	package_temp := generator.Package{Name: "main", Imports: []string{"fmt", "math/rand"}, Functions: []generator.Function{}}

	mainfunc := generator.Function{Name: "main", Statements: []generator.Statement{}}

	ifstatement := generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "i", Right: "0"}, Body: []generator.Statement{generator.Statement{IfStatement: &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "k", Right: "3"}}}}}

	ifstatement.Else = &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "k", Right: "20"}}
	statement := generator.Statement{IfStatement: &ifstatement}

	mainfunc.Statements = append(mainfunc.Statements, statement)

	package_temp.Functions = append(package_temp.Functions, mainfunc)

	buf, err := ioutil.ReadFile("../sample.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("package").Parse(string(buf))
	if err != nil {
		panic(err)
	}

	f, err := os.Create("generated/if.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, package_temp)
	if err != nil {
		panic(err)
	}

}
func TestPackage(t *testing.T) {

	package_temp := generator.Package{Name: "main", Imports: []string{"fmt", "math/rand"}, Functions: []generator.Function{}}

	mainfunc := generator.Function{Name: "main", Statements: []generator.Statement{}}

	statement := generator.Statement{FunctionCall: &generator.FunctionCall{Name: "fmt.Println", Args: []generator.Args{generator.Args{Value: "My favorite number is", IsString: true}, generator.Args{Value: "rand.Intn(10)", IsString: false}}}}

	mainfunc.Statements = append(mainfunc.Statements, statement)

	package_temp.Functions = append(package_temp.Functions, mainfunc)

	buf, err := ioutil.ReadFile("../sample.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("package").Parse(string(buf))
	if err != nil {
		panic(err)
	}

	f, err := os.Create("generated/hello.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, package_temp)
	if err != nil {
		panic(err)
	}

}
