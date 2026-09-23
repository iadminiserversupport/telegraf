//go:build !custom || inputs || inputs.linux_file_descriptors

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/linux_file_descriptors" // register plugin
