package autism

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func AutismEncrypt(text string) []byte {
	shatext := fmt.Sprintf("%X", sha256.Sum256([]byte(text)))
	contents, err := ioutil.ReadFile("./key.json")
	if err != nil {
		fmt.Println("Had an err: ", err)
	}

	var key map[string]interface{}
	errr := json.Unmarshal([]byte(contents), &key)
	if errr != nil {
		fmt.Println("Had an err: ", errr)
	}
	var output = ""
	for _, char := range shatext {
		autism := key[string(char)]
		byteautism, errrr := json.Marshal(autism)
		if errrr != nil {
			fmt.Println("Had an err: ", errrr)
		}
		stringautism := string(byteautism)
		if stringautism == "null" {
			fmt.Println("NULL for ", string(char))
		}
		output += strings.ReplaceAll(stringautism, "\"", "")
	}
	return []byte(output)
}
