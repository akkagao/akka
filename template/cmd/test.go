package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "启动test环境服务",
	Long:  `启动test环境服务 读取resources目录下的conf-test.yaml 作为配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("test called")
		start("./resources/conf-test.yaml")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
