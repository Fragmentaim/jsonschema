package jsonschema

import (
	"strings"
	"testing"
)

// No content decoder is enabled, so string validation should not materialize bytes.
func BenchmarkStringWithoutContentAssertions(b *testing.B) {
	c := NewCompiler()
	const resource = "https://test.invalid/string.json"
	if err := c.AddResource(resource, map[string]any{"type": "string"}); err != nil {
		b.Fatal(err)
	}
	schema, err := c.Compile(resource)
	if err != nil {
		b.Fatal(err)
	}
	value := strings.Repeat("A", 4<<20)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := schema.Validate(value); err != nil {
			b.Fatal(err)
		}
	}
}
