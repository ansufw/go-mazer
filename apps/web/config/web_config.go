package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

// SidebarItem represents a single item in the sidebar menu
type SidebarItem struct {
	Name    string        `json:"name"`
	IsTitle *bool         `json:"isTitle,omitempty"`
	Key     string        `json:"key,omitempty"`
	Icon    string        `json:"icon,omitempty"`
	URL     string        `json:"url,omitempty"`
	Submenu []SidebarItem `json:"submenu,omitempty"`
}

// IsTitleItem checks if this item is a title
func (s *SidebarItem) IsTitleItem() bool {
	return s.IsTitle != nil && *s.IsTitle
}

// HasSubmenu checks if this item has a submenu
func (s *SidebarItem) HasSubmenu() bool {
	return len(s.Submenu) > 0
}

// IsLink checks if this item is a simple link
func (s *SidebarItem) IsLink() bool {
	return s.URL != "" && !s.HasSubmenu()
}

type Sidebar struct {
	Items []SidebarItem `json:"items"`
}

var (
	sidebarCache []SidebarItem
	cacheOnce    sync.Once
)

// LoadSidebar loads sidebar configuration from embedded data with caching
func LoadSidebar() ([]SidebarItem, error) {
	cacheOnce.Do(func() {
		sidebarCache = SidebarData
	})

	return sidebarCache, nil
}

// ReloadSidebar forces a reload of the sidebar configuration
func ReloadSidebar() ([]SidebarItem, error) {
	cacheOnce = sync.Once{} // Reset the sync.Once
	sidebarCache = nil
	return LoadSidebar()
}

// LoadConfig loads general configuration using Viper
func LoadConfig(name, format string) error {
	viper.SetConfigName(name)
	viper.SetConfigType(format)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	return nil
}
