package autism

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func AutismEncrypt(text string) []byte {
	shatext := fmt.Sprintf("%X", sha256.Sum256([]byte(text)))
	// TODO: add all emojis
	var key map[string]string = map[string]string{
		"a": "😜",
		"b": "🤪",
		"c": "",
		"d": "", //...
	}
	var output = ""
	for _, char := range shatext {
		autistiChar := key[string(char)]
		output += strings.ReplaceAll(autistiChar, "\"", "")
	}
	return []byte(output)
}
