/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"context"
	"fmt"

	"github.com/samanthapuri/kubeui/clientset"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// podCmd represents the pod command
var kubeconfig *string
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset := clientset.NewClientSet()
		deploy, err := clientset.AppsV1().Deployments("kube-system").Get(context.TODO(), "coredns", v1.GetOptions{})
		if err != nil {
			fmt.Println("error getting deployments")
		}
		fmt.Println(deploy.Name, deploy.CreationTimestamp, deploy.Namespace)
	},
}

func init() {
	GetCmd.AddCommand(podCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
