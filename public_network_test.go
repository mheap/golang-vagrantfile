package vagrantfile

import (
	"testing"
)

func TestRenderPublicNetworkIp(t *testing.T) {
	publicNetwork := &PublicNetwork{
		Ip: "192.168.33.10",
	}

	output, err := publicNetwork.Render()
	expectedOutput := "config.vm.network \"public_network\", ip: \"192.168.33.10\""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPublicNetworkDhcp(t *testing.T) {
	publicNetwork := &PublicNetwork{
		Dhcp: true,
	}

	output, err := publicNetwork.Render()
	expectedOutput := "config.vm.network \"public_network\", dhcp: true"

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPublicNetworkDisableAutoConfig(t *testing.T) {
	publicNetwork := &PublicNetwork{
		Ip:                "192.168.33.10",
		DisableAutoConfig: true,
	}

	output, err := publicNetwork.Render()
	expectedOutput := "config.vm.network \"public_network\", ip: \"192.168.33.10\", auto_config: false"

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPublicNetworkBridge(t *testing.T) {
	publicNetwork := &PublicNetwork{
		Ip:     "192.168.33.10",
		Bridge: "en1: Wi-Fi (AirPort)",
	}

	output, err := publicNetwork.Render()
	expectedOutput := "config.vm.network \"public_network\", ip: \"192.168.33.10\", bridge: \"en1: Wi-Fi (AirPort)\""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderPublicNetworkErrorNoIpNoDHCP(t *testing.T) {
	publicNetwork := &PublicNetwork{
		DisableAutoConfig: true,
	}

	output, err := publicNetwork.Render()
	expectedOutput := "You must either provide an IP address or enable DHCP"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", err, expectedOutput)
	}
}

func TestRenderPublicNetworkErrorNoIpDisableAutoConfig(t *testing.T) {
	publicNetwork := &PublicNetwork{
		Dhcp:              true,
		DisableAutoConfig: true,
	}

	output, err := publicNetwork.Render()
	expectedOutput := "You must provide an IP address when disabling auto config"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", err, expectedOutput)
	}
}
