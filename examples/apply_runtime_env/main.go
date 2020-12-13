package main

import (
	"encoding/json"
	"log"
	"os"

	environmap "github.com/yinyin/go-environmap"
)

func applyChecker2(envKey, envValue string) (shouldApply bool) {
	return (envValue == ("${" + envKey + "}"))
}

func applyChecker3(envKey, envValue string) (shouldApply bool) {
	return (envValue == ("%" + envKey))
}

func runApply(m map[string]string, fnShouldApply func(envKey, envValue string) (shouldApply bool)) {
	var envMap environmap.EnvironMap = make(map[string]string)
	for k, v := range m {
		envMap[k] = v
	}
	envMap.ApplyRuntimeEnviron(fnShouldApply)
	if buf, err := json.MarshalIndent((map[string]string)(envMap), "  ", "  "); nil != err {
		log.Fatalf("ERR: encode failed: %v", err)
	} else {
		log.Printf("INFO: result: %v", string(buf))
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		log.Print("Usage: ./apply_runtime_env examples/runtime-env.json")
		return
	}
	inputJSONFilePath := os.Args[1]
	fp, err := os.Open(inputJSONFilePath)
	if nil != err {
		log.Fatalf("ERR: load JSON failed [%s]: %v", inputJSONFilePath, err)
	}
	defer fp.Close()
	dec := json.NewDecoder(fp)
	var m map[string]string
	if err = dec.Decode(&m); nil != err {
		log.Fatalf("ERR: parse JSON failed [%s]: %v", inputJSONFilePath, err)
	}
	log.Print("INFO: apply when value is empty")
	runApply(m, nil)
	log.Print("INFO: apply when value is ${KEY}")
	runApply(m, applyChecker2)
	log.Print("INFO: apply when value is %KEY")
	runApply(m, applyChecker3)
}
