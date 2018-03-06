/**
testa is a package for my test
*/
package testa
import (
	"fmt";
	"strconv"
)

const Pi float64 = 3.1415926
var testa int = 12
var teststring string = "this is a good world"
func Tt() {
	fmt.Println(strconv.Itoa(testa) + teststring + strconv.FormatFloat(Pi, 'E', -1, 64))
	for i := 0; i < 20; i++ {
		fmt.Println("this is round" + strconv.Itoa(i))
		if i == 12 {
			fmt.Println("1212 is a good girl")
		}
	}
}