package webssh

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type AuthModel int8

type SSHClientOption struct {
	RemoteAddr string
	User       string
	Password   string
	PulicKey   string
}

func NewSSHClient(option *SSHClientOption) (*ssh.Client, error) {

	if !strings.Contains(option.RemoteAddr, ":") {
		option.RemoteAddr = option.RemoteAddr + ":22"
	}

	if option.Password != "" {
		return NewSSHClientWithPassword(option)
	}

	if option.PulicKey != "" {
		return NewSSHClientWithPulicKey(option)
	}

	return nil, errors.New("no Password or PublicKey")

}

func NewSSHClientWithPassword(option *SSHClientOption) (*ssh.Client, error) {

	auth := ssh.Password(option.Password)

	config := &ssh.ClientConfig{
		User:            option.User,
		Auth:            []ssh.AuthMethod{auth},
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", option.RemoteAddr, config)

}

func NewSSHClientWithPulicKey(option *SSHClientOption) (*ssh.Client, error) {

	signer, err := ssh.ParsePrivateKey([]byte(option.PulicKey))
	if err != nil {
		return nil, err
	}

	auth := ssh.PublicKeys(signer)

	config := &ssh.ClientConfig{
		User:            option.User,
		Auth:            []ssh.AuthMethod{auth},
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", option.RemoteAddr, config)

}
