package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	
	// Error handling - cannot move into itself
	if name == dst {
		return nil, errors.New("Cannot move a folder to itself")
	}
	
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
		return nil, errors.New("Source folder does not exist")
	}

	if dstFolder == nil {
		return nil, errors.New("Destination folder does not exist")
	}

	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	oldPath := srcFolder.Paths
	newPath := dstFolder.Paths + "." + srcFolder.Name
	srcFolder.Paths = newPath

	for i := range folders {
		if strings.HasPrefix(folders[i].Paths, oldPath) {
			newChildPath := strings.Replace(folders[i].Paths, oldPath, newPath, 1)
			folders[i].Paths = newChildPath
		}
	}

	return folders, nil
}
