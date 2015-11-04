package vagrantfile

import (
	"testing"
)

func TestRenderSyncedFolder(t *testing.T) {
	folder := &SyncedFolder{
		Local:  "../data",
		Remote: "/var/www/folder",
	}

	output, err := folder.Render()
	expectedOutput := "config.vm.synced_folder \"../data\", \"/var/www/folder\""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderSyncedFolders(t *testing.T) {
	folders := []SyncedFolder{
		SyncedFolder{
			Local:  "../data",
			Remote: "/var/www/folder",
		},
		SyncedFolder{
			Local:  "/path/to/folder",
			Remote: "/tmp/bees",
		},
	}

	output, err := RenderSyncedFolders(folders)
	expectedOutput := `config.vm.synced_folder "../data", "/var/www/folder"
	config.vm.synced_folder "/path/to/folder", "/tmp/bees"
	`

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderSyncedFolderNoSettings(t *testing.T) {
	folder := &SyncedFolder{}

	output, err := folder.Render()
	expectedOutput := ""

	if err != nil {
		t.Errorf(".\nGot an unexpected error: %s", err)
		return
	}

	if output != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderSyncedFolderNoLocalFolder(t *testing.T) {
	folder := &SyncedFolder{
		Remote: "/tmp",
	}

	output, err := folder.Render()
	expectedOutput := "You must provide a local folder to sync"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
		return
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderSyncedFolderNoRemoteFolder(t *testing.T) {
	folder := &SyncedFolder{
		Local: "/tmp",
	}

	output, err := folder.Render()
	expectedOutput := "You must provide a remote folder to sync"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
		return
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}

func TestRenderSyncedFolderListErr(t *testing.T) {
	folders := []SyncedFolder{
		SyncedFolder{
			Local: "/tmp",
		},
	}

	output, err := RenderSyncedFolders(folders)
	expectedOutput := "You must provide a remote folder to sync"

	if output != "" {
		t.Errorf("Got an unexpected result (Was expecting error): %s", output)
		return
	}

	if err.Error() != expectedOutput {
		t.Errorf(".\nGot: %s\nExpected: %s", output, expectedOutput)
	}
}
