package vagrantfile

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRenderBasic(t *testing.T) {
	v := NewVagrantfile()
	output, _ := v.Render()

	expectedOutput := fmt.Sprintf("VagrantFile contents here: %d", 2)

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
