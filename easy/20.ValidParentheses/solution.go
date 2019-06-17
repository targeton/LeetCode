package solution

import (
	"container/list"
	"strings"
)

func isValid(s string) bool {
	strs := strings.Split(s, "")
	l := list.New()
	for _, str := range strs {
		if str == "(" {
			l.PushBack(")")
		} else if str == "[" {
			l.PushBack("]")
		} else if str == "{" {
			l.PushBack("}")
		} else {
			if l.Len() < 1 {
				return false
			}
			last := l.Back()
			if last.Value.(string) == str {
				l.Remove(last)
			} else {
				return false
			}
		}
	}
	return !(l.Len() > 0)
}
