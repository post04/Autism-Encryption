# Autism-Encryption

```
go get github.com/postrequest69/Autism-Encryption
```

```go
package main

import(
      "log"
      "os"
      
      "github.com/postrequest69/Autism-Encryption"
)

func main() {
//This is what will be encrypted
var string = "test string"
//This is your string encrypted returned in bytes. To save it in a json or something do string(autismBytes)
autismBytes := autism.AutismEncrypt(string)
//Optional
f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
		f.Close()
	}
	f.WriteString(string(autismBytes))
	f.Close()
}
```

Input: test string

Output: 🍁💸💸🧾📪🍀💷💳🍁🍃🍀🍀🧾🍃💵📦💴✏️🧾✏️💵💶🍂💳💸☘️💷💷🍂💷🍀☘️💷🍂💴🍀💴💴📪📦🍃💷🌿🍀💷💸🧾☘️🌿📦🍃📦💴🧾💷💶🍃💶💵🍂📪💶✏️☘️
