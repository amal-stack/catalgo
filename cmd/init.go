package cmd

import "github.com/spf13/cobra"


type InitOptions struct {
	Path string
	
}

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create a new Catalgo workspace",
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}
