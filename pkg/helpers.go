package ipcheck

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ProcessList(outputFormat string, args *[]string) (b *bytes.Buffer) {
	var a Accumulator
	switch outputFormat {
	case "j", "json":
		a = NewJsonAccumulator()
	default:
		a = NewTextAccumulator()
	}

	for _, input := range *args {
		ip, err := Type(input)

		if err != nil {
			a.Append(input, err.Error())
			fmt.Println(err)
		}
		a.Append(input, ip.Name)
	}
	return a.Read()
}

type Accumulator interface {
	Append(ip string, tName string)
	Read() *bytes.Buffer
}

type TextAccumulator struct {
	buffer *bytes.Buffer
}

func NewTextAccumulator() *TextAccumulator {
	b := bytes.NewBufferString("")
	return &TextAccumulator{b}
}

func (a TextAccumulator) Append(ip string, tName string) {
	a.buffer.WriteString(tName)
	a.buffer.WriteString("\n")
}

func (a TextAccumulator) Read() *bytes.Buffer {
	return a.buffer
}

type JsonAccumulator struct {
	m map[string]string
}

func NewJsonAccumulator() *JsonAccumulator {
	m := make(map[string]string)
	return &JsonAccumulator{m}
}

func (a JsonAccumulator) Append(ip string, tName string) {
	a.m[ip] = tName
}

func (a JsonAccumulator) Read() *bytes.Buffer {
	r, _ := json.Marshal(a.m)
	return bytes.NewBuffer(r)
}
