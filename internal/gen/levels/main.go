//go:generate go run . -output ../../echologger/level_gen.go
package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"

	"golang.org/x/tools/imports"
)

type (
	Bindings struct {
		Levels      []Level
		PackageName string
	}

	Level struct {
		LevelName  string
		MethodName string
	}
)

var (
	filename    = flag.String("output", "gen.go", "output file name")
	packageName = flag.String("name", "echologger", "package name")
	methodNames = []string{
		"Debug",
		"Info",
		"Warn",
		"Error",
		"Fatal",
		"Panic",
		"Print",
	}
	levelDict = map[string]string{
		"Print": "No",
	}
)

func main() {
	flag.Parse()

	bind := NewBindings()
	cont := content(bind)

	if err := render(cont); err != nil {
		log.Fatal(err)
	}
}

func NewBindings() Bindings {
	data := make([]Level, len(methodNames))
	for i, method := range methodNames {
		if levelName, ok := levelDict[method]; ok {
			data[i] = Level{
				LevelName:  levelName,
				MethodName: method,
			}
		} else {
			data[i] = Level{
				LevelName:  method,
				MethodName: method,
			}
		}
	}
	return Bindings{
		Levels:      data,
		PackageName: *packageName,
	}
}

func content(bind Bindings) []byte {
	const cont = `// Code generated. DO NOT EDIT.

package {{ .PackageName }}
{{ range .Levels }}
func (l Logger) {{ .MethodName }}(i ...interface{}) {
	l.msg(zerolog.{{ .LevelName }}Level, i...)
}

func (l Logger) {{ .MethodName }}f(f string, i ...interface{}) {
	l.msgf(zerolog.{{ .LevelName }}Level, f, i...)
}

func (l Logger) {{ .MethodName }}j(j log.JSON) {
	l.fields(zerolog.{{ .LevelName }}Level, j)
}
{{ end }}
`

	var b bytes.Buffer

	t := template.Must(template.New("content").Parse(cont))

	if err := t.Execute(&b, bind); err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}

func render(cont []byte) error {
	imported, err := imports.Process(*filename, cont, nil)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(*filename, imported, 0644)
}
