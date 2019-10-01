package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/bombsimon/enum"
	"github.com/iancoleman/strcase"
)

func main() {
	var (
		fileName   = kingpin.Arg("input-file", "Uses $GOFILE if that is set.").Envar("GOFILE").Required().ExistingFile()
		lineNum    = kingpin.Flag("line", "Location of const statement to generate enum from. Uses $GOLINE if that is set.").Default("1").Envar("GOLINE").Int()
		trimPrefix = kingpin.Flag("trim", "Prefix to trim from the enum type name when generating the strings.").Default("").String()
		formatFunc = kingpin.Flag("format", "snake|camel|upper|lower").Default("snake").Enum("snake", "camel", "upper", "lower")
		json       = kingpin.Flag("json", "Generate code implementing (un)marshal interface").Default("true").Bool()
	)

	kingpin.Parse()

	formatFuncs := map[string]func(s string) string{
		"snake": strcase.ToSnake,
		"camel": strcase.ToLowerCamel,
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}

	e := enum.New(*fileName, *trimPrefix, *lineNum, *json, formatFuncs[*formatFunc])

	if err := e.GetEnum(); err != nil {
		fmt.Printf("Could not get enums: %s\n", err.Error())
		os.Exit(1)
	}

	if err := e.CreateFile(); err != nil {
		fmt.Printf("Could not create enum file: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully generated enums!")
}
