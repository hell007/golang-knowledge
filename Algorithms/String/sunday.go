package main
 
import "fmt"
 
//字符串匹配
func catch(a, b string) bool {
	i := 0
	j := 0
	for i < len(a) {
		if a[i] == b[j] {
			i++
			j++
			if j == len(b) {
				return true
			}
			continue
		}
		if i+len(b) >= len(a) {
			return false
		}
		k := string(a[i+len(b)])
		find := false
		for key, value := range b {
			if string(value) == k {
				i = i + len(b) - key
				j = 0
				find = true
				break
			}
		}
		if find == false {
			i = i + len(b) + 1
			j = 0
		}
	}
	return false
}
 
func main() {
	a := "abcdef"
	b := "bcd"
	result := catch(a, b)
	fmt.Println(result)
}
