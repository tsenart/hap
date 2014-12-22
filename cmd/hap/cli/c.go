// Hap - the simple and effective provisioner
// Copyright (c) 2014 Garrett Woodworth (https://github.com/gwoo)
// The BSD License http://opensource.org/licenses/bsd-license.php.
package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gwoo/hap"
)

// Load all the available commands
func init() {
	Commands.Add("c", &ArbitraryCmd{})
}

type ArbitraryCmd struct {
	result []byte
	log    string
}

// Return the result of the command
func (cmd *ArbitraryCmd) String() string {
	return string(cmd.result)
}

// Return the log generated by the command
func (cmd *ArbitraryCmd) Log() string {
	return cmd.log
}

// Get help on c (arbitrary) command
func (cmd *ArbitraryCmd) Help() string {
	return "hap c <command>\t\t\tRun an arbitrary command on the remote host."
}

// Run an arbitrary command on the remote host
func (cmd *ArbitraryCmd) Run(remote *hap.Remote) error {
	args := flag.Args()
	if len(args) <= 1 {
		return fmt.Errorf("%s", cmd.Help())
	}
	arbitrary := strings.Join(args[1:], " ")
	result, err := remote.Execute([]string{arbitrary})
	cmd.result = result
	cmd.log = fmt.Sprintf("Executed `%s` on %s.", arbitrary, remote.Host.Addr)
	return err
}
