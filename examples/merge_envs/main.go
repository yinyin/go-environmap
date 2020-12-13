package main

import (
	"encoding/json"
	"log"
	"os"

	environmap "github.com/yinyin/go-environmap"
)

func loadMapJSON(jsonFilePath string) (m map[string]string) {
	fp, err := os.Open(jsonFilePath)
	if nil != err {
		log.Fatalf("ERR: open JSON failed [%s]: %v", jsonFilePath, err)
	}
	defer fp.Close()
	dec := json.NewDecoder(fp)
	if err = dec.Decode(&m); nil != err {
		log.Fatalf("ERR: parse JSON failed [%s]: %v", jsonFilePath, err)
	}
	return
}

func logEnvMap(subjectTitle string, envMap environmap.EnvironMap) {
	if buf, err := json.MarshalIndent((map[string]string)(envMap), "  ", "  "); nil != err {
		log.Fatalf("ERR: encode %s failed: %v", subjectTitle, err)
	} else {
		log.Printf("INFO: %s: %s", subjectTitle, string(buf))
	}
}

func main() {
	envMap := (environmap.EnvironMap)(loadMapJSON(os.Args[1]))
	logEnvMap("base", envMap)
	for _, f := range os.Args[2:] {
		m := loadMapJSON(f)
		envMap.Merge(m)
		logEnvMap("applied "+f, envMap)
	}
	return
}
