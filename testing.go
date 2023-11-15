package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func testing() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	keyMap := make(map[string]string)
	for _, key := range cfg.Section("SIMPLEFAZ").Keys() {
		keyMap[key.Name()] = key.Value()
	}

	for k, v := range keyMap {
		fmt.Println(k, v)
	}
}

func main() {
	testing()
}
