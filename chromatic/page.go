package chromatic

import "github.com/mafredri/cdp/protocol/network"

type Page struct {
	Title   string
	URL     string
	Cookies []network.Cookie
	Body    string
}
