package config

import "fmt"

type Group string

type Directory string

type Config struct {
	CurrentGroup Group
	Groups       map[Group][]Directory
}

func (c *Config) GetDirectoriesFromGroup(group Group) []Directory {
	return c.Groups[group]
}

func (c *Config) AddGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; exists {
		return fmt.Errorf("group %s already exists", groupName)
	}

	c.Groups[groupName] = []Directory{}

	err := SaveConfig(c)
	if err != nil {
		return fmt.Errorf("Error saving config: %v", err)
	}

	return nil
}

func (c *Config) RemoveGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	delete(c.Groups, groupName) // Remove the group

	err := SaveConfig(c)
	if err != nil {
		return fmt.Errorf("error saving config: %v", err)
	}

	return nil
}

func (c *Config) SetCurrentGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	c.CurrentGroup = groupName // Set the group

	err := SaveConfig(c)
	if err != nil {
		return fmt.Errorf("error saving config: %v", err)
	}

	return nil
}

func ToGroup(value string) Group {
	return Group(value)
}
