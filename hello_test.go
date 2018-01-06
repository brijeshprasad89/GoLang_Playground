package fosscafe

import (
	"fmt"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	var g string
	g = "hello everyone"
	fmt.Println(g)
	if strings.Compare(g, "hello everyone") != 0 {
		t.Error("Expected hello everyone, got", g)
	}
}

func TestGreetingsToMe(t *testing.T) {
	// var g string
	// g = "hello everyone"
	fmt.Println("Saying hi to myself")
	// if(strings.Compare(g,"hello everyone")!= 0){
	// 	t.Error("Expected hello everyone, got",g)
	// }
}
