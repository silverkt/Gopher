/**
所有模块包
*/
package libs

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func MarkDownParser() {
	str := "![avatar](/home/picture/1.png)"
	html := blackfriday.MarkdownCommon([]byte(str))
	fmt.Printf("%s/n", html)
}