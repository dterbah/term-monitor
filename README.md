# term-monitor

![CI](https://github.com/dterbah/term-monitor/actions/workflows/go-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/dterbah/term-monitor/branch/main/graph/badge.svg)](https://codecov.io/gh/dterbah/term-monitor)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=dterbah_term-monitor&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=dterbah_term-monitor)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=dterbah_term-monitor&metric=bugs)](https://sonarcloud.io/summary/new_code?id=dterbah_term-monitor)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=dterbah_term-monitor&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=dterbah_term-monitor)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_term-monitor&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_term-monitor)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=dterbah_term-monitor&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=dterbah_term-monitor)

## Overview

**term-monitor** is a lightweight command-line tool that provides real-time system monitoring information. It allows users to display various system metrics such as RAM usage, CPU usage, network latency (ping), and the number of running processes.
⚠️ Only tested on MACOS !!

## Features

- 🖥️ **RAM Usage**: Displays the current memory usage.
- ⚙️ **CPU Usage**: Shows the current CPU utilization.
- 🌐 **Network Ping**: Measures the current ping to a predefined destination.
- 📊 **Process Count**: Displays the number of currently running processes.

## Installation

### Prerequisites

- Go 1.24 or later

### Install the CLI

```sh
go install github.com/your-username/term-monitor@latest
```

## Usage

Run the tool with the desired flags:

```sh
./term-monitor [options]
```

### Available Flags

| Flag    | Description                 |
| ------- | --------------------------- |
| `-ram`  | Display RAM usage           |
| `-cpu`  | Display CPU usage           |
| `-ping` | Display network ping        |
| `-proc` | Display number of processes |

Example usage:

```sh
term-monitor -ram -cpu -proc -ping
```

The CLI should display something like this :

![Example of term-monitor in action](https://raw.githubusercontent.com/dterbah/term-monitor/main/docs/screenshot.png)

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Author

Created by **Dorian TERBAH**
