[![CircleCI][circleci-badge]][circleci-link]
[![Go Report Card][go-report-card-badge]][go-report-card-link]
[![License][license-badge]][license-link]
[![GitHub release][github-release-badge]][github-release-link]

# Chromatic

🍪 Configurable human assisted Chrome automation

![chromatic-demo](https://user-images.githubusercontent.com/307183/43058824-72d6681e-8dfe-11e8-94d2-2e71d847d425.gif)

## Installing

### From source

You can use `go get` to install this tool by running:

```bash
$ go get -u github.com/joshdk/chromatic
```

### Release binary

Alternatively, you can download an OSX [release][github-release-link] binary by running:

```bash
$ wget -O chromatic -q https://github.com/joshdk/chromatic/releases/download/0.1.0/chromatic_darwin_amd64
$ sudo install chromatic /usr/local/bin
```

## Motivations

There are certain services that require programmatic interaction, but otherwise make it difficult. Web-only APIs, [CAPTCHAs](https://en.wikipedia.org/wiki/CAPTCHA), and MFA are a few examples of things that might not work well with typical automation.

This tool, `chromatic`, implements an escape-hatch around these difficulties. It can be configured to open up a Chrome window to a given starting URL, and will continue to be interactive until a given set of ending conditions are met, at which time the Chrome window will close. This user interaction is typically a login flow.

When `chromatic` exits, it will dump details about the last visited page as json. This json (and the session data contained within) is meant to be consumed by additional automation.

## Usage

In this example we will use `chromatic` to extract and consume a web session for CircleCI.

### Configuration

The following [`examples/circleci.yml`](https://github.com/joshdk/chromatic/blob/master/examples/circleci.yml) will launch Chrome to CircleCI's Github/Bitbucket authorization page.

```yaml
start:
  url: https://circleci.com/vcs-authorize/

end:
  timeout: 60
  url: https://circleci.com/dashboard
  cookie:
    name: ring-session
    domain: circleci.com
  title: CircleCI
```

### Interaction

You can then login to Github using your username, password, & maybe an MFA code. When you eventually login and are redirected to your CircleCI dashboard, `chromatic` will check if the desired URL, title, & cookie match the current page.

### Output

Once a match is found, Chrome will exit successfully and `chromatic` will dump details about the last page as json.

```
$ chromatic examples/circleci.yml | tee page.json
{
  "url": "https://circleci.com/dashboard",
  "title": "CircleCI",
  "cookies": [
    {
      "name": "ring-session",
      "domain": "circleci.com",
      "value": "vKAm...72Af"
    }
    ...
  ]
}
```

### Consumption

Using this output, the session value can be stored and later used with the web API.

```
$ SESSION="$(cat page.json | jq -r '.cookies[] | select(.name=="ring-session") | .value')"

$ echo $SESSION
vKAm...72Af

$ curl -H "Cookie: ring-session=$SESSION" https://circleci.com/api/v1/me
{
  "name" : "...",
  "selected_email" : "...",
  ...
}
```

## License

This library is distributed under the [MIT License][license-link], see [LICENSE.txt][license-file] for more information.

[circleci-badge]:         https://circleci.com/gh/joshdk/chromatic.svg?&style=shield
[circleci-link]:          https://circleci.com/gh/joshdk/chromatic/tree/master
[github-release-badge]:   https://img.shields.io/github/release/joshdk/chromatic.svg
[github-release-link]:    https://github.com/joshdk/chromatic/releases/latest
[go-report-card-badge]:   https://goreportcard.com/badge/github.com/joshdk/chromatic
[go-report-card-link]:    https://goreportcard.com/report/github.com/joshdk/chromatic
[license-badge]:          https://img.shields.io/badge/license-MIT-green.svg
[license-file]:           https://github.com/joshdk/chromatic/blob/master/LICENSE.txt
[license-link]:           https://opensource.org/licenses/MIT

