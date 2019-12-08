package main

import (
	"github.com/spf13/cobra"
	"github.com/uzxmx/k8s-event-generator/pkg/generator"
	"k8s.io/klog"
)

func main() {
	gen := generator.NewGenerator()

	cmd := &cobra.Command{
		Use:   "k8s-event-generator",
		Short: "Generate fake event in kubernetes cluster",
		Long: `
k8s-event-generator is an utility that can help you generate fake event in kubernetes
cluster, especially useful when you use event-based tools like brigade.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := gen.Run(); err != nil {
				klog.Error(err)
			}
		},
	}
	gen.AddFlags(cmd)

	if err := cmd.Execute(); err != nil {
		klog.Error(err)
	}
}
