package main

import (
	"errors"
	"github.com/lazyguru/AdventOfCode/utils/pkg/download"
	"github.com/lazyguru/AdventOfCode/utils/pkg/submit"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var year, day, part int
var answer string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Get a puzzle",
	Long:  "Get a puzzle from year, day, and part inputs",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			year = getNow().Year()
		}
		if day <= 0 {
			if getNow().Month() == 12 {
				day = getToday()
			} else {
				return errors.New("invalid day")
			}
		}
		if part <= 0 {
			part = 1
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		download.SetupPuzzle(year, day, part)
	},
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports getting puzzle data, including inputs directly from the website, and submitting answers",
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit an answer",
	Long:  "Submit an answer for a given year, day, and part",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			year = getNow().Year()
		}
		if day <= 0 {
			if getNow().Month() == 12 {
				day = getToday()
			} else {
				return errors.New("invalid day")
			}
		}
		if part <= 0 {
			part = 1
		}
		if strings.TrimSpace(answer) == "" {
			return errors.New("missing answer")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		submit.SubmitAnswer(year, day, part, answer)
	},
}

func getNow() time.Time {
	location, _ := time.LoadLocation("America/New_York")
	return time.Now().In(location)
}

func getToday() int {
	location, _ := time.LoadLocation("America/New_York")
	return time.Now().In(location).Day()
}

func main() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "year input")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "day input")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", 0, "part input")
	rootCmd.PersistentFlags().StringVarP(&answer, "answer", "a", "", "answer input")

	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(submitCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
