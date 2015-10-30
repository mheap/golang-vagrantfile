package vagrantfile

import (
	"fmt"
)

type ForwardedPort struct {
	Guest   int
	Host    int
	BoxName string
}

func (p *ForwardedPort) Render() string {

	// Sensible defaults
	if p.BoxName == "" {
		p.BoxName = "vm"
	}

	return fmt.Sprintf("config.%s.network \"forwarded_port\", guest: %d, host: %d", p.BoxName, p.Guest, p.Host)
}

func RenderForwardedPorts(ports []ForwardedPort) (output string) {

	for _, v := range ports {
		output = output + v.Render() + "\n	"
	}

	return output
}
