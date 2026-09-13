package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// var rootCmd = &cobra.Command{
// 	Use:   "catalgo",
// 	Short: "A CLI tool for managing algorithm problems",
// 	Long:  `catalgo is a command-line interface (CLI) tool designed to help users manage algorithm problems, solutions, and implementations. It provides commands for creating new problems, adding solutions, and managing implementations across different programming languages.`,
// 	Run:   runRoot,
// }

type Invocation[T any] struct {
	Cmd     *cobra.Command
	Options T
	Args    []string
}

type RunFunc[T any] func(ctx *Invocation[T]) error

func NewCommand[T any](parseOptions func(*Invocation[T]) error, run RunFunc[T]) *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := &Invocation[T]{
				Cmd:  cmd,
				Args: args,
			}
			err := parseOptions(ctx)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
			return run(ctx)
		},
	}

	return cmd
}

func NewCmdRoot() *cobra.Command {
	cmd := NewCommand(parseRootOptions, runRoot)

	return cmd
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catalgo",
		Short: "Catalog and manage your alogrithmic problem solving knowledge",
		Long: `Catalgo is a tool for managing a personal knowledge-base for alogrithmic problem solving.

		It organizes problems, solutions, implementations, notes and metadata into a structured collection that 
		lives in a simple unified workspace rather than being scattered across multiple websites. It acts as a single source of truth
		to track problems and multiple variants of their associated solutions, associate them with metadata like patterns, categories or
		tags that helps explore patterns, deepen understanding and establish relationships between problems.  
		`,
		Args: cobra.NoArgs,
	}

	return cmd;
}

type RootOptions struct {
	// Add any fields you need for root command options here
}

func parseRootOptions(ctx *Invocation[RootOptions]) error {
	// Parse and set ctx.Options fields
	return nil
}

func invocation(cmd *cobra.Command, args []string) *Invocation[RootOptions] {
	return &Invocation[RootOptions]{
		Cmd:     cmd,
		Options: RootOptions{},
		Args:    args,
	}
}

func runRoot(ctx *Invocation[RootOptions]) error {
	return nil
}

func Execute() {
	// if err := rootCmd.Execute(); err != nil {
	// 	os.Exit(1)
	// }
}

// ====


