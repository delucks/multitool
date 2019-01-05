package main

import (
	"strings"
	"testing"
)

func TestYamlToJson(_ *testing.T) {
	r := strings.NewReader("a:\n  b:\n  - item\n  - item2\n")
	YamlToJson(nil, r)
	// Output:
	// {"a": {"b": ["item", "item2"]}}
}

func TestJsonToYaml(_ *testing.T) {
	r := strings.NewReader(`{"my": {"cool": ["json", "object"]}}`)
	JsonToYaml(nil, r)
	// Output:
	// my:\n  cool:\n  - json\n  - object\n
}
