package hook

import (
	"fmt"
	config "github.com/kairos-io/kairos-agent/v2/pkg/config"
	"github.com/sanity-io/litter"
)

type Interface interface {
	Run(c config.Config) error
}

var AfterInstall = []Interface{
	&RunStage{},    // Shells out to stages defined from the container image
	&GrubOptions{}, // Set custom GRUB options
	&BundleOption{},
	&CustomMounts{},
	&Kcrypt{},
	&Lifecycle{}, // Handles poweroff/reboot by config options
}

var AfterReset = []Interface{}

var FirstBoot = []Interface{
	&BundlePostInstall{},
	&GrubPostInstallOptions{},
}

func Run(c config.Config, hooks ...Interface) error {
	for _, h := range hooks {
		fmt.Println(litter.Sdump(h))
		if err := h.Run(c); err != nil {
			return err
		}
	}
	return nil
}
