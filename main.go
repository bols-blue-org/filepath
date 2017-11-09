package main

import (
	"os"
)

func main() {
	testCases := map[string]bool{
		"dir":                          false,
		"file.txt":                     false,
		".dotfile":                     false,
		"link-to-file.txt":             false,
		"unknown-dir":                  true,
		"unknown-file":                 true,
		".unknown-dotfile":             true,
		"unknown-link-to-file.txt":     true,
		"dir/unknown-dir":              true,
		"file.txt/unknown-dir":         false,
		".dotfile/unknown-dir":         false,
		"link-to-file.txt/unknown-dir": false,
	}

	for filename, expects := range testCases {
		if fileNotExist("/tmp/golang-test/"+filename) == expects {
			println("PASS: " + filename)
		} else {
			println("FAIL: " + filename)
		}
	}
	for filename, expects := range testCases {
		if fileNotExist2("/tmp/golang-test/"+filename) == expects {
			println("PASS: " + filename)
		} else {
			println("FAIL: " + filename)
		}
	}
}

func fileNotExist(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return true
	}

	return false
}

func fileNotExist2(filename string) bool {
	_, err := os.Stat(filename)
	return err != nil
}
