package version

import "github.com/hashicorp/go-version"

func Assert(v, constraint string) (bool, error) {
	cv, err := version.NewVersion(v)
	if err != nil {
		return false, err
	}
	constraints, err := version.NewConstraint(constraint)
	if err != nil {
		return false, err
	}
	return constraints.Check(cv), nil
}
