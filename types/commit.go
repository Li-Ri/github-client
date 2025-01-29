package github


type CommitResponse struct {
	Sha string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Message string `json:"message"`
}


