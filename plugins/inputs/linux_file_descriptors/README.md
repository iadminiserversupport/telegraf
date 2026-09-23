

# Linux File Descriptors Input Plugin

The `linux_file_descriptors` input plugin collects system-wide Linux file descriptor statistics from `/proc/sys/fs/file-nr`.

### Features

The plugin reports:

- `allocated` — allocated file descriptors
- `unused` — unused file descriptors
- `maximum` — system-wide maximum file descriptors
- `used_percent` — percentage of the maximum currently allocated

### Configuration

```toml
[[inputs.linux_file_descriptors]]
  ## Path to Linux file descriptor statistics.
  ## Default: /proc/sys/fs/file-nr
  # path = "/proc/sys/fs/file-nr"
````

### Example output

```text
linux_file_descriptors allocated=13223i,unused=0i,maximum=9223372036854775807i,used_percent=0.000000
```

The plugin is intended for Linux systems where system-wide file descriptor usage needs to be monitored.

For Linux and cloud server management:

[https://iserversupport.com/cloud-server-management/](https://iserversupport.com/cloud-server-management/)
