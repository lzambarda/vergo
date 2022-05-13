//nolint:revive // Self-explanatory.
package vergo

import "github.com/pkg/errors"

var (
	ErrMalformed      = errors.New("malformed semver")
	ErrMalformedMajor = errors.Wrap(ErrMalformed, "major")
	ErrMalformedMinor = errors.Wrap(ErrMalformed, "minor")
	ErrMalformedPatch = errors.Wrap(ErrMalformed, "patch")
	ErrMalformedLabel = errors.Wrap(ErrMalformed, "only labels in the form of rc<NUMBER> can be bumped")
)
