package webssh

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type AuthModel int8

type SSHClientOption struct {
	Addr       string
	User       string
	Password   string
	PrivateKey string
}

func NewSSHClient(option *SSHClientOption) (*ssh.Client, error) {

	if !strings.Contains(option.Addr, ":") {
		option.Addr = option.Addr + ":22"
	}

	if option.Password != "" {
		return NewSSHClientWithPassword(option)
	}

	if option.PrivateKey != "" {
		return NewSSHClientWithPrivateKey(option)
	}

	return nil, errors.New("SSHClient: no Password or PrivateKey")

}

func NewSSHClientWithPassword(option *SSHClientOption) (*ssh.Client, error) {

	auth := ssh.Password(option.Password)

	config := &ssh.ClientConfig{
		User:            option.User,
		Auth:            []ssh.AuthMethod{auth},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", option.Addr, config)

}

func NewSSHClientWithPrivateKey(option *SSHClientOption) (*ssh.Client, error) {

	signer, err := ssh.ParsePrivateKey([]byte(option.PrivateKey))

	if err != nil {
		return nil, err
	}

	auth := ssh.PublicKeys(signer)

	config := &ssh.ClientConfig{
		User:            option.User,
		Auth:            []ssh.AuthMethod{auth},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", option.Addr, config)

}
