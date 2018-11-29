package libs

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func MarkDownParser() {
	str := "# markdown"
	html := blackfriday.MarkdownCommon([]byte(str))
	fmt.Printf("%s/n", html)
}