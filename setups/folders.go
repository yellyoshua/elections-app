package setups

import (
	"os"

	"github.com/yellyoshua/elections-app/server/modules/api"
)

type folder struct {
	path        string
	permissions os.FileMode
}

// Folders create and setup permissions if don't exist
func Folders() {

	folders := []folder{
		{path: api.PublicFolder, permissions: 0755},
		{path: api.UploadFolder, permissions: 0755},
	}

	for _, f := range folders {
		if notExistFolder(f.path) {
			go os.Mkdir(f.path, f.permissions)
		}
	}
	return
}

func notExistFolder(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
