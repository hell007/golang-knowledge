/**
 * name: switch
 * author: jie
 * note:
 * switch sExpr {
		case expr1:
			some instructions
		case expr2:
			some other instructions
		case expr3:
			some other instructions
		default:
			other code
	}
*
* sExpr和expr1、expr2、expr3的类型必须一致。
*
* Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true
*/

package main

import (
	"fmt"
)

func main() {

	// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码

	a := 8

	switch a {
	case 7:
		fmt.Println("a=", 7)
	case 8:
		fmt.Println("a=", 8)
		fallthrough // 强制执行 case 9:
	case 9:
		fmt.Println("a=", 9)
	default:
		fmt.Println("a=", 0)
	}

}
