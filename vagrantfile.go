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

type Renderable interface {
	Render() (string, error)
}

func RenderGroup(input []Renderable) (output string, err error) {
	for _, v := range input {
		content, err := v.Render()
		if err != nil {
			return "", err
		}
		output = output + content + "\n	"
	}

	return output, nil

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

	publicNetwork, err := v.PublicNetwork.Render()
	if err != nil {
		return "", err
	}

	syncedFolders, err := RenderSyncedFolders(v.SyncedFolders)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`Vagrant.configure(%d) do |config|
	config.vm.box = "%s"
	config.vm.box_check_update = %t
	%s
	%s
	%s
	%s
end`,

		v.Version, v.Box, v.BoxCheckUpdate, forwardedPorts, privateNetworks, publicNetwork, syncedFolders), nil
}

func NewVagrantfile() VagrantFile {
	return VagrantFile{
		Version:        2,
		Box:            "ubuntu/trusty64",
		BoxCheckUpdate: true,
	}
}
