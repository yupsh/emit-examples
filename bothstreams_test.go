package emit_test

import (
	"bytes"
	"fmt"

	emit "github.com/yupsh/emit"
)

func ExampleEmit_bothStreams() {
	// emit to both stdout and stderr
	// Note: in pipelines, stderr doesn't flow to the next command
	var stdout, stderr bytes.Buffer

	cmd := emit.Emit("output message", "error message")
	err := cmd.Executor()(nil, nil, &stdout, &stderr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("stdout: %s", stdout.String())
	fmt.Printf("stderr: %s", stderr.String())
	// Output:
	// stdout: output message
	// stderr: error message
}

