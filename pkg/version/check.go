package version

import "github.com/hashicorp/go-version"

func Check(v string) (string, error) {
	cv, err := version.NewVersion(v)
	if err != nil {
		return "", err
	}

	return cv.String(),nil
}
