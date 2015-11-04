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

func TestRenderPrivateNetworkNoDetailsProvided(t *testing.T) {
	privateNetwork := &PrivateNetwork{}

	output, err := privateNetwork.Render()
	expectedOutput := ""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPrivateNetworkErrorNoIpNoDHCP(t *testing.T) {
	privateNetwork := &PrivateNetwork{
		DisableAutoConfig: true,
	}

	output, err := privateNetwork.Render()
	expectedOutput := "You must either provide an IP address or enable DHCP"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", err, expectedOutput)
	}
}

func TestRenderPrivateNetworkErrorNoIpDisableAutoConfig(t *testing.T) {
	privateNetwork := &PrivateNetwork{
		Dhcp:              true,
		DisableAutoConfig: true,
	}

	output, err := privateNetwork.Render()
	expectedOutput := "You must provide an IP address when disabling auto config"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", err, expectedOutput)
	}
}

func TestRenderPrivateNetworksError(t *testing.T) {
	privateNetworks := []PrivateNetwork{
		PrivateNetwork{
			Ip: "192.168.33.10",
		},
		PrivateNetwork{
			DisableAutoConfig: true,
		},
	}

	output, err := RenderPrivateNetworks(privateNetworks)
	expectedOutput := "You must either provide an IP address or enable DHCP"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", err, expectedOutput)
	}
}
