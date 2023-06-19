package config

import (
	"fmt"
)

type Group string
type FilePath string
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
	groupDirs, exists := c.Groups[group]
	if !exists {
		return fmt.Errorf("group %s does not exist", group)
	}

	for _, existingDir := range groupDirs {
		if existingDir == dir {
			return fmt.Errorf("directory `%s` already exists in group `%s`", dir, group)
		}
	}

	c.Groups[group] = append(groupDirs, dir)
	return nil
}

func (c *Config) PrintDirectoriesFromCurrentGroup() {
	currentGroupDirectoryList, exists := c.Groups[c.CurrentGroup]
	if !exists {
		fmt.Printf("Group %s does not exist\n", c.CurrentGroup)
		return
	}

	if len(currentGroupDirectoryList) == 0 {
		fmt.Printf("Your current group `%s` have no directories into it!\n", c.CurrentGroup)
		return
	}

	fmt.Printf("Directories from current group `%s`:\n", c.CurrentGroup)
	for _, dir := range currentGroupDirectoryList {
		fmt.Println(dir)
	}
}

func (c *Config) RemoveDirectoryFromCurrentGroup(dir Directory) error {
	groupDirs, exists := c.Groups[c.CurrentGroup]
	if !exists {
		return fmt.Errorf("group %s does not exist", c.CurrentGroup)
	}

	// Check if the directory exists in the group
	index := -1
	for i, existingDir := range groupDirs {
		if existingDir == dir {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("directory %s does not exist in group %s", dir, c.CurrentGroup)
	}

	// Remove the directory from the group
	c.Groups[c.CurrentGroup] = append(groupDirs[:index], groupDirs[index+1:]...)
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
