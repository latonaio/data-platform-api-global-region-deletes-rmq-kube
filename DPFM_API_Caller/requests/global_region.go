package requests

type GlobalRegion struct {
	GlobalRegion        string `json:"GlobalRegion"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
