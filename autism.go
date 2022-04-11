package autism

import (
	"crypto/sha256"
	"encoding/json"
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
		autism := key[string(char)]
		byteautism, err := json.Marshal(autism)
		if err != nil {
			fmt.Println("Had an err: ", err)
		}
		stringautism := string(byteautism)
		if stringautism == "null" {
			fmt.Println("NULL for ", string(char))
		}
		output += strings.ReplaceAll(stringautism, "\"", "")
	}
	return []byte(output)
}
