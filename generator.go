package generator

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Package struct {
	Name      string
	Imports   []string
	Functions []Function
	Structs   []Struct
}

type Function struct {
	Name       string
	Reciever   string
	Statements []Statement
}

type Statement struct {
	AssignStatement string
	IfStatement     *IfStatement
	FunctionCall    *FunctionCall
}

type AssignStatement struct {
	Left  string
	Right *Expression
}

type Expression struct {
}

type IfStatement struct {
	BoolExpression *BooleanExpression
	Body           []Statement
	Else           *IfStatement
}

type BooleanExpression struct {
	Left  string
	Right string
}

type FunctionCall struct {
	Name string
	Args []Args
}

type Args struct {
	IsString bool
	Value    string
}

type Struct struct {
	Name     string
	Elements []StructElement
}

type StructElement struct {
	Name string
	Type string
}
