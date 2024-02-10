// Copyright (c) 2021 Blacknon. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Shell connection Example file.
// Change the value of the variable and compile to make sure that you can actually connect.
//
// This file has a simple ssh proxy connection.
// Also, the authentication method is password authentication.
// Please replace as appropriate.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abakum/bnssh"
	"github.com/blacknon/go-sshlib"
	"golang.org/x/crypto/ssh"
)

var (
	// http proxy server
	// httpProxyHost = "http-proxy.com"
	// httpProxyPort = "4444"

	httpProxyHost = "127.0.0.1"
	httpProxyPort = "8888"

	// socks5 proxy server
	socks5ProxyHost = "127.0.0.1"
	socks5ProxyPort = "10080"

	// ssh target server
	// host     = "target.com"
	// port     = "22"
	host = "10.161.115.160"
	port = "22"
	user = "root"
)

func main() {
	// proxy := goproxy.NewProxyHttpServer()
	// proxy.Verbose = true
	// go log.Fatal(http.ListenAndServe(httpProxyHost+":"+httpProxyPort, proxy))
	// time.Sleep(time.Second * 7)
	// ==========
	// http proxy connect
	// ==========

	httpProxy := &sshlib.Proxy{
		Type: "http",
		Addr: httpProxyHost,
		Port: httpProxyPort,
	}
	httpProxyDialer, err := httpProxy.CreateProxyDialer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	socks5Proxy := &sshlib.Proxy{
		Type:      "socks5",
		Addr:      socks5ProxyHost,
		Port:      socks5ProxyPort,
		Forwarder: httpProxyDialer,
	}
	socks5ProxyDialer, err := socks5Proxy.CreateProxyDialer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ==========
	// target connect
	// ==========

	log.Println(httpProxyDialer)
	// Create target sshlib.Connect
	targetCon := &sshlib.Connect{
		ProxyDialer: socks5ProxyDialer,
	}

	// Create ssh.AuthMethods
	// store ConnectSshAgent() to con.Agent
	authMethod, err := bnssh.CreateAuthMethodAgent(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Connect target server
	err = targetCon.CreateClient(host, port, user, []ssh.AuthMethod{authMethod})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create Session
	session, err := targetCon.CreateSession()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start ssh shell
	targetCon.Shell(session)
}
