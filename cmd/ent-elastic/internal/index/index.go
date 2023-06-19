package index

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdIndex Index command.
var CmdIndex = &cobra.Command{
	Use:   "index",
	Short: "new elasticsearch",
	Long:  "new elasticsearch. Example: go-elasticsearch index xxxx",
	Run:   Index,
}

func Index(cmd *cobra.Command, args []string) {
	fmt.Println("aaaaa")
}
