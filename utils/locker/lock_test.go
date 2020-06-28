package locker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErr(t *testing.T) {
	var err = func() error{
		return ErrNotHeld
	}()
	require.Equal(t, err, ErrNotHeld)
	require.True(t, errors.Is(err, ErrNotHeld))
	require.True(t, errors.As(err, &ErrNotHeld))
	t.Logf("%+q", err)
}