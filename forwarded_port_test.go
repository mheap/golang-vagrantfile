package vagrantfile

import (
	"testing"
)

func TestRender(t *testing.T) {
	forwardedPort := &ForwardedPort{
		Guest: 80,
		Host:  8080,
	}

	output := forwardedPort.Render()
	expectedOutput := "config.vm.network \"forwarded_port\", guest: 80, host: 8080"

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderForwardedPortsSinglePort(t *testing.T) {

	ports := []ForwardedPort{
		ForwardedPort{
			Guest: 80,
			Host:  8080,
		},
	}

	output := RenderForwardedPorts(ports)
	expectedOutput := "config.vm.network \"forwarded_port\", guest: 80, host: 8080\n	"

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func FooTestRenderForwardedPortsMultiplePorts(t *testing.T) {

	ports := []ForwardedPort{
		ForwardedPort{
			Guest: 80,
			Host:  8080,
		},

		ForwardedPort{
			Guest: 1234,
			Host:  5678,
		},
	}

	output := RenderForwardedPorts(ports)
	expectedOutput := `config.vm.network "forwarded_port", guest: 80, host: 8080
	config.vm.network "forwarded_port", guest: 1234, host: 5678
`

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
