package vagrantfile

import (
	"errors"
	"fmt"
)

type PublicNetwork struct {
	Dhcp              bool
	Ip                string
	Bridge            string
	BoxName           string
	DisableAutoConfig bool
}

func (p *PublicNetwork) Render() (output string, err error) {

	// We don't have any public networks if all settings are nil
	if !p.Dhcp && p.Ip == "" && p.Bridge == "" && !p.DisableAutoConfig {
		return "", nil
	}

	if !p.Dhcp && p.Ip == "" {
		return "", errors.New("You must either provide an IP address or enable DHCP")
	}

	if p.DisableAutoConfig && p.Ip == "" {
		return "", errors.New("You must provide an IP address when disabling auto config")
	}

	if p.BoxName == "" {
		p.BoxName = "vm"
	}

	output = fmt.Sprintf("config.%s.network \"public_network\"", p.BoxName)

	if p.Dhcp {
		output = output + ", dhcp: true"
	}

	if p.Ip != "" {
		output = output + fmt.Sprintf(", ip: \"%s\"", p.Ip)
	}

	if p.Bridge != "" {
		output = output + fmt.Sprintf(", bridge: \"%s\"", p.Bridge)
	}

	if p.DisableAutoConfig {
		output = output + ", auto_config: false"
	}

	return output, nil
}
