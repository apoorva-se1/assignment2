package main

import (
	"fmt"
)

 var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe := 0
	fmt.Println(WhatIsThe)
}

func main() {
	if WhatIsThe == 42 {
		fmt.Println("It's all a lie.")
	}
}
