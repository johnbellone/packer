package diskimage

import (
	"archive/tar"
	"errors"
	"fmt"
	"github.com/mitchellh/packer/packer"
	"os"
)

type VirtualBoxProvider struct{}

func (*VirtualBoxProvider) KeepInputArtifact() bool {
	return false
}

func (self *VirtualBoxProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) error {
}
