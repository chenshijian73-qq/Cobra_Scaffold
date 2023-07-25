package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	a int
	b *string
	c *int
	d *[]string

	e *int
)

var groups = []*cobra.Group{
	{ID: "1", Title: "Basic Commands:"},
	{ID: "2", Title: "Other Commands:"},
}

var rootCmd = &cobra.Command{
	Use:   "cobra [Name]",
	Short: "cobra demo command",
	// 最低输入参数个数为 1
	//Args: cobra.MinimumNArgs(1),
	// 自定义参数校验
	/*Args: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			_, err := strconv.Atoi(arg)
			if err == nil {
				return fmt.Errorf("仅支持字符串, 不支持使用数字[%s]作为参数", arg)
			}
		}
		return nil
	},*/
	// validArgs的配置
	/*ValidArgs: []string{"a", "b"},*/
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a:", a, "\nb:", *b, "\nc:", *c, "\nd:", *d, "\ne:", *e)
	},
}

func init() {
	rootCmd.AddGroup(groups...)
	// 非持久化参数
	rootCmd.Flags().IntVarP(&a, "a", "a", 0, "specify int")
	b = rootCmd.Flags().String("b", "", "specify string")
	c = rootCmd.Flags().Int("c", 1, "specify int")
	d = rootCmd.Flags().StringArray("d", []string{}, "specify string array")
	// 持久化参数, 子命令继承使用
	e = rootCmd.PersistentFlags().Int("e", 2, "persistent int")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
