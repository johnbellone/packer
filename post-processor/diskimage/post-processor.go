package diskimage

import (
	"fmt"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
	"os"
)

var builtins = map[string]string{
	"mitchellh.virtualbox": "virtualbox",
	"mitchellh.vmware":     "vmware",
	"mitchellh.qemu":       "qemu",
}

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Compression bool   `mapstructure:"compression"`
	OutputPath  string `mapstructure:"output"`

	tpl *packer.ConfigTemplate
}

type PostProcessor struct {
	configs map[string]*Config
}

func (self *PostProcessor) Configure(raws ...interface{}) error {

}

func (self *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	name, ok := builtins[artifact.BuilderId()]
	if !ok {
		return nil, false, fmt.Errorf(
			"Unknown artifact type, can't build diskimage: %s", artifact.BuilderId())
	}

	// Sanity check since all of these are hardcoded in the application.
	provider := providerForName(name)
	if provider == nil {
		panic(fmt.Sprintf("Bad provider name: %s", name))
	}

	ui.Say(fmt.Sprintf("Creating diskimage for '%s' provider", name))

	return NewArtifact(name, outputPath), provider.KeepInputArtifact(), nil
}

func providerForName(name string) Provider {
	switch name {
	case "virtualbox":
		return new(VirtualBoxProvider)
	case "vmware":
		return new(VMwareProvider)
	default:
		return nil
	}
}
