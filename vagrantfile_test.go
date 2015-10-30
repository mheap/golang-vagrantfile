package vagrantfile

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRenderDefaults(t *testing.T) {
	v := NewVagrantfile()
	output, _ := v.Render()

	expectedOutput := `Vagrant.configure(2) do |config|
	config.vm.box = "ubuntu/trusty64"
	config.vm.box_check_update = true
	
	
end`

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderCustom(t *testing.T) {
	v := &VagrantFile{
		Box:            "fedora/fedora",
		BoxCheckUpdate: false,
		ForwardedPorts: []ForwardedPort{
			ForwardedPort{
				Guest: 80,
				Host:  8080,
			},

			ForwardedPort{
				Guest: 1234,
				Host:  5678,
			},
		},
		PrivateNetworks: []PrivateNetwork{
			PrivateNetwork{
				Ip: "192.168.33.10",
			},
			PrivateNetwork{
				Dhcp: true,
			},
		},
	}

	output, _ := v.Render()

	expectedOutput := `Vagrant.configure(2) do |config|
	config.vm.box = "fedora/fedora"
	config.vm.box_check_update = false
	config.vm.network "forwarded_port", guest: 80, host: 8080
	config.vm.network "forwarded_port", guest: 1234, host: 5678
	
	config.vm.network "private_network", ip: "192.168.33.10"
	config.vm.network "private_network", dhcp: true
	
end`

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
