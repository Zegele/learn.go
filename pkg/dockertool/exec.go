package dockertool

import (
	"bytes"
	"context"
	_ "github.com/docker/docker/api/types"
	_ "github.com/docker/docker/client"
	_ "github.com/docker/docker/pkg/stdcopy"
)

// ExecResult repressents a result returned from Exec()
type ExecResult struct {
	ExitCode  int
	outBuffer *bytes.Buffer
	errBuffer *bytes.Buffer
}

// Stdout returns stdout output of a command run by Exec()
func (res *ExecResult) Stdout() string {
	return res.outBuffer.String()
}

// Stderr returns stderr output of a command run by Exec()
func (res *ExecResult) Stderr() string {
	return res.errBuffer.String()
}

// Combined returns combined stdout and stderr output of a command run by Exec()
func (res *ExecResult) Combined() string {
	return res.outBuffer.String() + res.errBuffer.String()
}

func Exec(ctx context.Context, cli client.APIClient, id string, cmd []string) (ExecResult, error) {
	// prepare exec
	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	}
	cresp, err := cli.ContainnerExecCreate(ctx, id, execConfig)
	if err != nil {
		return ExecResult{}, err
	}
	execID := cresp.ID

	// run it, with stdout/stderr attached
	aresp, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer aresp.Close()

	// read th e output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		// StdCopy demultiplexes the stream into two buffers
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, aresp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break

	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	//get the exit code
	iresp, err := cli.ContainerExecInspect(ctx, execID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{ExitCode: iresp.ExitCode, outBuffer: &outBuf, errBuffer: &errBuf}, nil
}
