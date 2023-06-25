package main

import (
	"github.com/guihouchang/go-elasticsearch/cmd/ent-elastic/internal/generate"
	"github.com/spf13/cobra"
	"log"
)

const release = "v1.0.5"

var rootCmd = &cobra.Command{
	Use:     "ent-elastic",
	Short:   "ent-elastic: Tool for generating elastic search orm",
	Long:    "ent-elastic: Tool for generating elastic search orm",
	Version: release,
}

func init() {
	rootCmd.AddCommand(generate.GenCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
