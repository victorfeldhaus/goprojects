package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var task, cluster, env string

	var rootCmd = &cobra.Command{Use: "aws-task-runner"}

	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Run task on AWS",
		Run: func(cmd *cobra.Command, args []string) {
			if task == "" {
				fmt.Println("Write which task definition to execute.")
				return
			}

			if cluster == "" {
				fmt.Println("Write which cluster to execute task definition.")
				return
			}

			if cluster != "production" && cluster != "homolog" && cluster != "staging" {
				fmt.Println("Name cluster invalid.")
				return
			}

			fmt.Printf("Task Definition: %s\nCluster: %s\nEnv: %s\n", task, cluster, env)
		},
	}

	cmd.Flags().StringVarP(&task, "task", "t", "", "Task Definition Name")
	cmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Cluster Name")
	cmd.Flags().StringVarP(&env, "env", "e", "", "Env Overwrite")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}

func clusterSubnets() {

}
