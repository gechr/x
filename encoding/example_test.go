package encoding_test

import (
	"fmt"

	xencoding "github.com/gechr/x/encoding"
)

func ExampleNewPath() {
	p := xencoding.NewPath("items").Index(0).Child("foo", "bar").Wildcard()
	fmt.Println(p)
	// Output:
	// items[0].foo.bar[*]
}

// Names that cannot appear in dot notation are bracket-quoted automatically.
func ExamplePath_Child() {
	p := xencoding.NewPath("metadata", "labels").Child("kubernetes.io/hostname")
	fmt.Println(p)
	// Output:
	// metadata.labels["kubernetes.io/hostname"]
}

func ExamplePath_Render() {
	p := xencoding.NewPath("items").Index(0)
	fmt.Println(p.Render())
	fmt.Println(p.Render(xencoding.WithRoot()))
	fmt.Println(p.Render(xencoding.WithRoot('@')))
	// Output:
	// items[0]
	// $.items[0]
	// @.items[0]
}

func ExamplePath_Lookup() {
	doc := map[string]any{"spec": map[string]any{"replicas": 3}}
	v, ok := xencoding.NewPath("spec", "replicas").Lookup(doc)
	fmt.Println(v, ok)
	// Output:
	// 3 true
}

func ExamplePath_LookupAll() {
	doc := map[string]any{"items": []any{
		map[string]any{"name": "a"},
		map[string]any{"name": "b"},
	}}
	names := xencoding.NewPath("items").Wildcard().Child("name").LookupAll(doc)
	fmt.Println(names)
	// Output:
	// [a b]
}
