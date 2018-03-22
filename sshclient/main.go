package main

import (
	"log"
	"flag"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
	"net"
	"io"
	"sync"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := flag.String("u", "", "user name")
	keyFilePath := flag.String("i", "", "key file path")
	cmd := flag.String("cmd", "", "command")
	forwardAddr := flag.String("forward", "", "remote address to forward")

	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("host address needed")
	}

	hostAddr := flag.Arg(0)

	key, err := ioutil.ReadFile(*keyFilePath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	//var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", hostAddr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	if *forwardAddr != "" {
		fowardToLocalPort(client, *forwardAddr)
	} else if *cmd != "" {
		// Each ClientConn can support multiple interactive sessions,
		// represented by a Session.
		session, err := client.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		defer session.Close()

		// Once a Session is created, you can execute a single command on
		// the remote side using the Run method.
		var b bytes.Buffer
		session.Stdout = &b
		if err := session.Run(*cmd); err != nil {
			log.Fatal("Failed to run: " + err.Error())
		}
		fmt.Println(b.String())
	} else {
		// Each ClientConn can support multiple interactive sessions,
		// represented by a Session.
		session, err := client.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		defer session.Close()

		// Set up terminal modes
		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     // disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}
		// Request pseudo terminal
		if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
			log.Fatal("request for pseudo terminal failed: ", err)
		}

		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		// Start remote shell
		if err := session.Shell(); err != nil {
			log.Fatal("failed to start shell: ", err)
		}

		if err := session.Wait(); err != nil {
			log.Fatal("failed to wait: ", err)
		}
	}
}

func fowardToLocalPort(client *ssh.Client, addr string) {
	localAddr := ":8000"
	log.Println("listen", localAddr)
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Failed to accept: ", err)
		}

		forwardConnToRemote(client, addr, conn)
	}

}
func forwardConnToRemote(client *ssh.Client, addr string, connIn net.Conn) {
	connOut, err := client.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Failed to Dial forward address: ", err)
	}

	defer connOut.Close()
	defer connIn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(connOut, connIn)
	}()

	go func() {
		defer wg.Done()
		io.Copy(connIn, connOut)
	}()

	wg.Wait()
}
