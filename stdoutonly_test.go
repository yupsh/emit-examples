package emit_test

import (
	emit "github.com/yupsh/emit"
	yup "github.com/gloo-foo/framework"
)

func ExampleEmit_stdoutOnly() {
	// emit only to stdout
	yup.MustRun(
		emit.Emit("stdout only", ""),
	)
	// Output:
	// stdout only
}

