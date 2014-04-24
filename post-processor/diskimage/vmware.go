package diskimage

import (
	"archive/tar"
	"errors"
	"fmt"
	"github.com/mitchellh/packer/packer"
	"os"
)

type VMwareProvider struct{}

func (*VMwareProvider) KeepInputArtifact() bool {
	return false
}

func (self *VMwareProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) error {
}
