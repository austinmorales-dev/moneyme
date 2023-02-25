/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github/austinmorales/moneyme/logic"
	"os"

	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "returns the current market price of the requested stock",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 3 {
			price := logic.StockPrice(os.Args[2])
			fmt.Printf("%v: $%v\n", os.Args[2], price)
		} else {
			fmt.Println("invalid syntax, try adding -h to the end of the command you wish to run")
		}

	},
}

func init() {
	rootCmd.AddCommand(stockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
