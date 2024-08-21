package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type FetchPaginatedFolderRequest struct {
	OrgID uuid.UUID
  Limit int 
  Token string
}

type FetchPaginatedFolderResponse struct {
	Folders []*Folder
  Token string
}

