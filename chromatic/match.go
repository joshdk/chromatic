// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"strings"
)

func Match(actual Page, targetTitle string, targetURL string, targetCookie cookieConfig) bool {
	// If all conditions are blank, nothing will ever match.
	if targetTitle == "" && targetURL == "" && targetCookie == (cookieConfig{}) {
		return false
	}

	if targetTitle != "" {
		if !strings.Contains(actual.Title, targetTitle) {
			return false
		}
	}

	if targetURL != "" {
		if !strings.Contains(actual.URL, targetURL) {
			return false
		}
	}

	if targetCookie != (cookieConfig{}) {
		for _, cookie := range actual.Cookies {
			if cookie.Name == targetCookie.Name && cookie.Domain == targetCookie.Domain {
				return true
			}
		}
		return false
	}

	return true
}
