package goback

import (
	"fmt"
)

func Contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func CheckErr(err error){
	if err != nil {
		fmt.Printf("Listen on %v\n", err)
	}
}