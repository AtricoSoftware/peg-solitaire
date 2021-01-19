// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package settings

import (
"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const verboseSettingName = "verbose"
const verboseSettingShortcut = "v"

// Fetch the setting
func (theSettings) Verbose() bool {
	return viper.GetBool(verboseSettingName)
}

func AddVerboseFlag(flagSet *pflag.FlagSet) {
	viperEx.AddBoolSettingP(flagSet, verboseSettingName, verboseSettingShortcut, "Output all stages of operation")
}