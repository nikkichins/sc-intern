package folder_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	tests := [...]struct {
		name string
		orgID uuid.UUID
		src string
		dst string
		want []folder.Folder
	}{
		// TODO: tests
		// moving 1 level up
		// moving 1 level down
		// moving the second node
		// moving to second node
		// moving to leaf node
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
	// TODO: your tests here
	tests := [...]struct {
		name string
		orgID uuid.UUID
		src string
		dst string
		want error
	}{
		// TODO: tests
		// test for cannot move to itself if theyre both invalid but w same name
		// movign it itself
		// src does not exist
		// dest does not exist
		// same org moving to child (child error)
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

