// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"encoding/json"
)

type output struct {
	URL     string         `json:"url"`
	Title   string         `json:"title"`
	Cookies []cookieConfig `json:"cookies,omitempty"`
}

func Report(result Page) (string, error) {
	var o = output{
		URL:     result.URL,
		Title:   result.Title,
		Cookies: make([]cookieConfig, len(result.Cookies)),
	}

	for index, cookie := range result.Cookies {
		o.Cookies[index] = cookieConfig{
			cookie.Name,
			cookie.Domain,
			cookie.Value,
		}
	}

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
