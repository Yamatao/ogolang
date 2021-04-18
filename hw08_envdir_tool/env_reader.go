package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"
)

type Environment map[string]string

// convert to the form: an array of "key=value" strings
func (e Environment) AsArray() []string {
	arr := make([]string, 0, len(e))

	for k, v := range e {
		arr = append(arr, k+"="+v)
	}

	return arr
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dirpath string) (env Environment, retErr error) {
	entries, err := ioutil.ReadDir(dirpath)
	if err != nil {
		retErr = err
		return
	}

	env = make(Environment)

	for _, entry := range entries {
		if !entry.Mode().IsRegular() {
			continue
		}
		// имя `S` не должно содержать `=`
		if strings.Contains(entry.Name(), "=") {
			continue
		}

		value, err := ioutil.ReadFile(path.Join(dirpath, entry.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file '%s': %v\n", entry.Name(), err)
			continue
		}

		// пробелы и табуляция в конце `T` удаляются;
		svalue := strings.TrimRight(string(value), " \t")

		// take only the first line
		if nlpos := strings.Index(svalue, "\n"); nlpos != -1 {
			svalue = svalue[:nlpos]
		}

		// терминальные нули (`0x00`) заменяются на перевод строки (`\n`);
		svalue = strings.ReplaceAll(svalue, "\000", "\n")

		// ignore lines with non-conforming characters
		if strings.IndexFunc(svalue, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsPunct(c) && c != ' ' && c != '\n'
		}) != -1 {
			continue
		}

		env[entry.Name()] = svalue
	}
	return
}
