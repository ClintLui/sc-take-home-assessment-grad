package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
  // Unused declared variables
	//var (
	//	err error
	//	f1  Folder
	//	fs  []*Folder
	//)

  //This variable becomes redundant once we comment out the below
	//f := []Folder{}
	// r, _ := FetchAllFoldersByOrgID(req.OrgID)
  // Use a more meaningful variable name like folders instead of 'r'
	folders, err := FetchAllFoldersByOrgID(req.OrgID)

  //Handle error from FetchAllFoldersByOrgID
  // While FetchAllFoldersByOrgID only ever returns nil
  // this could change and is hence useful for future error handling
  if err != nil {
    return nil, err
  }
  // The code below assigns all folders of the specified organisation
  // into f, before storing the folder pointers into fp 

  //This is redundant as FetchAllFoldersByOrgID already returns
  //folder pointers of the folders from the desired organisation

  // Unused index variables k and k1
	//for k, v := range r {
	//	f = append(f, *v)
	//}
	//var fp []*Folder
	//for k1, v1 := range f {
	//	fp = append(fp, &v1)
	//}

	var ffr *FetchFolderResponse
  // We no longer need the variable fp, and can
  // source the folder pointers directly from variable r
	// ffr = &FetchFolderResponse{Folders: fp}
	ffr = &FetchFolderResponse{Folders: folders}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
