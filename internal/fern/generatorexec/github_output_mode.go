// Generated by Fern. Do not edit.

package generatorexec

type GithubOutputMode struct {
	Version string `json:"version"`
	// A full repo url (i.e. https://github.com/fern-api/fern)
	RepoUrl     string             `json:"repoUrl"`
	PublishInfo *GithubPublishInfo `json:"publishInfo"`
}
