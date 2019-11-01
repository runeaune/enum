package enum

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// FormatFunc is the function to format the string value of the enum.
type FormatFunc func(string) string

// Parser is the parser that will search a file for constants and add each
// constant as an enum.
type Parser struct {
	Enums      []Enum
	File       string
	Format     FormatFunc
	LineStart  int
	Package    string
	TrimPrefix string
	TypeName   string
	ValueType  string
	WithJSON   bool
	WithValue  bool
}

// Enum is one enum with a number mapped to a string. The name of the enun will
// be the constant value.
type Enum struct {
	Int    int
	Value  string
	Name   string
	String string
}

// New will create a new parser to use for a given file.
func New(file, trimPrefix string, lineStart int, json, value bool, ff FormatFunc) *Parser {
	return &Parser{
		File:       file,
		Format:     ff,
		LineStart:  lineStart,
		TrimPrefix: trimPrefix,
		WithJSON:   json,
		WithValue:  value,
	}
}

// GetEnumFromFile will read a file and pass the content to GetEnum()
func (ep *Parser) GetEnumFromFile() error {
	fileData, err := ioutil.ReadFile(ep.File)
	if err != nil {
		return err
	}

	return ep.GetEnum(fileData)
}

// GetEnum will find all enum in one const block starting from the parsers
// LineStart.
func (ep *Parser) GetEnum(fileData []byte) error {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", fileData, parser.ParseComments)
	if err != nil {
		return err
	}

	constDeclaration := ep.findConstDeclaration(fset, file)
	if constDeclaration == nil {
		return errors.New("could not find const declaration")
	}

	ep.findEnum(constDeclaration)

	ep.Package = file.Name.Name

	if len(ep.Enums) < 1 {
		return errors.New("no enum found")
	}

	return nil
}

// CreateFile will create a file on disk with the enum found.
func (ep *Parser) CreateFile() error {
	tmpl := template.Must(template.New("").Parse(enumTemplate))
	buf := bytes.Buffer{}

	if err := tmpl.Execute(&buf, ep); err != nil {
		return err
	}

	fileBytes, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	outputFile := filepath.Join(
		filepath.Dir(ep.File),
		fmt.Sprintf("%s.gen.go", strcase.ToSnake(ep.TypeName)),
	)

	if err := ioutil.WriteFile(outputFile, fileBytes, 0644); err != nil {
		return err
	}

	return nil
}

func (ep *Parser) findConstDeclaration(fset *token.FileSet, f *ast.File) *ast.GenDecl {
	for _, decl := range f.Decls {
		// If we're not at the given line number yet, move one!
		if fset.Position(decl.Pos()).Line < ep.LineStart {
			continue
		}

		// Ensure the declaration is a GenDecl.
		gdecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		// Ensure the token is a constant.
		if gdecl.Tok != token.CONST {
			continue
		}

		return gdecl
	}

	return nil
}

// nolint: funclen
func (ep *Parser) findEnum(cd *ast.GenDecl) {
	var (
		iotaValue = 0
		addEnum   = func(name string, value interface{}) {
			enum := Enum{
				String: ep.Format(
					strings.TrimPrefix(name, ep.TrimPrefix),
				),
				Name: name,
			}

			if i, ok := value.(int); ok {
				enum.Int = i
				ep.ValueType = "int"
			}

			if v, ok := value.(string); ok {
				enum.Value = v
				ep.ValueType = "string"
			}

			ep.Enums = append(ep.Enums, enum)
		}
	)

	for _, spec := range cd.Specs {
		vspec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		for _, n := range vspec.Names {
			item, ok := n.Obj.Decl.(*ast.ValueSpec)
			if !ok {
				continue
			}

			// Add non first iota item to enums.
			if item.Type == nil && iotaValue != 0 {
				addEnum(n.Name, iotaValue)
				iotaValue++
			}

			typIdent, ok := item.Type.(*ast.Ident)
			if !ok {
				continue
			}

			typeName := typIdent.Name

			// Set the name the first time we find a type.
			if ep.TypeName == "" {
				ep.TypeName = typeName
			}

			// New type found, move on
			if ep.TypeName != typeName {
				continue
			}

			var (
				value interface{}
				err   error
			)

			switch v := item.Values[0].(type) {
			case *ast.BasicLit:
				switch v.Kind {
				case token.STRING:
					value = strings.Trim(v.Value, "\"")
				case token.INT:
					value, err = strconv.Atoi(v.Value)
					if err != nil {
						continue
					}
				default:
					continue
				}
			case *ast.Ident:
				if v.Name == "iota" {
					value = iotaValue
					iotaValue++
				}
			default:
				continue
			}

			addEnum(n.Name, value)
		}
	}
}
