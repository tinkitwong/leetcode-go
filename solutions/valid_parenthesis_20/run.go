package valid_parenthesis_20

import "fmt"

func Run() {
	s := "[{()]"
	res := isValid(s)
	fmt.Println(res)
}
