package folder

import (
	"strings"
)

// Custom errors
type InvalidMoveErr struct {
	Err string
}

func (invalidMoveErr InvalidMoveErr) Error() string {
	return invalidMoveErr.Err
}

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders
	var srcFolder *Folder
	var dstFolder *Folder

	for i := range folders {
		if folders[i].Name == name {
			srcFolder = &folders[i]
		}

		if folders[i].Name == dst {
			dstFolder = &folders[i]
		}
	}

	// Error handling
	if srcFolder == nil {
		return nil, NoFolderError{"Error: Source folder does not exist"}
	}

	if dstFolder == nil {
		return nil, NoFolderError{"Error: Destination folder does not exist"}
	}

	if name == dst {
		return nil, InvalidMoveErr{"Error: Cannot move a folder to itself"}
	}

	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, InvalidMoveErr{"Error: Cannot move a folder to a child of itself"}
	}

	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, InvalidMoveErr{"Error: Cannot move a folder to a different organization"}
	}

	oldPath := srcFolder.Paths
	newPath := dstFolder.Paths + "." + srcFolder.Name
	srcFolder.Paths = newPath

	// Changing children paths
	for i := range folders {
		if strings.HasPrefix(folders[i].Paths, oldPath) {
			newChildPath := strings.Replace(folders[i].Paths, oldPath, newPath, 1)
			folders[i].Paths = newChildPath
		}
	}

	return folders, nil
}
