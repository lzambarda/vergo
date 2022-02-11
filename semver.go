// Package vergo is a simple tool to parse, compare and bump semver.
package vergo

import (
	"fmt"
	"strconv"
	"strings"
)

// Bump describes a modify operation on a semver object.
type Bump int

//nolint:revive // Self-explanatory.
const (
	BumpInvalid Bump = iota
	BumpPatch
	BumpMinor
	BumpMajor
)

var bumpString = [4]string{"invalid", "patch", "minor", "major"}

func (b Bump) String() string {
	return bumpString[b]
}

// ParseBump returns a Bump equivalent to the given string.
func ParseBump(s string) (Bump, error) {
	for i, v := range bumpString {
		if s == v {
			return Bump(i), nil
		}
	}
	return BumpInvalid, fmt.Errorf("unable to parse %q as Bump", s)
}

// Semver represents a... Semver!?
type Semver struct {
	Major int
	Minor int
	Patch int
	hasV  bool
}

// New correctly instantiates a semver instance using the given major, minor and
// patch.
func New(major, minor, patch int) *Semver {
	return &Semver{major, minor, patch, false}
}

func (s *Semver) String() string {
	v := ""
	if s.hasV {
		v = "v"
	}
	return fmt.Sprintf("%s%d.%d.%d", v, s.Major, s.Minor, s.Patch)
}

// ParseSemver returns a semver from the given string.
// This is capable of deadling with prependend "v" but can't yet process any
// suffix.
func ParseSemver(semver string) (s *Semver, err error) {
	s = &Semver{}
	semver = strings.TrimSpace(semver)
	if strings.HasPrefix(semver, "v") {
		s.hasV = true
		semver = semver[1:]
	}
	parts := strings.Split(semver, ".")
	if len(parts) != 3 {
		return nil, ErrMalformed
	}
	s.Major, err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, ErrMalformedMajor
	}
	s.Minor, err = strconv.Atoi(parts[1])
	if err != nil {
		return nil, ErrMalformedMinor
	}
	s.Patch, err = strconv.Atoi(parts[2])
	if err != nil {
		return nil, ErrMalformedPatch
	}
	return s, nil
}

// Bump applies the given bump to this semver.
func (s *Semver) Bump(b Bump) {
	ns := s.PeekBump(b)
	*s = ns
}

// PeekBump returns a copy of this semver with the applied Bump.
// Applying an invalid bump will result in no op.
func (s *Semver) PeekBump(b Bump) Semver {
	ns := Semver{
		s.Major,
		s.Minor,
		s.Patch,
		s.hasV,
	}
	switch b {
	case BumpPatch:
		ns.Patch++
	case BumpMinor:
		ns.Patch = 0
		ns.Minor++
	case BumpMajor:
		ns.Patch = 0
		ns.Minor = 0
		ns.Major++
	}
	return ns
}

// UnmarshalYAML sets the fields in the unmarshal target according to the
// content of the stringified version.
func (s *Semver) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawVersion string
	err := unmarshal(&rawVersion)
	if err != nil {
		return err
	}
	ns, err := ParseSemver(rawVersion)
	if err != nil {
		return err
	}
	*s = *ns
	return nil
}

// MarshalYAML returns the formatted stringified version of this semver.
//nolint:unparam // False negative, can't do much about it.
func (s *Semver) MarshalYAML() (interface{}, error) {
	return s.String(), nil
}

// After returns whether this current semver is a newer version of the other.
func (s *Semver) After(other *Semver) bool {
	if s.Major > other.Major {
		return true
	}
	if s.Major < other.Major {
		return false
	}
	if s.Minor > other.Minor {
		return true
	}
	if s.Minor < other.Minor {
		return false
	}
	if s.Patch > other.Patch {
		return true
	}
	return false
}

// Before returns whether this current semver is an older version of the other.
func (s *Semver) Before(other *Semver) bool {
	if s.Major < other.Major {
		return true
	}
	if s.Major > other.Major {
		return false
	}
	if s.Minor < other.Minor {
		return true
	}
	if s.Minor > other.Minor {
		return false
	}
	if s.Patch < other.Patch {
		return true
	}
	return false
}

// Equals only checks equality of version, not whether they both have a
// prepended "v" or not.
func (s *Semver) Equals(other *Semver) bool {
	return s.Patch == other.Patch && s.Minor == other.Minor && s.Major == other.Major
}
