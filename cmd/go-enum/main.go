package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/bombsimon/enum"
)

func main() {
	var (
		fileName   = kingpin.Arg("input-file", "Uses $GOFILE if that is set.").Envar("GOFILE").Required().ExistingFile()
		lineNum    = kingpin.Flag("line", "Location of const statement to generate enum from. Uses $GOLINE if that is set.").Default("1").Envar("GOLINE").Int()
		trimPrefix = kingpin.Flag("trim", "Prefix to trim from the enum type name when generating the strings.").Default("").String()
		formatFunc = kingpin.Flag("format", "How to format string value").Default("snake").Enum("space", "snake", "camel", "upper", "lower", "first", "first-upper", "first-lower", "capitalize-first", "capitalize-all")
		json       = kingpin.Flag("json", "Generate code implementing (un)marshal interface").Default("true").Bool()
		value      = kingpin.Flag("with-value", "Generate code implementing Value() to allow the actual value").Default("false").Bool()
	)

	kingpin.Parse()

	formatFuncs := enum.FormatFuncs()

	e := enum.New(*fileName, *trimPrefix, *lineNum, *json, *value, formatFuncs[*formatFunc])

	if err := e.GetEnumFromFile(); err != nil {
		fmt.Printf("Could not get enums: %s\n", err.Error())
		os.Exit(1)
	}

	if err := e.CreateFile(); err != nil {
		fmt.Printf("Could not create enum file: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully generated enums!")
}
