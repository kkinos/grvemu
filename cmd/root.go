package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kinpoko/grvemu/rv32i"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grvemu [binary file]",
	Short: "RISC-V emulator for cli written in Go",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		arch, err := cmd.Flags().GetString("arch")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		test, err := cmd.Flags().GetBool("test")
		if err != nil {
			return err
		}
		switch arch {
		case "rv32i":
			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer file.Close()

			binary, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			end, err := cmd.Flags().GetUint32("eof")
			if err != nil {
				return err
			}

			if err := rv32i.Run(binary, end, debug, test); err != nil {
				return err
			}

		// TODO support rv64i

		default:
			fmt.Printf("%s is not supported\n", arch)
			fmt.Printf("this emulator supports rv32i\n")
		}
		return nil

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("arch", "a", "rv32i", "architecture")
	rootCmd.Flags().BoolP("debug", "d", false, "debug mode")
	rootCmd.Flags().BoolP("test", "t", false, "display global pointer")
	rootCmd.Flags().Uint32P("eof", "e", 0x0, "end of binary")

}
