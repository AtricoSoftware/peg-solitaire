// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package cmd

import (
	"fmt"
	"os"

	"github.com/atrico-go/container"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/peg-solitaire/pkg"
	"github.com/AtricoSoftware/peg-solitaire/settings"
)

func CreateCommands(c container.Container) *cobra.Command {
	cobra.OnInitialize(initConfig)
	rootCmd := CreateRootCommand()
	rootCmd.AddCommand(CreateVersionCommand())
	// Add commands
	rootCmd.AddCommand(CreateSolveCommand(c))
	return rootCmd
}

func CreateRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   pkg.Name,
		Short: pkg.Summary,
		Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "alternate config file")
	settings.AddVerboseFlag(cmd.PersistentFlags())
	return cmd
}

var cfgFile string

func initConfig() {
	// Config file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := tryReadConfig(); err != nil {
			// Fail if specified config cannot be read
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Standard name for config
		viper.SetConfigName(fmt.Sprintf(".%s", pkg.Name))
		// Try current working directory
		dir, err := os.Getwd()
		if err == nil {
			viper.AddConfigPath(dir)
			err = tryReadConfig()
		}
		if err != nil {
			// Finally, try home directory
			dir, err = homedir.Dir()
			if err == nil {
				viper.AddConfigPath(dir)
				tryReadConfig()
			}
		}
	}
}

func tryReadConfig() error {
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	return err
}
