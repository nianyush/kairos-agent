package hook

import (
	"fmt"
	"github.com/kairos-io/kairos-agent/v2/pkg/config"
	"github.com/kairos-io/kairos-sdk/utils"
	"github.com/sanity-io/litter"
)

type Lifecycle struct{}

func (s Lifecycle) Run(c config.Config) error {
	fmt.Println(litter.Sdump(c))
	if c.Install.Reboot {
		utils.Reboot()
	}

	if c.Install.Poweroff {
		utils.PowerOFF()
	}
	return nil
}
