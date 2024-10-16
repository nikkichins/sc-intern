package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	res2, err := folderDriver.GetAllChildFolders(uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), "sacred-moonstar")
	if err != nil {
		fmt.Println("Error:", err)
	}
	folder.PrettyPrint(res2)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
}
