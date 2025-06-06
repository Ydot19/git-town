package flags

import (
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/spf13/cobra"
)

const noPushLong = "no-push"

// type-safe access to the CLI arguments of type configdomain.PushBranches
func NoPush() (AddFunc, ReadNoPushFlagFunc) {
	addFlag := func(cmd *cobra.Command) {
		cmd.Flags().BoolP(noPushLong, "", false, "do not push local branches")
	}
	readFlag := func(cmd *cobra.Command) (configdomain.PushBranches, error) {
		value, err := cmd.Flags().GetBool(noPushLong)
		return configdomain.PushBranches(!value), err
	}
	return addFlag, readFlag
}

// ReadNoPushFlagFunc is the type signature for the function that reads the "no-push" flag from the args to the given Cobra command.
type ReadNoPushFlagFunc func(*cobra.Command) (configdomain.PushBranches, error)
