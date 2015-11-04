package vagrantfile

import (
	"errors"
	"fmt"
)

type PrivateNetwork struct {
	Dhcp              bool
	Ip                string
	DisableAutoConfig bool
	BoxName           string
}

func (p PrivateNetwork) Render() (output string, err error) {

	// We don't have any private networks if all settings are nil
	if !p.Dhcp && p.Ip == "" && !p.DisableAutoConfig {
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

	output = fmt.Sprintf("config.%s.network \"private_network\"", p.BoxName)

	if p.Dhcp {
		output = output + ", dhcp: true"
	}

	if p.Ip != "" {
		output = output + fmt.Sprintf(", ip: \"%s\"", p.Ip)
	}

	if p.DisableAutoConfig {
		output = output + ", auto_config: false"
	}

	return output, nil
}

func RenderPrivateNetworks(networks []PrivateNetwork) (output string, err error) {
	for _, v := range networks {
		content, err := v.Render()
		if err != nil {
			return "", err
		}
		output = output + content + "\n	"
	}

	return output, nil
}
