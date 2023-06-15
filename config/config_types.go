package config

import (
	"fmt"
)

type Group string

type Directory string

type Config struct {
	CurrentGroup Group
	Groups       map[Group][]Directory
}

func (c *Config) GetDirectoriesFromGroup(group Group) ([]Directory, error) {
	if dirs, ok := c.Groups[group]; ok {
		return dirs, nil
	}
	return nil, fmt.Errorf("group %s does not exist", group)
}

func (c *Config) AddGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; exists {
		return fmt.Errorf("group %s already exists", groupName)
	}

	c.Groups[groupName] = []Directory{}
	return nil
}

func (c *Config) RemoveGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	delete(c.Groups, groupName)
	return nil
}

func (c *Config) SetCurrentGroup(groupName Group) error {
	if _, exists := c.Groups[groupName]; !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	c.CurrentGroup = groupName
	return nil
}

func (c *Config) AddDirectoryToGroup(group Group, dir Directory) error {
	if _, exists := c.Groups[group]; !exists {
		return fmt.Errorf("group %s does not exist", group)
	}

	c.Groups[group] = append(c.Groups[group], dir)
	return nil
}

func (c *Config) ApplyChanges() error {
	err := SaveConfig(c)
	if err != nil {
		return fmt.Errorf("error saving config: %v", err)
	}

	return nil
}

func ToGroup(value string) Group {
	return Group(value)
}

func ToDirectory(value string) Directory {
	return Directory(value)
}
