package main

import "fmt"

func main()  {
	s := "()[]{}"
	res :=	isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {
	if len(s)%2 != 0{
		return false
	}
	stack := make([]byte,0,len(s)%2)
	for k,v := range s{
		if v == '(' || v == '[' || v == '{'{
			switch v {
			case '(':
				stack = append(stack,')')
			case '[':
				stack = append(stack, ']')
			case '{':
				stack = append(stack,'}')
			}
		}else{
			if len(stack) == 0{
				return false
			}
			front := stack[len(stack) -1]
			stack = stack[:len(stack)-1]
			if front != s[k]{
				return false
			}
		}
	}
	return len(stack) == 0
}
