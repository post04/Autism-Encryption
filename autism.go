package autism

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Why expect a string and return a byte?
func AutismEncrypt(text string) []byte {
	// What does this even mean?
	shatext := fmt.Sprintf("%X", sha256.Sum256([]byte(text)))
	// TODO: add all emojis (see git history & my commit messages)
	var key map[string]string = map[string]string{
		"a": "😜",
		"b": "🤪",
		"c": "",
		"d": "", //...
	}
	var output = ""
	// this could be reduced even more
	for _, char := range shatext {
		autistiChar := key[string(char)]
		output += strings.ReplaceAll(autistiChar, "\"", "")
	}
	// Should use strings
	return []byte(output)
}
