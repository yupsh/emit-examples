package emit_test

import (
	emit "github.com/yupsh/emit"
	yup "github.com/gloo-foo/framework"
)

func ExampleEmit_basic() {
	// emit "Hello" to stdout and "Warning" to stderr
	yup.MustRun(
		emit.Emit("Hello", "Warning"),
	)
	// Output:
	// Hello
}

