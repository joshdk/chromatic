// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type browserConfig struct {
	Flags []string `yaml:"flags"`
}

type startConfig struct {
	URL string `yaml:"url"`
}

type endConfig struct {
	URL     string       `yaml:"url"`
	Title   string       `yaml:"title"`
	Cookie  cookieConfig `yaml:"cookie"`
	Timeout int64        `yaml:"timeout"`
}

type cookieConfig struct {
	Name   string `yaml:"name" json:"name"`
	Domain string `yaml:"domain" json:"domain"`
	Value  string `yaml:"-" json:"value"`
}

type Config struct {
	Browser browserConfig `yaml:"browser"`
	Start   startConfig   `yaml:"start"`
	End     endConfig     `yaml:"end"`
}

func Load(filename string) (*Config, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return LoadBytes(body)
}

func LoadBytes(body []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(body, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
