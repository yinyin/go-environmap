package main

import (
	"encoding/json"
	"log"
	"os"

	environmap "github.com/yinyin/go-environmap"
)

func main() {
	envMap := environmap.ParseEnviron(os.Environ())
	if buf, err := json.MarshalIndent((map[string]string)(envMap), "  ", "  "); nil != err {
		log.Fatalf("ERR: encode result failed: %v", err)
	} else {
		log.Printf("INFO: result: %s", string(buf))
	}
	return
}
