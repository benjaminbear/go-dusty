package sshhandler

import (
	"io/ioutil"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
	ssh2 "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// GetKey asks a host for PublicKeys in a specific order
// If a key is available the loop is stopped
func GetKey(host string, timeout time.Duration) (KnownHost, error) {
	khEntry := KnownHost{
		Hosts: []string{host},
	}

	supportedHostKeyAlgos := []string{
		ssh.KeyAlgoRSA, ssh.KeyAlgoECDSA521, ssh.KeyAlgoECDSA384, ssh.KeyAlgoECDSA256,
		ssh.KeyAlgoED25519, ssh.KeyAlgoDSA,
	}

	for _, algo := range supportedHostKeyAlgos {
		keyResult, err := getPublicKey(host, algo, timeout)
		if err != nil {
			return khEntry, err
		}

		if keyResult != nil {
			khEntry.Key = keyResult
			break
		}
	}

	furtherHosts, err := net.LookupHost(host)
	if err != nil {
		return khEntry, err
	}

	for _, newHost := range furtherHosts {
		if newHost != host {
			khEntry.Hosts = append(khEntry.Hosts, newHost)
		}
	}

	return khEntry, err
}

// getPublicKey gathers a public key in a specific algorithm
func getPublicKey(host, algo string, timeout time.Duration) (key ssh.PublicKey, err error) {
	d := net.Dialer{Timeout: timeout}

	conn, err := d.Dial("tcp", net.JoinHostPort(host, "22"))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	config := ssh.ClientConfig{
		HostKeyAlgorithms: []string{algo},
		HostKeyCallback:   hostKeyCallback(&key),
	}

	sshconn, _, _, err := ssh.NewClientConn(conn, host, &config)
	if err == nil {
		sshconn.Close()
	}

	return key, nil
}

// ParsePrivateKey reads the private key from given path
func ParsePrivateKey(keyPath string) (*ssh2.PublicKeys, error) {
	sshKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	pubKey, err := ssh2.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return pubKey, err
	}

	return pubKey, nil
}
