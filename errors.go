//nolint:revive // Self-explanatory.
package vergo

import "github.com/pkg/errors"

var (
	ErrMalformed      = errors.New("malformed semver")
	ErrMalformedMajor = errors.Wrap(ErrMalformed, "major")
	ErrMalformedMinor = errors.Wrap(ErrMalformed, "minor")
	ErrMalformedPatch = errors.Wrap(ErrMalformed, "patch")
)
