package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "启动dev环境服务",
	Long:  `启动dev环境服务 读取resources目录下的conf-dev.yaml 作为配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("dev called")
		start("./resources/conf-dev.yaml")
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
