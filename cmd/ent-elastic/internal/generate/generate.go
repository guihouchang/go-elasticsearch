package generate

import (
	_ "embed"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/schema/gen"
	"github.com/spf13/cobra"
	"log"
)

var GenCmd = &cobra.Command{
	Use:     "generate [flags] path",
	Short:   "generate go code for the schema directory",
	Example: "ent-elastic generate  ./es/schema",
	Run:     Gen,
}

func Gen(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln(fmt.Errorf("schema error for args"))
	}

	path := args[0]
	gen.Gen(path)
}
