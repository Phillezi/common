package cli

import (
	"fmt"
	"os"

	viperconf "github.com/Phillezi/common/config/viper"
	zetup "github.com/Phillezi/common/logging/zap"
	"github.com/Phillezi/common/scripts/glint/internal/api"
	"github.com/Phillezi/common/scripts/glint/internal/cyclo"
	"github.com/Phillezi/common/scripts/glint/internal/defaults"
	"github.com/Phillezi/common/scripts/glint/internal/ineffassign"
	"github.com/Phillezi/common/scripts/glint/pkg/mod"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:  "glint",
	Long: glint,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		zetup.Setup()
	},
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		banner()
		root, err := mod.FindGoModRoot(args...)
		if err != nil {
			zap.L().Fatal("could not find go.mod", zap.Error(err))
		}

		var errors []error
		for _, cmd := range []api.Command{
			(&cyclo.Command{Threshold: viper.GetInt("gocyclo.over")}),
			(&ineffassign.Command{}),
		} {
			if err := cmd.Run(root); err != nil {
				errors = append(errors, err)
			}
		}
		if len(errors) > 0 {
			zap.L().Error("")
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	cobra.OnInitialize(func() { viperconf.InitConfig("lint", "glint") })

	rootCmd.PersistentFlags().IntP("gocyclo-over", "t", defaults.DefaultGocycloOverValue, "Set the cyclomatic complexity level threshold")
	viper.BindPFlag("gocyclo.over", rootCmd.PersistentFlags().Lookup("gocyclo-over"))

	rootCmd.PersistentFlags().String("loglevel", "info", "Set the logging level (info, warn, error, debug)")
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))

	rootCmd.PersistentFlags().String("profile", "", "Set the logging profile (production or empty)")
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))

	rootCmd.PersistentFlags().Bool("stacktrace", false, "Show the stack trace in error logs")
	viper.BindPFlag("stacktrace", rootCmd.PersistentFlags().Lookup("stacktrace"))

	rootCmd.AddCommand(versionCmd)
}

func ExecuteE() error {
	return rootCmd.Execute()
}

func GetRootCMD() *cobra.Command {
	return rootCmd
}
