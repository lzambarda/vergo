package vergo

import (
	"testing"

	"github.com/blokur/testament"
	"github.com/stretchr/testify/assert"
)

func TestSemver(t *testing.T) {
	t.Parallel()
	t.Run("Bump", TestBump)
	t.Run("Parse", TestParse)
	t.Run("After", TestAfter)
	t.Run("Before", TestBefore)
	t.Run("Equals", TestEquals)
}

func TestBump(t *testing.T) {
	t.Parallel()
	runs := map[string]struct {
		input    *Semver
		bump     Bump
		expected *Semver
	}{
		"invalid": {
			New(1, 2, 3),
			BumpInvalid,
			New(1, 2, 3),
		},
		"patch": {
			New(1, 2, 9),
			BumpPatch,
			New(1, 2, 10),
		},
		"minor": {
			New(3, 4, 6),
			BumpMinor,
			New(3, 5, 0),
		},
		"major": {
			New(0, 4, 2),
			BumpMajor,
			New(1, 0, 0),
		},
	}

	for run, data := range runs {
		data := data
		t.Run(run, func(t *testing.T) {
			t.Parallel()
			data.input.Bump(data.bump)
			assert.EqualValues(t, data.expected, data.input)
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	runs := map[string]struct {
		input  string
		outErr error
		outVal *Semver
	}{
		"malformed": {
			"v1",
			ErrMalformed,
			nil,
		},
		"malformed major": {
			"va.0.0",
			ErrMalformedMajor,
			nil,
		},
		"malformed minor": {
			"v1.123a.0",
			ErrMalformedMinor,
			nil,
		},
		"malformed patch": {
			"v1.10.fef",
			ErrMalformedPatch,
			nil,
		},
		"success no v": {
			"1.2.3",
			nil,
			&Semver{1, 2, 3, false},
		},
		"success with v": {
			"v41.52.63",
			nil,
			&Semver{41, 52, 63, true},
		},
		"success pad": {
			"  v1.2.3 ",
			nil,
			&Semver{1, 2, 3, true},
		},
	}

	for run, data := range runs {
		data := data
		t.Run(run, func(t *testing.T) {
			t.Parallel()
			gotVal, gotErr := ParseSemver(data.input)
			if data.outErr == nil {
				assert.NoError(t, gotErr)
			} else {
				testament.AssertInError(t, gotErr, data.outErr)
			}
			assert.EqualValues(t, data.outVal, gotVal)
		})
	}
}

func TestAfter(t *testing.T) {
	t.Parallel()
	runs := map[string]struct {
		a        *Semver
		b        *Semver
		expected bool
	}{
		"major false": {
			&Semver{1, 0, 0, false},
			&Semver{2, 0, 0, false},
			false,
		},
		"major true": {
			&Semver{2, 0, 0, false},
			&Semver{1, 0, 0, false},
			true,
		},
		"minor false": {
			&Semver{0, 1, 0, false},
			&Semver{0, 2, 0, false},
			false,
		},
		"minor true": {
			&Semver{0, 2, 0, false},
			&Semver{0, 1, 0, false},
			true,
		},
		"patch false": {
			&Semver{0, 0, 1, false},
			&Semver{0, 0, 2, false},
			false,
		},
		"patch true": {
			&Semver{0, 0, 2, false},
			&Semver{0, 0, 1, false},
			true,
		},
		"equal false": {
			&Semver{1, 2, 3, false},
			&Semver{1, 2, 3, false},
			false,
		},
	}

	for run, data := range runs {
		data := data
		t.Run(run, func(t *testing.T) {
			t.Parallel()
			got := data.a.After(data.b)
			assert.Equal(t, data.expected, got)
		})
	}
}

func TestBefore(t *testing.T) {
	t.Parallel()
	runs := map[string]struct {
		a        *Semver
		b        *Semver
		expected bool
	}{
		"major false": {
			&Semver{2, 0, 0, false},
			&Semver{1, 0, 0, false},
			false,
		},
		"major true": {
			&Semver{1, 0, 0, false},
			&Semver{2, 0, 0, false},
			true,
		},
		"minor false": {
			&Semver{0, 2, 0, false},
			&Semver{0, 1, 0, false},
			false,
		},
		"minor true": {
			&Semver{0, 1, 0, false},
			&Semver{0, 2, 0, false},
			true,
		},
		"patch false": {
			&Semver{0, 0, 2, false},
			&Semver{0, 0, 1, false},
			false,
		},
		"patch true": {
			&Semver{0, 0, 1, false},
			&Semver{0, 0, 2, false},
			true,
		},
		"equal false": {
			&Semver{1, 2, 3, false},
			&Semver{1, 2, 3, false},
			false,
		},
	}

	for run, data := range runs {
		data := data
		t.Run(run, func(t *testing.T) {
			t.Parallel()
			got := data.a.Before(data.b)
			assert.Equal(t, data.expected, got)
		})
	}
}

func TestEquals(t *testing.T) {
	t.Parallel()
	runs := map[string]struct {
		a        *Semver
		b        *Semver
		expected bool
	}{
		"major false": {
			&Semver{1, 0, 0, false},
			&Semver{2, 0, 0, false},
			false,
		},
		"minor false": {
			&Semver{0, 1, 0, false},
			&Semver{0, 2, 0, false},
			false,
		},
		"patch false": {
			&Semver{0, 0, 1, false},
			&Semver{0, 0, 2, false},
			false,
		},
		"true": {
			&Semver{1, 2, 3, false},
			&Semver{1, 2, 3, false},
			true,
		},
		"true diff v": {
			&Semver{1, 2, 3, true},
			&Semver{1, 2, 3, false},
			true,
		},
	}

	for run, data := range runs {
		data := data
		t.Run(run, func(t *testing.T) {
			t.Parallel()
			got := data.a.Equals(data.b)
			assert.Equal(t, data.expected, got)
		})
	}
}
