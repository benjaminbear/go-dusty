package sshhandler

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	giturls "github.com/whilp/git-urls"
)

// ParseUrl parses the url
// Only http(s) and ssh is allowed
// Tries to parse unknown format to ssh
func ParseUrl(url string) (*url.URL, error) {
	triedToFix := false
	for {
		parsedUrl, err := giturls.Parse(url)
		if err != nil {
			return parsedUrl, err
		}

		switch parsedUrl.Scheme {
		case "http":
			fallthrough
		case "https":
			return parsedUrl, nil
		case "ssh":
			return parsedUrl, nil
		case "file":
			if triedToFix {
				return parsedUrl, fmt.Errorf("cannot handle url format %v", url)
			}

			url = fixUrl(url)
			triedToFix = true
		default:
			return parsedUrl, fmt.Errorf("cannot handle url format %v", url)
		}
	}
}

// fixUrl tries to fix a git url
// It makes a gitUrl out of a normal url without a scheme
func fixUrl(url string) string {
	urlParts := strings.SplitN(url, "/", 2)
	if len(urlParts) < 2 {
		return ""
	}

	return urlParts[0] + ":" + urlParts[1]
}

// BuildRepoPath splits a url path into its components
// example: https://github.com/owner/project.git will turn
// into ./github.com/owner/project
func BuildRepoPath(url *url.URL) (string, error) {
	if url.Host == "" {
		return "", fmt.Errorf("invalid url, host field is empty")
	}

	path := url.Host
	splitPath := strings.Split(strings.TrimSuffix(url.Path, ".git"), "/")

	for _, elem := range splitPath {
		if elem != "" {
			path = filepath.Join(path, elem)
		}
	}

	return path, nil
}
