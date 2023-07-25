package cmd

import (
	"github.com/chenshijian73-qq/cobra_scaffold/internal/logic"
	"github.com/spf13/cobra"
)

var s string

var subcmd = &cobra.Command{
	Use:     "subcmd",
	Short:   "sub cmd Usage",
	GroupID: "1",
	// 最低输入参数个数为 1
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logic.Hello(s)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&s, "s", "s", "", "specify string")
	rootCmd.AddCommand(subcmd)
}
