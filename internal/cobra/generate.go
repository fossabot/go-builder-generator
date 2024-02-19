package cobra

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
)

var (
	generateOpts = generate.CLIOptions{}

	generateCmd = &cobra.Command{
		Use:    "generate",
		Short:  "Generate builders for structs arguments present in file argument.",
		PreRun: SetLogLevel,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := generate.Run(generateOpts); err != nil {
				logrus.WithContext(cmd.Context()).
					WithError(err).
					Fatal("failed to generate structs builder")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// files flag
	generateCmd.Flags().StringVarP(&generateOpts.File, "file", "f", "",
		"input files containing struct types on which generate builders")
	_ = generateCmd.MarkFlagRequired("file")

	// structs flag
	generateCmd.Flags().StringSliceVarP(&generateOpts.Structs, "structs", "s", []string{},
		"golang struct names on which generate builders (they must be contained in given input files)")
	_ = generateCmd.MarkFlagRequired("structs")

	// dest flag
	generateCmd.Flags().StringVarP(&generateOpts.Destdir, "dest", "d", ".",
		"destination directory in which generate builders")

	// no notice flag
	generateCmd.Flags().BoolVar(&generateOpts.NoNotice, "no-notice", false,
		"remove top notice ('generated by ...') in generated files")

	// no validate flag
	generateCmd.Flags().BoolVar(&generateOpts.UserValidator, "use-validator", false,
		"adds validation on build function only if validate tag is present in struct(s) tags")
}
