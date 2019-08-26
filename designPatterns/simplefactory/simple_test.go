package simplefactory

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("Tom")
	fmt.Println(s)
	if s == "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}
