package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var prodCmd = &cobra.Command{
	Use:   "prod",
	Short: "启动prod环境服务",
	Long:  `启动prod环境服务 读取resources目录下的conf-prod.yaml 作为配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("prod called")
		start("./resources/conf-prod.yaml")
	},
}

func init() {
	rootCmd.AddCommand(prodCmd)
}
