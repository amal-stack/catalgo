package cmd

import "io"

type Env struct {
	IO *IOStreams
	
}

type IOStreams struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer

	IsInTTY  bool
	IsOutTTY bool
	IsErrTTY bool
}