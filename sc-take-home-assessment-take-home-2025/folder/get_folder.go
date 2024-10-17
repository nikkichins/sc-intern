package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

// Custom Error Types
type InvalidOrgError struct {
	Err string
}

type NotInOrgError struct {
	Err string
}

type NoFolderError struct {
	Err string
}

func (invalidErr InvalidOrgError) Error() string {
	return invalidErr.Err
}

func (notInOrgErr NotInOrgError) Error() string {
	return notInOrgErr.Err
}

func (noFolderErr NoFolderError) Error() string {
	return noFolderErr.Err
}

// Helper data retrieval methods
func GetAllFolders() []Folder {
	return GetSampleData()
}

func GetTest() []TestAns {
	return getTestAns()
}

// Implemented methods
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
		return nil, InvalidOrgError{Err: "Error: Organization does not exist"}
		// return nil, errors.New("Error: Organization does not exist")
	}

	// Find path of parent folder
	for _, f := range orgFolders {
		if f.Name == name {
			branchPath = f.Paths
		}
	}

	// Error handling - if parent folder does not exist or 
	// is not in specified org
	if branchPath == "" {
		for _, f := range folders {
			if f.Name == name && f.OrgId != orgID {
				return nil, NotInOrgError{Err: "Error: Folder does not exist in the specified organization"}
			}
		}
		return nil, NoFolderError{"Error: Folder does not exist"}
	}

	res := []Folder{}
	for _, f := range orgFolders {
		if strings.HasPrefix(f.Paths, branchPath) && f.Name != name {
			res = append(res, f)
		}
	}

	return res, nil
}
