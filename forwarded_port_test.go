package vagrantfile

import (
	"testing"
)

func TestRenderForwardedPort(t *testing.T) {
	forwardedPort := &ForwardedPort{
		Guest: 80,
		Host:  8080,
	}

	output, err := forwardedPort.Render()
	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

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

	output, err := RenderForwardedPorts(ports)

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

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

	output, err := RenderForwardedPorts(ports)

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	expectedOutput := `config.vm.network "forwarded_port", guest: 80, host: 8080
	config.vm.network "forwarded_port", guest: 1234, host: 5678
`

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderForwardedPortsMissingGuest(t *testing.T) {

	ports := []ForwardedPort{
		ForwardedPort{
			Host: 8080,
		},
	}

	expectedOutput := "ForwardedPort.Guest must be a valid port number"
	_, err := RenderForwardedPorts(ports)

	if err.Error() != expectedOutput {
		t.Errorf(".\nExpected an error that didn't occur: %s", expectedOutput)
		return
	}
}

func TestRenderForwardedPortsMissingHost(t *testing.T) {

	ports := []ForwardedPort{
		ForwardedPort{
			Guest: 8080,
		},
	}

	expectedOutput := "ForwardedPort.Host must be a valid port number"
	_, err := RenderForwardedPorts(ports)

	if err.Error() != expectedOutput {
		t.Errorf(".\nExpected an error that didn't occur: %s", expectedOutput)
		return
	}
}
