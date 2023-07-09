package cipher

import (
	"fmt"
	"testing"
)

func TestRc4(t *testing.T) {
	text := "test123"
	key := "abcd1234"
	desText, err := Rc4(text, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(desText)
	//oriText, err := Rc4(desText, key)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(oriText)
}
