package main

import (
	"github.com/kohdice/work-sandbox/baz/pkg/greet"
	hoge "github.com/kohdice/work-sandbox/hoge/pkg/greet"
)

func main() {
	greet.Greet("foo")
	hoge.Greet("foo")
}
