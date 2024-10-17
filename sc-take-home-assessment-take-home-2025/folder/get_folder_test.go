package folder_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	ans := folder.GetTest()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		// default orgid
		{"Test_default_orgID", ans[0].OrgID, folder.GetAllFolders(), ans[0].FolderList},
		{"Test_de5349c1_orgID", ans[5].OrgID, folder.GetAllFolders(), ans[5].FolderList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("got %s, want %s", get, tt.want)
			}
		})
	}
}

func Test_folder_GetChildFolders(t *testing.T) {
	ans := folder.GetTest()
	var tests = []struct {
		name		string
		orgID 		uuid.UUID
		folderName 	string
		want 		[]folder.Folder
	}{
		// TODO
		{"No_children", ans[3].OrgID, "liberal-legion", ans[3].FolderList},
		{"One_child", ans[4].OrgID, "still-inertia", ans[4].FolderList},
		{"Multiple_leaf_children", ans[1].OrgID, "warm-skyrocket", ans[1].FolderList},
		{"Multiple_subchildren", ans[2].OrgID, "flowing-shaman", ans[2].FolderList},
	}
	for _, tt := range tests {
		f := folder.GetAllFolders()
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(f)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folderName)
			if err != nil {
				t.Errorf("got error %s, want %s", err, tt.want)
			}
			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("got %s, want %s", get, tt.want)
			}
			fmt.Println(get)
		})
	}
}

func Test_folder_GetChildFolders_Errors(t *testing.T) {
	testUUID := uuid.FromStringOrNil("335b32d4-5d58-4923-9f8d-f9c4f63040b9")
	var tests = []struct {
		name		string
		orgID 		uuid.UUID
		folderName 	string
		want 		error
	}{
		// TODO
		{"Invalid_Org", uuid.FromStringOrNil("non-existent"), "liberal-legion", folder.InvalidOrgError{Err: "Error: Organization does not exist"}},
		{"Folder_does_not_exist", testUUID, "non-existent", folder.NoFolderError{Err: "Error: Folder does not exist"}},
		{"Folder_not_in_org", testUUID, "topical-maginty", folder.NotInOrgError{Err: "Error: Folder does not exist in the specified organization"}},
	}
	for _, tt := range tests {
		f := folder.GetAllFolders()
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(f)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folderName)
			if get != nil {
				t.Errorf("Should have given error %s, but got %s", tt.want, get)
			}
			if !errors.Is(err, tt.want) {
				t.Errorf("Should have given: %s, but got: %s", tt.want, err)
			}
			
		})
	}
}