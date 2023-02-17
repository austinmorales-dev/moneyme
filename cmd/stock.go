/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "returns the current market price of the requested stock",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: not enough arguments, try `moneyme -h` for all commands")
			return
		}
		stock, err := StockPrice(os.Args[2])
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%v (live): %v\n", os.Args[2], stock)
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

func StockPrice(ticker string) (float64, error) {
	q, err := quote.Get(ticker)
	if err != nil || q == nil {
		fmt.Println("Check inputted symbol for typos")
		os.Exit(1)
	}
	// if q.Symbol != ticker {
	// 	os.Exit(1)
	// }
	// if err == nil {
	// 	log.Fatal("Fatal error, check for typos in stock symbol/ticker")
	// }
	return q.RegularMarketPrice, nil
}
