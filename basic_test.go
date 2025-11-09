package emit_test

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/emit"
)

func ExampleEmit_basic() {
	// emit "Hello" to stdout and "Warning" to stderr
	gloo.MustRun(
		emit.Emit("Hello", "Warning"),
	)
	// Output:
	// Hello
}
