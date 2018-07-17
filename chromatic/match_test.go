package chromatic

import (
	"fmt"
	"testing"

	"github.com/mafredri/cdp/protocol/network"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		page     Page
		title    string
		url      string
		cookie   cookieConfig
		expected bool
	}{
		{
			name: "No page, no target, no match",
		},
		{
			name:     "Title target match",
			page:     Page{Title: "example | login page"},
			title:    "login",
			expected: true,
		},
		{
			name:  "Title target no match",
			page:  Page{Title: "example | sign up page"},
			title: "login",
		},
		{
			name:     "URL target match",
			page:     Page{URL: "example.com/login.html"},
			url:      "login",
			expected: true,
		},
		{
			name: "URL target no match",
			page: Page{URL: "example.com/sign-up.html"},
			url:  "login",
		},

		{
			name: "Cookie target matches",
			page: Page{Cookies: []network.Cookie{
				{Name: "SESSIONID", Domain: "example.com"},
			}},
			cookie: cookieConfig{
				Name: "SESSIONID", Domain: "example.com",
			},
			expected: true,
		},
		{
			name: "Cookie target matches one of many",
			page: Page{Cookies: []network.Cookie{
				{Name: "JSESSIONID", Domain: "example.com"},
				{Name: "SESSIONID", Domain: "example.com"},
				{Name: "PHP_SESSION_ID", Domain: "example.com"},
			}},
			cookie: cookieConfig{
				Name: "SESSIONID", Domain: "example.com",
			},
			expected: true,
		},
		{
			name: "Cookie target doesn't match",
			page: Page{Cookies: []network.Cookie{
				{Name: "JSESSIONID", Domain: "example.com"},
				{Name: "PHP_SESSION_ID", Domain: "example.com"},
			}},
			cookie: cookieConfig{
				Name: "SESSIONID", Domain: "example.com",
			},
		},
	}

	for index, test := range tests {
		name := fmt.Sprintf("Case #%d %s", index+1, test.name)
		t.Run(name, func(t *testing.T) {
			actual := Match(test.page, test.title, test.url, test.cookie)

			if actual != test.expected {
				panic(fmt.Sprintf("Expected %v, actual %v", test.expected, actual))
			}
		})

	}
}
