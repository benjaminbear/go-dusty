package sshhandler

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type KnownHosts struct {
	knownHosts     []KnownHost
	knownHostsFile string
}

type KnownHost struct {
	Hosts []string
	Key   ssh.PublicKey
}

// AddHostIfMissing adds a host to knownhosts file if needed
func AddHostIfMissing(host, knownHostFile string) error {
	kh, err := ParseKnownHosts(knownHostFile)
	if err != nil {
		return err
	}

	if !kh.ContainsHost(host) {
		hostKey, err := GetKey(host, time.Second*5)
		if err != nil {
			return err
		}

		kh.AddHost(hostKey)

		err = kh.SaveKnownHosts()
		if err != nil {
			return err
		}
	}

	return nil
}

// ParseKnownHosts reads the given known_hosts file and parses the entries
// If the file doesn't exist, it will be created
func ParseKnownHosts(knownHostsFile string) (*KnownHosts, error) {
	if _, err := os.Stat(knownHostsFile); os.IsNotExist(err) {
		f, err := os.Create(knownHostsFile)
		if err != nil {
			return nil, err
		}

		defer f.Close()
	}

	knownHostsBytes, err := ioutil.ReadFile(knownHostsFile)
	if err != nil {
		return nil, err
	}

	knownHosts := &KnownHosts{knownHostsFile: knownHostsFile}

	for len(knownHostsBytes) > 0 {
		kh := KnownHost{}

		_, kh.Hosts, kh.Key, _, knownHostsBytes, err = ssh.ParseKnownHosts(knownHostsBytes)
		if err != nil {
			return nil, err
		}
		knownHosts.knownHosts = append(knownHosts.knownHosts, kh)
	}

	return knownHosts, nil
}

// ContainsHost searches all known hosts entries for a host
// if the host is not found the hashed host is searched for
func (kh *KnownHosts) ContainsHost(host string) bool {
	for _, kH := range kh.knownHosts {
		for _, h := range kH.Hosts {
			if h == host {
				return true
			}
		}
	}

	hashedHost := knownhosts.HashHostname(host)

	for _, kH := range kh.knownHosts {
		for _, h := range kH.Hosts {
			if h == hashedHost {
				return true
			}
		}
	}

	return false
}

// Adds an known host entry to the cached entries
func (kh *KnownHosts) AddHost(host KnownHost) {
	kh.knownHosts = append(kh.knownHosts, host)
}

// Overwrites existing knownhosts file with all cached entries
// including the newly added once
func (kh *KnownHosts) SaveKnownHosts() error {
	f, err := os.OpenFile(kh.knownHostsFile, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, entry := range kh.knownHosts {
		khLine := knownhosts.Line(entry.Hosts, entry.Key)
		khLine = strings.ReplaceAll(khLine, "[", "")
		khLine = strings.ReplaceAll(khLine, "]", "")

		_, err = f.WriteString(fmt.Sprintln(khLine))
		if err != nil {
			return err
		}
	}

	return nil
}

// hostKeyCallback is the callback function of the server answer to get save the public key
func hostKeyCallback(publicKey *ssh.PublicKey) func(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		*publicKey = key
		return nil
	}
}
