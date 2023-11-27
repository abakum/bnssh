//go:build !windows && !plan9 && !nacl
// +build !windows,!plan9,!nacl

package bnssh

import "github.com/blacknon/go-sshlib"

func ConnectSshAgent() (ag sshlib.AgentInterface) {
	return sshlib.ConnectSshAgent()
}
