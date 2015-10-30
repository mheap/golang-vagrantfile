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
