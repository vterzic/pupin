/*
Copyright © 2024 <veljkot@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/vterzic/pupin/internal/menu"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <config-path>",
	Short: "Renders menu based on the provided configuration",
	Long: `Renders menu based on the provided configuration. 
	
Menu configuration path must be provided as an argument. 
Make sure that your config file is aligned with config-schema.json

Example usage:
pupin run ./config-example.json

For more information, visit https://github.com/vterzic/pupin/blob/main/readme.md`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Missing config path argument")
			fmt.Println("Run pupin --help")
			os.Exit(0)
		}

		initMenu(args)
	},
}

func initMenu(args []string) {

	path := args[0]
	menu := menu.FromPath(path)
	render(menu)
}

func render(m *menu.Menu) {
	var index = -1

	labels := generateBreadcrumbs(m)
	items := generateItems(m)

	searcher := func(input string, index int) bool {
		opt := items[index]
		name := strings.Replace(strings.ToLower(opt), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:    labels,
		Items:    items,
		Searcher: searcher,
	}

	for index < 0 {
		var err error
		index, _, err = prompt.Run()

		if err != nil {
			os.Exit(1)
		}
	}

	// back is added on top of options
	if index >= len(m.Current().Options) {
		if m.IsRoot() {
			fmt.Println("Bye!")
			os.Exit(0)
		} else {
			m.GoBack()
			render(m)
		}
	} else {
		navigatedMenu := m.Navigate(index)
		if navigatedMenu == nil {
			runCommand(m.Current().Options[index].Command)
		} else {
			render(m)
		}
	}
}

func runCommand(c string) {
	command := strings.Split(strings.Trim(c, " "), " ")

	var shCmd = exec.Command(command[0], command[1:]...)
	shCmd.Stdin = os.Stdin
	shCmd.Stdout = os.Stdout
	shCmd.Stderr = os.Stderr

	err := shCmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateBreadcrumbs(m *menu.Menu) string {
	items := m.Breadcrumbs()

	if len(items) == 1 {
		return items[0]
	}

	breadcrumbs := ""

	for index, item := range items {
		if index != len(items)-1 {
			item = item + " > "
		}

		breadcrumbs += item
	}

	return breadcrumbs
}

func generateItems(m *menu.Menu) []string {
	items := []string{}
	menuItems := m.Current().Options

	for _, item := range menuItems {
		items = append(items, item.Name)
	}

	if !m.IsRoot() {
		items = append(items, "Back")
	} else {
		items = append(items, "Exit")
	}

	return items
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// menuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
