/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"os"

	utils "github.com/matrixbegins/klauncher/core"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Schedules a Kubernetes CronJob.",
	Long:  `Schedules a Kubernetes CronJob.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scheduling the CronJob....." + utils.GoDotEnvVariable("PROFILER_NAME"))

		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		jobClient := clientset.BatchV1().CronJobs(apiv1.NamespaceDefault)

		results, err := jobClient.Create(context.TODO(), utils.GetCronJobSpec(), metav1.CreateOptions{})

		if err != nil {
			panic(err)
		}
		fmt.Printf("Created deployment %q.\n", results.GetObjectMeta().GetName())

	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	fmt.Println(os.Environ())
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
