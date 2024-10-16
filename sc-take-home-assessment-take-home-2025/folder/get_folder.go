package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...
	folders := f.folders

	// get path of data item with orgID and name
	var branchPath string;
	orgFolders := []Folder{}

	for _, f := range folders {
		if f.OrgId == orgID {
			orgFolders = append(orgFolders, f)
		}
	}

	// Error handling - invalid orgID
	if len(orgFolders) == 0 {
		// Todo: handle error
		return nil, errors.New("Organization does not exist")
	}

	// Find path of parent folder
	for _, f := range orgFolders {
		if f.Name == name {
			branchPath = f.Paths
		}
	}

	// Error handling - if folder does not exist or 
	// is not in specified org
	if branchPath == "" {
		for _, f := range folders {
			if f.Name == name && f.OrgId != orgID {
				// Todo: handle error if folder exists but not in org
				return nil, errors.New("Folder does not exist in the specified organization")
			}
		}
		return nil, errors.New("Folder does not exist")
	}

	res := []Folder{}
	for _, f := range orgFolders {
		if strings.HasPrefix(f.Paths, branchPath) && f.Name != name {
			res = append(res, f)
		}
	}

	return res, nil
}
