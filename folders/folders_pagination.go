package folders
import (
	"github.com/gofrs/uuid"
  "encoding/base64"
  "strconv"
)

//In this pagination implementation I created two new types, namely, FetchPaginatedFolderRequest, which has a Limit
//and Token variables on top of the OrgID field from the original FetchFolderRequest and 
//FetchPaginatedFolderResponse containing the Token field as an addition to the Folders field from the original
//FetchFolderResponse. 
//FetchPaginatedFolderRequest intends to allow the user to specify, how many folders they want returned (through Limit), and
//the token of next 'page' of folders.
//We create these tokens by encoding the index of the next folder of the currently fetched list of folders

// The token given through the request can then be parsed back as the end index, to which we can 
// fetch the next specified amount of folders desired. When there are less folders than the limit to be retrieved, 
// we only fetch the remaining folders. At this point the Token field is "" in the FetchPaginatedFolderResponse, since 
// there are no folders left to fetch.
func GetAllPaginatedFolders(req *FetchPaginatedFolderRequest) (*FetchPaginatedFolderResponse, error) {
	folders, nextToken, err := FetchAllPaginatedFoldersByOrgID(req.OrgID, req.Limit, req.Token)

  if err != nil {
    return nil, err
  }

  ffr := &FetchPaginatedFolderResponse{
    Folders: folders,
    Token: nextToken,
  }
	return ffr, nil
}

func FetchAllPaginatedFoldersByOrgID(orgID uuid.UUID, limit int, token string) ([]*Folder, string, error) {
  startIndex, err := ParseToken(token)
  if err != nil {
    return nil, "", err
  }

	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

  endIndex := startIndex + limit
  if endIndex > len(folders) {
    endIndex = len(folders)
  }

  nextToken := ""
  if endIndex < len(folders) {
    nextToken = GenerateToken(endIndex)
  }

  return resFolder[startIndex:endIndex], nextToken, nil
}

func GenerateToken(index int) string {
  return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(index)))
} 

func ParseToken(token string) (int, error) {
  // Start from the beginning if no token is provided
  if token == "" {
    return 0, nil 
  }

  decoded, err := base64.StdEncoding.DecodeString(token)
  if err != nil {
    return 0, err 
  }
  return strconv.Atoi(string(decoded))
}
