package diskimage

import (
	"archive/tar"
	"errors"
	"fmt"
	"github.com/mitchellh/packer/packer"
	"os"
)

type QemuProvider struct{}

func (*QemuProvider) KeepInputArtifact() bool {
	return false
}

func (self *QemuProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) error {
}
