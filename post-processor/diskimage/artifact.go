package diskimage

import (
	"fmt"
	"os"
)

const BuilderId = "johnbellone.post-processor.diskimage"

type Artifact struct {
	Path     string
	Provider string
}

func NewArtifact(provider, path string) *Artifact {
	return &Artifact{
		Path:     path,
		Provider: provider,
	}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Id() string {
	return ""
}

func (self *Artifact) Files() []string {
	return []string{self.Path}
}

func (self *Artifact) String() string {
	return fmt.Sprintf("'%s' diskimage: %s", self.Provider, self.Path)
}

func (self *Artifact) Destroy() error {
	return os.Remove(self.Path)
}
