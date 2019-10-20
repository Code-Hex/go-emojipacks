package emojipacks

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestA(t *testing.T) {
	var typ EmojiPacks
	err := yaml.Unmarshal([]byte(data), &typ)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	fmt.Println(typ.Title)
	for _, v := range typ.Emojis {
		fmt.Println("name:", v.Name)
		fmt.Println("src:", v.Src)
		fmt.Println()
	}
}
