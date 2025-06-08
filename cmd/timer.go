package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "timer [minutes]",
	Short: "⏲️ Timer created using Golang",
	Long:  "⏲️ Timer created using Golang for fast compilation time and Cobra for superior CLI experience.",
	Run:   Tick,
}

func init() {
	rootCmd.AddCommand(start)

	start.Flags().IntP("duration", "d", 0, "Duration of timer in minutes")
	start.Flags().StringP("color", "c", "default", "Color of timer output (red, green, blue, yellow)")
}

func Tick(cmd *cobra.Command, args []string) {
	var duration time.Duration

	colorFlag, _ := cmd.Flags().GetString("color")

	var printer *color.Color
	switch colorFlag {
	case "red":
		printer = color.New(color.FgRed)
	case "green":
		printer = color.New(color.FgGreen)
	case "blue":
		printer = color.New(color.FgBlue)
	case "yellow":
		printer = color.New(color.FgYellow)
	default:
		printer = color.New(color.FgYellow)
	}

	var printerBg *color.Color
	switch colorFlag {
	case "red":
		printerBg = color.New(color.BgRed)
	case "green":
		printerBg = color.New(color.BgGreen)
	case "blue":
		printerBg = color.New(color.BgBlue)
	case "yellow":
		printerBg = color.New(color.BgYellow)
	default:
		printerBg = color.New(color.BgYellow)
	}

	durationFlag, _ := cmd.Flags().GetInt("duration")
	if durationFlag != 60 {
		duration = time.Duration(durationFlag) * time.Minute
	} else if len(args) > 0 {
		minutes, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid number of minutes")
			return
		}
		duration = time.Duration(minutes) * time.Minute
	} else {
		cmd.Help()
		return
	}

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	printerBg.Printf("Timer started for %02d:%02d:%02d", hours, minutes, seconds)
	fmt.Println()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeLeft := duration
	for range ticker.C {
		timeLeft -= time.Second
		if timeLeft <= 0 {
			printer.Println("\nTime's up!")
			if err := beeep.Alert("Timer Complete", "Time's up❗", ""); err != nil {
				fmt.Println("Error showing notification:", err)
			}
			break
		}
		hours := int(timeLeft.Hours())
		minutes := int(timeLeft.Minutes()) % 60
		seconds := int(timeLeft.Seconds()) % 60
		printer.Printf("\r🍕 Time remaining: %02d:%02d:%02d", hours, minutes, seconds)
	}
}
