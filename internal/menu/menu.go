package menu

import (
	"encoding/json"
	"fmt"
	"os"
)

type MenuItem struct {
	Name    string
	Command string
	Options []*MenuItem
}

type Menu struct {
	History []*MenuItem
}

func (m *Menu) Current() *MenuItem {
	return m.History[len(m.History)-1]
}

func (m *Menu) IsRoot() bool {
	return m.Current() == m.History[0]
}

func (m *Menu) GoBack() *MenuItem {
	if len(m.History) == 1 {
		return m.Current()
	}

	m.History = m.History[:len(m.History)-1]

	return m.Current()
}

func (m *Menu) Breadcrumbs() []string {
	breadcrumbs := []string{}

	for _, menu := range m.History {
		breadcrumbs = append(breadcrumbs, menu.Name)
	}

	return breadcrumbs
}

func (m *Menu) Navigate(index int) *MenuItem {
	selected := m.Current().Options[index]

	if selected.Options == nil {
		// it is a command
		return nil
	}

	m.History = append(m.History, selected)

	return selected
}

func FromPath(filepath string) *Menu {
	rootMenu := loadJson(filepath)

	menu := &Menu{
		History: []*MenuItem{rootMenu},
	}

	return menu
}

func loadJson(filepath string) *MenuItem {
	var file, err = os.ReadFile(filepath)
	var menuItem MenuItem

	if err != nil {
		fmt.Printf("Could not retrieve a file on path %s\n", filepath)
		os.Exit(1)
	}

	if json.Unmarshal(file, &menuItem) != nil {
		fmt.Printf("Could not parse menu. Validate your config file against config-schema.json")
		os.Exit(1)
	}

	return &menuItem
}
