package folder_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	testUUID := uuid.FromStringOrNil("335b32d4-5d58-4923-9f8d-f9c4f63040b9")
	testUUID_2 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	tests := [...]struct {
		name string
		orgID uuid.UUID
		src string
		dst string
		want []folder.Folder
	}{
		{"Move_up_one_lvl", testUUID, "still-inertia", "select-smiling-tiger", folder.GetSampleData("move_folder_set0.json")},
		{"Move_down_one_lvl", testUUID, "accepted-kitty", "still-inertia", folder.GetSampleData("move_folder_set1.json")},
		{"Moving_up_to_second_lvl", testUUID_2, "topical-maginty", "massive-deathbird", folder.GetSampleData("move_folder_set2.json")},
		{"Moving_to_leaf", testUUID_2, "pleased-tombstone", "unique-wonder-man", folder.GetSampleData("move_folder_set3.json")},
	}
	for _, tt := range tests {
		f := folder.GetAllFolders()
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(f)
			get, err := f.MoveFolder(tt.src, tt.dst)
			if err != nil {
				t.Errorf("got error %s, want %s", err, tt.want)
			}
			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("got %s, want %s", get, tt.want)
			}
		})
	}
}

func Test_folder_MoveFolder_Errors(t *testing.T) {
	testUUID := uuid.FromStringOrNil("335b32d4-5d58-4923-9f8d-f9c4f63040b9")
	testUUID_2 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	tests := [...]struct {
		name string
		orgID uuid.UUID
		src string
		dst string
		want error
	}{
		// TODO: tests
		{"Moving_to_itself_invalid", testUUID, "invalid", "invalid", folder.NoFolderError{Err: "Error: Source folder does not exist"}},
		{"Moving_to_itself", testUUID, "still-inertia", "still-inertia", folder.InvalidMoveErr{Err: "Error: Cannot move a folder to itself"}},
		{"Src_not_exist", testUUID, "invalid", "still-inertia", folder.NoFolderError{Err: "Error: Source folder does not exist"}},
		{"Dst_not_exist", testUUID, "still-inertia", "invalid", folder.NoFolderError{Err: "Error: Destination folder does not exist"}},
		{"Moving_to_child", testUUID, "select-smiling-tiger", "accepted-kitty", folder.InvalidMoveErr{Err: "Error: Cannot move a folder to a child of itself"}},
		{"Moving_diff_org", testUUID_2, "pleased-tombstone", "select-smiling-tiger", folder.InvalidMoveErr{Err: "Error: Cannot move a folder to a different organization"}},
	}
	for _, tt := range tests {
		f := folder.GetAllFolders()
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(f)
			get, err := f.MoveFolder(tt.src, tt.dst)
			if get != nil {
				t.Errorf("Should have given error %s, but got %s", tt.want, get)
			}
			if !errors.Is(err, tt.want) {
				t.Errorf("Should have given: %s, but got: %s", tt.want, err)
			}
		})
	}
}

