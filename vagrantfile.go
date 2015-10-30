package vagrantfile

import (
	"fmt"
)

type VagrantFile struct {
	Version int
}

func (v *VagrantFile) Render() (s string, err error) {
	return fmt.Sprintf("VagrantFile contents here: %d", v.Version), nil
}

func NewVagrantfile() VagrantFile {
	return VagrantFile{
		Version: 2,
	}
}
