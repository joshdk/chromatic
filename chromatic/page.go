// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import "github.com/mafredri/cdp/protocol/network"

type Page struct {
	Title   string
	URL     string
	Cookies []network.Cookie
	Body    string
}
