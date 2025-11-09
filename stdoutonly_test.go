package emit_test

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/emit"
)

func ExampleEmit_stdoutOnly() {
	// emit only to stdout
	gloo.MustRun(
		emit.Emit("stdout only", ""),
	)
	// Output:
	// stdout only
}
