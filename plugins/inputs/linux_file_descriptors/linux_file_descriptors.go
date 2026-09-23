//go:generate ../../../tools/readme_config_includer/generator

package linux_file_descriptors

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

type LinuxFileDescriptors struct {
	Path string `toml:"path"`
}

func (*LinuxFileDescriptors) SampleConfig() string {
	return sampleConfig
}

func (l *LinuxFileDescriptors) Gather(acc telegraf.Accumulator) error {
	path := l.Path
	if path == "" {
		path = "/proc/sys/fs/file-nr"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	fields, err := parseFileNr(string(data))
	if err != nil {
		return err
	}

	acc.AddFields("linux_file_descriptors", fields, nil)

	return nil
}

func parseFileNr(data string) (map[string]interface{}, error) {
	fields := strings.Fields(data)
	if len(fields) != 3 {
		return nil, fmt.Errorf("expected 3 values in file-nr, got %d", len(fields))
	}

	allocated, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid allocated value: %w", err)
	}

	unused, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid unused value: %w", err)
	}

	maximum, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid maximum value: %w", err)
	}

	if maximum == 0 {
		return nil, errors.New("maximum file descriptors is zero")
	}

	usedPercent := float64(allocated) / float64(maximum) * 100

	return map[string]interface{}{
		"allocated":    allocated,
		"unused":       unused,
		"maximum":      maximum,
		"used_percent": usedPercent,
	}, nil
}

func init() {
	inputs.Add("linux_file_descriptors", func() telegraf.Input {
		return &LinuxFileDescriptors{}
	})
}
