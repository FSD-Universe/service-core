// Package global
package utils

import (
	"strings"
)

type CheckVersionResult int

const (
	AllMatch CheckVersionResult = iota
	MajorUnmatch
	MinorUnmatch
	PatchUnmatch
)

type Version struct {
	major   int
	minor   int
	patch   int
	version string
}

func NewVersion(version string) *Version {
	versions := strings.Split(version, ".")
	if len(versions) < 3 {
		return nil
	}
	return &Version{
		major:   StrToInt(versions[0], 0),
		minor:   StrToInt(versions[1], 0),
		patch:   StrToInt(versions[2], 0),
		version: version,
	}
}

func (v *Version) CheckVersion(version *Version) CheckVersionResult {
	if v.major != version.major {
		return MajorUnmatch
	}
	if v.minor != version.minor {
		return MinorUnmatch
	}
	if v.patch != version.patch {
		return PatchUnmatch
	}
	return AllMatch
}

func (v *Version) String() string {
	return v.version
}
