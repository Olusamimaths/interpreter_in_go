package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIndentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}

func tracePrint(s string) {
	fmt.Printf("%s%s", identLevel(), s)
}

func incIdent()  { traceLevel += 1 }
func decIndent() { traceLevel -= 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIndent()
}
