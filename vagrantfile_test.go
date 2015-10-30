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
	}

	output, _ := v.Render()

	expectedOutput := `Vagrant.configure(2) do |config|
	config.vm.box = "fedora/fedora"
	config.vm.box_check_update = false
	
end`

	if output != expectedOutput {
		t.Errorf("Default VagrantFile was incorrect.\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
