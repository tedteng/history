/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package history

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type History struct {
	ConfigPath string
}

func (h *History) Settings(path string) *History {
	return &History{ConfigPath: path}
}

func (h *History) Write(i interface{}) {
	f, err := os.OpenFile(h.ConfigPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("History file write error ", h.ConfigPath)
	}
	defer f.Close()
	if _, err := f.WriteString(convert(i)); err != nil {
		log.Fatalln("History file write error ", h.ConfigPath)
	}
}

func (h *History) Load() []string {
	f, err := os.Open(h.ConfigPath)
	if err != nil {
		log.Fatalln("cannot open the file", err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	var items []string

	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Load err:", err)
		}
		items = append(items, string(line))
	}
	return items
}

func (h *History) Previous() []string {
	load := h.Load()

	if len(load) <= 1 {
		return nil
	}

	item := make([]string, 1)
	copy(item, load[len(load)-2:len(load)-1])

	return item
}

func convert(i interface{}) string {
	var str string
	switch x := i.(type) {
	case []string:
		str = strings.Join(x, " ")
	case []byte:
		str = string(x[:])
	/*
		add more type assertion if need
	*/

	default:
		log.Fatalln("type unknown") // here v has type interface{}
	}
	return isLineBreak(str)
}

func isLineBreak(str string) string {
	if strings.HasSuffix(str, "\n") {
		return str
	}
	return str + "\n"
}
