package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
  "github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
  t.Run("GetAllFolders successfully", func(t *testing.T) {
    orgId := uuid.FromStringOrNil(folders.DefaultOrgID)
    req := &folders.FetchFolderRequest{
      OrgID: orgId,
    }

    res, err := folders.GetAllFolders(req)

    if err != nil {
      t.Fatalf("Unexpected error: %v", err)
    }

    for _, folder := range res.Folders {
      if folder.OrgId != orgId {
        t.Errorf("Expected OrgId %s, but got %s", orgId, folder.OrgId)
      }
    }
  })

  t.Run("Test non-existent OrgID", func(t *testing.T) {
    req := &folders.FetchFolderRequest{
      OrgID: uuid.FromStringOrNil("abcd"),
    }

    res, err := folders.GetAllFolders(req)

    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if res == nil {
      t.Errorf("Expected empty slice, got %v", res)
    }

    if len(res.Folders) != 0 {
        t.Errorf("Expected 0 folders, but got %d", len(res.Folders))
    }
  })

  t.Run("Test with empty OrgID", func(t *testing.T) {
    req := &folders.FetchFolderRequest{
        OrgID: uuid.FromStringOrNil(""), 
    }

    res, err := folders.GetAllFolders(req)
    
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if res == nil {
      t.Errorf("Expected empty slice, got %v", res)
    }

    if len(res.Folders) != 0 {
        t.Errorf("Expected 0 folders, but got %d", len(res.Folders))
    }
  })
}

// Extra pagination testing
// func Test_GetAllPaginatedFolders(t *testing.T) {
//   t.Run("GetAllPaginatedFolders test", func(t *testing.T) {
//     req := &folders.FetchPaginatedFolderRequest{
//       OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
//       Limit: 10,
//       Token: "",
//     }
//
//     var err error
//     // Fetch the first page
//     response, _ := folders.GetAllPaginatedFolders(req)
//     if len(response.Folders) != 10 {
//       t.Errorf("Expected 10 folders, got %d", len(response.Folders))
//     }
//
//     // Get token to fetch next page
//     nextPageToken := response.Token
//     if nextPageToken == "" {
//       t.Errorf("Expected a next token, got none")
//     }
//
//     // Fetch next page using the next token
//     req.Token = nextPageToken
//     response, err = folders.GetAllPaginatedFolders(req)
//
//     if err != nil {
//       t.Fatalf("Unexpected error: %v", err)
//     }
//
//     if len(response.Folders) != 10 {
//       t.Errorf("Expected 2 folders, got %d", len(response.Folders))
//     }
//
//     // To test last page fetched
//     for {
//       response, err = folders.GetAllPaginatedFolders(req)
//       if err != nil {
//         t.Fatalf("Unexpected error: %v", err)
//       }
//
//       if response.Token == "" {
//         break
//       }
//
//       req.Token = response.Token
//     }
//
//     if len(response.Folders) == 0 {
//       t.Error("Expected some folders, but got none")
//     }
// 	})
// }
