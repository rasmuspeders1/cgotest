package cmd

import (
	"sigidagi/qrparser/qrcode"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "QR parser for Matter device version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(version)
	},
}

var codeCmd = &cobra.Command{
	Use:     "code",
	Short:   "QR code as base38 string",
	Aliases: []string{"qr", "qrcode"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		//testQRCode := "MT:Y.K9042C00KA0648G00"
		return qrcode.Parse(args[0])
	},
}

var rootCmd = &cobra.Command{
	Use:   "qrparser",
	Short: "QR parser for Matter devices",
	Long: `QR parser for Matter controller 
	> documentation & support: https://www.skyaalborg.io/
	> source & copyright information: https://skyaalborg/qrparser`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	//i := args[0]
	log.Println("Whoops. You need to provide command, see list of commands with 'help'")
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(codeCmd)
}

// Called from main
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
