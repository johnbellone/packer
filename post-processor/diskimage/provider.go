package diskimage

import (
	"github.com/mitchellh/packer/packer"
)

type Provider interface {
	KeepInputArtifact() bool

	Process(packer.Ui, packer.Artifact, string) error
}
