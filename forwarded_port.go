package vagrantfile

import (
	"errors"
	"fmt"
)

type ForwardedPort struct {
	Guest   int
	Host    int
	BoxName string
}

func (p ForwardedPort) Render() (output string, err error) {

	// Sensible defaults
	if p.BoxName == "" {
		p.BoxName = "vm"
	}

	if p.Guest == 0 {
		return "", errors.New("ForwardedPort.Guest must be a valid port number")
	}

	if p.Host == 0 {
		return "", errors.New("ForwardedPort.Host must be a valid port number")
	}

	return fmt.Sprintf("config.%s.network \"forwarded_port\", guest: %d, host: %d", p.BoxName, p.Guest, p.Host), nil
}

func RenderForwardedPorts(ports []ForwardedPort) (output string, err error) {
	for _, v := range ports {
		content, err := v.Render()
		if err != nil {
			return "", err
		}
		output = output + content + "\n	"
	}

	return output, nil
}
