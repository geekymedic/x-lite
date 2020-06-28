package xerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var NotFound = By(errors.New("not found source"))

func TestBy(t *testing.T) {
	var err = func() error {
		return NotFound
	}()
	require.Equal(t, err, NotFound)
	require.True(t, errors.As(err, &NotFound))
	t.Log(fmt.Errorf("%w", func() error{
		return errors.New("Hello Word")
	}()))
}
