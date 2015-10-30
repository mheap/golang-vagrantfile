package vagrantfile

import (
	"errors"
	"fmt"
)

type SyncedFolder struct {
	Local   string
	Remote  string
	Type    string
	BoxName string
}

func (p SyncedFolder) Render() (output string, err error) {

	// We don't have any synced folders if all settings are nil
	if p.Local == "" && p.Remote == "" && p.Type == "" {
		return "", nil
	}

	if p.Local == "" {
		return "", errors.New("You must provide a local folder to sync")
	}

	if p.Remote == "" {
		return "", errors.New("You must provide a remote folder to sync")
	}

	if p.BoxName == "" {
		p.BoxName = "vm"
	}

	output = fmt.Sprintf("config.%s.synced_folder \"%s\", \"%s\"", p.BoxName, p.Local, p.Remote)

	return output, nil
}

func RenderSyncedFolders(folders []SyncedFolder) (output string, err error) {
	for _, v := range folders {
		content, err := v.Render()
		if err != nil {
			return "", err
		}
		output = output + content + "\n	"
	}

	return output, nil
}
