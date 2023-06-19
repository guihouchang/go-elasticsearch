package new

import "github.com/spf13/cobra"

// CmdNew new project command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "new elasticsearch",
	Long:  "new elasticsearch. Example: go-elasticsearch new",
	Run:   New,
}

func New(cmd *cobra.Command, args []string) {

}
