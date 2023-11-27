package bnssh

import (
	"github.com/blacknon/go-sshlib"
	"golang.org/x/crypto/ssh"
)

// CreateAuthMethodAgent returns ssh.AuthMethod from con.Agent.
// case con.Agent is nil then ConnectSshAgent to it
func CreateAuthMethodAgent(con *sshlib.Connect) (auth ssh.AuthMethod, err error) {
	var ag sshlib.AgentInterface
	if con == nil {
		ag = ConnectSshAgent()
	} else {
		if con.Agent == nil {
			ag = ConnectSshAgent()
			con.Agent = ag
		} else {
			ag = con.Agent
		}
	}
	signers, err := sshlib.CreateSignerAgent(ag)
	if err != nil {
		return
	}
	auth = ssh.PublicKeys(signers...)

	return
}
