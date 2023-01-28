//go:build clientlogmode
// +build clientlogmode

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var config = struct {
	ModeBits []string
}{
	// Items should be appended only to keep bit-flag positions stable
	ModeBits: []string{
		"Signing",
		"Retries",
		"Request",
		"RequestWithBody",
		"Response",
		"ResponseWithBody",
		"DeprecatedUsage",
		"RequestEventMessage",
		"ResponseEventMessage",
	},
}

func bitName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

var tmpl = template.Must(template.New("ClientLogMode").Funcs(map[string]interface{}{
	"symbolName": func(name string) string {
		return "Log" + bitName(name)
	},
	"bitName": bitName,
}).Parse(`// Code generated by aws/logging_generate.go DO NOT EDIT.

package aws

// ClientLogMode represents the logging mode of SDK clients. The client logging mode is a bit-field where
// each bit is a flag that describes the logging behavior for one or more client components.
// The entire 64-bit group is reserved for later expansion by the SDK.
//
// Example: Setting ClientLogMode to enable logging of retries and requests
//  clientLogMode := aws.LogRetries | aws.LogRequest
//
// Example: Adding an additional log mode to an existing ClientLogMode value
//  clientLogMode |= aws.LogResponse
type ClientLogMode uint64

// Supported ClientLogMode bits that can be configured to toggle logging of specific SDK events.
const (
{{- range $index, $field := .ModeBits }}
	{{ (symbolName $field) }}{{- if (eq 0 $index) }} ClientLogMode = 1 << (64 - 1 - iota){{- end }}
{{- end }}
)
{{ range $_, $field := .ModeBits }}
// Is{{- bitName $field }} returns whether the {{ bitName $field }} logging mode bit is set
func (m ClientLogMode) Is{{- bitName $field }}() bool {
	return m&{{- (symbolName $field) }} != 0
}
{{ end }}
{{- range $_, $field := .ModeBits }}
// Clear{{- bitName $field }} clears the {{ bitName $field }} logging mode bit
func (m *ClientLogMode) Clear{{- bitName $field }}() {
	*m &^= {{ (symbolName $field) }}
}
{{ end -}}
`))

func main() {
	uniqueBitFields := make(map[string]struct{})

	for _, bitName := range config.ModeBits {
		if _, ok := uniqueBitFields[strings.ToLower(bitName)]; ok {
			panic(fmt.Sprintf("duplicate bit field: %s", bitName))
		}
		uniqueBitFields[bitName] = struct{}{}
	}

	file, err := os.Create("logging.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, config)
	if err != nil {
		log.Fatal(err)
	}
}
