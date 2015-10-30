package vagrantfile

import (
	"fmt"
)

type VagrantFile struct {
	Version         int
	Box             string
	BoxCheckUpdate  bool
	ForwardedPorts  []ForwardedPort
	PrivateNetworks []PrivateNetwork
	PublicNetwork   PublicNetwork
	SyncedFolders   []SyncedFolder
}

type PublicNetwork struct {
	Dhcp       bool
	Ip         string
	Bridge     string
	AutoConfig bool
}

type SyncedFolder struct {
	Local  string
	Remote string
	Type   string
}

func (v *VagrantFile) Render() (s string, err error) {

	// Set some smart defaults
	if v.Version < 1 {
		v.Version = 2
	}

	forwardedPorts, err := RenderForwardedPorts(v.ForwardedPorts)

	if err != nil {
		return "", err
	}

	privateNetworks, err := RenderPrivateNetworks(v.PrivateNetworks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`Vagrant.configure(%d) do |config|
	config.vm.box = "%s"
	config.vm.box_check_update = %t
	%s
	%s
end`,

		v.Version, v.Box, v.BoxCheckUpdate, forwardedPorts, privateNetworks), nil
}

func NewVagrantfile() VagrantFile {
	return VagrantFile{
		Version:        2,
		Box:            "ubuntu/trusty64",
		BoxCheckUpdate: true,
	}
}
