// Hap - the simple and effective provisioner
// Copyright (c) 2015 Garrett Woodworth (https://github.com/gwoo)
// The BSD License http://opensource.org/licenses/bsd-license.php.

package cli

import (
	"fmt"

	"github.com/gwoo/hap"
)

// Add the init command
func init() {
	Commands.Add("init", &InitCmd{})
}

// Init command for setting up remote repo
type InitCmd struct {
	result []byte
	log    string
}

// Does this command expect a remote
func (cmd *InitCmd) IsRemote() bool {
	return true
}

// Return the result of the command
func (cmd *InitCmd) String() string {
	return string(cmd.result)
}

// Return the log generated by the command
func (cmd *InitCmd) Log() string {
	return cmd.log
}

// Get help on the init command
func (cmd *InitCmd) Help() string {
	return "hap init\tInitialize a new remote host."
}

// Run the command against the remote
func (cmd *InitCmd) Run(remote *hap.Remote) error {
	result, err := remote.Initialize()
	cmd.result = result
	if err != nil {
		cmd.log = fmt.Sprintf("%s failed to initialize on %s.", remote.Dir, remote.Host.Addr)
		return err
	}
	if len(cmd.result) <= 0 {
		cmd.result = []byte(fmt.Sprintf("[%s] Init successful.\n", remote.Host.Name))
	}
	cmd.log = fmt.Sprintf("%s initialized on %s.", remote.Dir, remote.Host.Addr)
	return nil
}
