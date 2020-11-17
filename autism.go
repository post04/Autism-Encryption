package autism

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

func AutismEncrypt(text string) []byte {
	shatext := fmt.Sprintf("%X", sha256.Sum256([]byte(text)))
	var key map[string]interface{}
	te := []byte("{\"a\": \"😜\",\"b\": \"🤪\",\"c\": \"😋\",\"d\": \"😛\",\"e\": \"😍\",\"f\": \"💯\",\"g\": \"💦\",\"h\": \"💣\",\"i\": \"🌲\",\"j\": \"🏔️\",\"k\": \"🌨️\",\"l\": \"❄️\",\"m\": \"☃️\",\"n\": \"⛄\",\"o\": \"🔥\",\"p\": \"💧\",\"q\": \"🎆\",\"r\": \"✨\",\"s\": \"🎉\",\"t\": \"🎁\",\"u\": \"🏆\",\"v\": \"🃏\",\"w\": \"♟️\",\"x\": \"👕\",\"y\": \"🥾\",\"z\": \"💰\",\"1\": \"💵\",\"2\": \"💴\",\"3\": \"💶\",\"4\": \"💷\",\"5\": \"💸\",\"6\": \"💳\",\"7\": \"🧾\",\"8\": \"📦\",\"9\": \"📪\",\"0\": \"✏️\",\"A\": \"🌿\",\"B\": \"☘️\",\"C\": \"🍀\",\"D\": \"🍁\",\"E\": \"🍂\",\"F\": \"🍃\",\"G\": \"🍇\",\"H\": \"🍈\",\"I\": \"🍉\",\"J\": \"🍊\",\"K\": \"🍋\",\"L\": \"🍌\",\"M\": \"🍍\",\"N\": \"🥭\",\"O\": \"🍎\",\"P\": \"🍏\",\"Q\": \"🍑\",\"R\": \"🍐\",\"S\": \"🍒\",\"T\": \"🍓\",\"U\": \"🥝\",\"V\": \"🍅\",\"W\": \"🥥\",\"X\": \"🥑\",\"Y\": \"🍆\",\"Z\": \"🥔\"}")
	errr := json.Unmarshal(te, &key)
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
