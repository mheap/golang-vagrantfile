package vagrantfile

import (
	"testing"
)

func TestRenderPrivateNetworkIp(t *testing.T) {
	privateNetwork := &PrivateNetwork{
		Ip: "192.168.33.10",
	}

	output, err := privateNetwork.Render()
	expectedOutput := "config.vm.network \"private_network\", ip: \"192.168.33.10\""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPrivateNetworkDhcp(t *testing.T) {
	privateNetwork := &PrivateNetwork{
		Dhcp: true,
	}

	output, err := privateNetwork.Render()
	expectedOutput := "config.vm.network \"private_network\", dhcp: true"

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPrivateNetworkDisableAutoConfig(t *testing.T) {
	privateNetwork := &PrivateNetwork{
		Ip:                "192.168.33.10",
		DisableAutoConfig: true,
	}

	output, err := privateNetwork.Render()
	expectedOutput := "config.vm.network \"private_network\", ip: \"192.168.33.10\", auto_config: false"

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPrivateNetworks(t *testing.T) {
	privateNetworks := []PrivateNetwork{
		PrivateNetwork{
			Ip: "192.168.33.10",
		},
		PrivateNetwork{
			Dhcp: true,
		},
	}

	output, err := RenderPrivateNetworks(privateNetworks)
	expectedOutput := `config.vm.network "private_network", ip: "192.168.33.10"
	config.vm.network "private_network", dhcp: true
	`

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
