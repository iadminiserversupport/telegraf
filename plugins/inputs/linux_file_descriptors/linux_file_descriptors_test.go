package linux_file_descriptors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFileNr(t *testing.T) {
	fields, err := parseFileNr("13223\t0\t9223372036854775807\n")
	require.NoError(t, err)

	require.Equal(t, uint64(13223), fields["allocated"])
	require.Equal(t, uint64(0), fields["unused"])
	require.Equal(t, uint64(9223372036854775807), fields["maximum"])
	require.InDelta(t, 0.0, fields["used_percent"], 0.000001)
}

func TestParseFileNrInvalid(t *testing.T) {
	_, err := parseFileNr("13223\t0\n")
	require.Error(t, err)
}

func TestParseFileNrZeroMaximum(t *testing.T) {
	_, err := parseFileNr("13223\t0\t0\n")
	require.Error(t, err)
}
