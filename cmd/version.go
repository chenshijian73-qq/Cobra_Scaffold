package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var bannerBase64 = "CiAgb29vb29vb284ICAgICAgICAgICAgb29vbyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCm84ODggICAgIDg4ICAgb29vb29vbyAgIDg4OG9vb29vICBvbyBvb29vb28gICBvb29vb29vICAgCjg4OCAgICAgICAgIDg4OCAgICAgODg4IDg4OCAgICA4ODggODg4ICAgIDg4OCBvb29vbzg4OCAgCjg4OG8gICAgIG9vIDg4OCAgICAgODg4IDg4OCAgICA4ODggODg4ICAgICAgODg4ICAgIDg4OCAgCiA4ODhvb29vODggICAgODhvb284OCAgbzg4OG9vbzg4ICBvODg4byAgICAgIDg4b29vODggOG8gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAoK"

var versionTpl = `%c[%d;%d;%dm%s%c[0m
Name: cobra_scaffold
Version: %s
BuildDate: %s
Arch: %s
CommitID: %s
`

var (
	Version   string
	BuildDate string
	CommitID  string
)

var (
	showVersion bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
		fmt.Printf(versionTpl, 0x1B, 0, 0, 34, banner, 0x1B, Version, BuildDate, runtime.GOOS+"/"+runtime.GOARCH, CommitID)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	rootCmd.AddCommand(versionCmd)
}
