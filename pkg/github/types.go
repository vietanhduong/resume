package github

type github struct {
	user   string
	repo   string
	branch string
}

type Permission struct {
	Admin bool `json:"admin"`
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
}

type RepoResponse struct {
	Name        string      `json:"name"`
	Permissions *Permission `json:"permissions,omitempty"`
}
