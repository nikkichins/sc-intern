package folder_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	testUUID := uuid.FromStringOrNil("335b32d4-5d58-4923-9f8d-f9c4f63040b9")
	testUUID_2 := uuid.FromStringOrNil("de5349c1-78aa-4052-b7a5-c3f600bafbf6")
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"Test_default_orgID", testUUID, folder.GetAllFolders(), folder.GetSampleData("get_orgid_set0.json")},
		{"Test_de5349c1_orgID", testUUID_2, folder.GetAllFolders(), folder.GetSampleData("get_orgid_set1.json")},
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
	testUUID := uuid.FromStringOrNil("335b32d4-5d58-4923-9f8d-f9c4f63040b9")
	testUUID_2 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	var tests = []struct {
		name		string
		orgID 		uuid.UUID
		folderName 	string
		want 		[]folder.Folder
	}{
		{"No_children", testUUID, "liberal-legion", []folder.Folder{}},
		{"One_child", testUUID, "still-inertia", folder.GetSampleData("get_child_set0.json")},
		{"Multiple_leaf_children", testUUID, "warm-skyrocket", folder.GetSampleData("get_child_set1.json")},
		{"Multiple_subchildren", testUUID_2, "flowing-shaman", folder.GetSampleData("get_child_set2.json")},
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