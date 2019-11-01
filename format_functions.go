package enum

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func FormatFuncs() map[string]func(s string) string {
	return map[string]func(s string) string{
		"snake":            strcase.ToSnake,
		"camel":            strcase.ToLowerCamel,
		"upper":            strings.ToUpper,
		"lower":            strings.ToLower,
		"first":            FirstLetter,
		"first-upper":      FirstLetterUpper,
		"first-lower":      FirstLetterLower,
		"capitalize-first": CapitalizeFirst,
		"capitalize-all":   CapitalizeAll,
	}
}

func FirstLetter(s string) string {
	return s[:1]
}

func FirstLetterUpper(s string) string {
	return strings.ToUpper(FirstLetter(s))
}

func FirstLetterLower(s string) string {
	return strings.ToLower(FirstLetter(s))
}

func CapitalizeFirst(s string) string {
	snakeForm := strcase.ToSnake(s)
	capitalized := strings.Title(snakeForm)

	return strings.ReplaceAll(capitalized, "_", " ")
}

func CapitalizeAll(s string) string {
	snakeForm := strcase.ToSnake(s)
	separated := strings.ReplaceAll(snakeForm, "_", " ")

	return strings.Title(separated)
}
