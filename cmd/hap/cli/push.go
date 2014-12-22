// Hap - the simple and effective provisioner
// Copyright (c) 2014 Garrett Woodworth (https://github.com/gwoo)
// The BSD License http://opensource.org/licenses/bsd-license.php.
package cli

import (
	"fmt"

	"github.com/gwoo/hap"
)

// Add the Push command
func init() {
	Commands.Add("push", &PushCmd{})
}

// Push command
type PushCmd struct {
	result []byte
	log    string
}

// Return the result of the command
func (cmd *PushCmd) String() string {
	return string(cmd.result)
}

// Return the log generated by the command
func (cmd *PushCmd) Log() string {
	return cmd.log
}

// Get help on the push command
func (cmd *PushCmd) Help() string {
	return "hap push\t\t\tPush current repo to the remote."
}

// Push to the remote
func (cmd *PushCmd) Run(remote *hap.Remote) error {
	result, err := remote.Push()
	cmd.result = result
	if err != nil {
		cmd.log = fmt.Sprintf("Failed to push to %s.", remote.Host.Addr)
		return err
	}
	if len(cmd.result) <= 0 {
		cmd.result = []byte(fmt.Sprintf("[%s] Push successful.\n", remote.Host.Name))
	}
	cmd.log = fmt.Sprintf("Pushed to %s.", remote.Host.Addr)
	return nil
}