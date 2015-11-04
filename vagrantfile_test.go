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
	output, err := v.Render()

	expectedOutput := `Vagrant.configure(2) do |config|
	config.vm.hostname = "default"
	config.vm.box = "ubuntu/trusty64"
	config.vm.box_check_update = true
	
	
	
	
end`

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderCustom(t *testing.T) {
	v := &VagrantFile{
		Hostname:       "michael",
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

		PublicNetwork: PublicNetwork{
			Dhcp:   true,
			Bridge: "en1 (Airport)",
		},

		SyncedFolders: []SyncedFolder{
			SyncedFolder{
				Local:  "../data",
				Remote: "/var/www/folder",
			},
			SyncedFolder{
				Local:  "/path/to/folder",
				Remote: "/tmp/bees",
			},
		},
	}

	output, err := v.Render()

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	expectedOutput := `Vagrant.configure(2) do |config|
	config.vm.hostname = "michael"
	config.vm.box = "fedora/fedora"
	config.vm.box_check_update = false
	config.vm.network "forwarded_port", guest: 80, host: 8080
	config.vm.network "forwarded_port", guest: 1234, host: 5678
	
	config.vm.network "private_network", ip: "192.168.33.10"
	config.vm.network "private_network", dhcp: true
	
	config.vm.network "public_network", dhcp: true, bridge: "en1 (Airport)"
	config.vm.synced_folder "../data", "/var/www/folder"
	config.vm.synced_folder "/path/to/folder", "/tmp/bees"
	
end`

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
