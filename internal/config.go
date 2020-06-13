// Copyright (c) 2000-present Liferay, Inc. All rights reserved.
//
// This library is free software; you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation; either version 2.1 of the License, or (at your option)
// any later version.
//
// This library is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.

package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const fileName = "config.yaml"

var dbContainerNames = map[string]string{
	"ce":       "db-ce",
	"commerce": "db-commerce",
	"dxp":      "db-dxp",
	"nightly":  "db-nightly",
	"release":  "db-release",
}
var portalContainerNames = map[string]string{
	"ce":       "lpn-ce",
	"commerce": "lpn-commerce",
	"dxp":      "lpn-dxp",
	"nightly":  "lpn-nightly",
	"release":  "lpn-release",
}

var dbImages = map[string]ImageConfig{
	"mysql": {
		Image: "mdelapenya/mysql-utf8",
		Tag:   "5.7",
	},
	"postgres": {
		Image: "postgres",
		Tag:   "9.6-alpine",
	},
}
var portalImages = map[string]ImageConfig{
	"ce": {
		Image: "liferay/portal",
		Tag:   "7.0.6-ga7",
	},
	"commerce": {
		Image: "liferay/commerce",
		Tag:   "1.1.1",
	},
	"dxp": {
		Image: "liferay/dxp",
		Tag:   "7.0.10.8",
	},
	"nightly": {
		Image: "liferay/portal-snapshot",
		Tag:   "master",
	},
	"release": {
		Image: "mdelapenya/liferay-portal",
		Tag:   "latest",
	},
}

// ImageConfig image configuration
type ImageConfig struct {
	Image string `yaml:"image"`
	Tag   string `yaml:"tag"`
}

// ImagesConfig image configuration
type ImagesConfig struct {
	Db     map[string]ImageConfig `mapstructure:"db"`
	Portal map[string]ImageConfig `mapstructure:"portal"`
}

// LPNConfig tool configuration
type LPNConfig struct {
	Container NamesConfig  `mapstructure:"container"`
	Images    ImagesConfig `mapstructure:"images"`
}

// GetDbContainerName name of the container for databases
func (c *LPNConfig) GetDbContainerName(t string) string {
	return c.Container.Names.Db[t]
}

// GetDbImageName name of the image used to run the portal
func (c *LPNConfig) GetDbImageName(t string) string {
	return c.Images.Db[t].Image
}

// GetDbImageTag name of the image used to run the portal
func (c *LPNConfig) GetDbImageTag(t string) string {
	return c.Images.Db[t].Tag
}

// GetPortalImageName name of the image used to run the portal
func (c *LPNConfig) GetPortalImageName(t string) string {
	return c.Images.Portal[t].Image
}

// GetPortalImageTag name of the image used to run the portal
func (c *LPNConfig) GetPortalImageTag(t string) string {
	return c.Images.Portal[t].Tag
}

// GetPortalContainerName name of the container for portal
func (c *LPNConfig) GetPortalContainerName(t string) string {
	return c.Container.Names.Portal[t]
}

// NamesConfig container configuration
type NamesConfig struct {
	Names NameConfig `mapstructure:"names"`
}

// NameConfig container configuration
type NameConfig struct {
	Db     map[string]string `mapstructure:"db"`
	Portal map[string]string `mapstructure:"portal"`
}

func initConfigFile(workspace string, configFile string, defaults map[string]interface{}) *os.File {
	log.Printf("Creating %s with default values in %s.", configFile, workspace)

	configFilePath := filepath.Join(workspace, configFile)

	f, _ := os.Create(configFilePath)

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(workspace)

	err := v.WriteConfig()
	if err != nil {
		log.Fatalf(`Cannot save default configuration file at %s: %v`, configFilePath, err)
	}

	return f
}

// NewConfig returns a new configuration
func NewConfig(workspace string) *LPNConfig {
	lpnConfig, err := readConfig(workspace, fileName, map[string]interface{}{
		"container": map[string]interface{}{
			"names": map[string]interface{}{
				"db":     dbContainerNames,
				"portal": portalContainerNames,
			},
		},
		"images": map[string]interface{}{
			"db":     dbImages,
			"portal": portalImages,
		},
	})
	if err != nil {
		log.Fatalf("Error when reading config: %v\n", err)
	}

	return &lpnConfig
}

func readConfig(
	workspace string, configFile string, defaults map[string]interface{}) (LPNConfig, error) {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(workspace)

	err := viper.ReadInConfig()
	if err != nil {
		initConfigFile(workspace, configFile, defaults)
		viper.ReadInConfig()
	}

	var lpnConfig LPNConfig
	err = viper.Unmarshal(&lpnConfig)
	if err != nil {
		log.Fatalf("Unable to decode configuration into struct, %v", err)
	}

	return lpnConfig, err
}
