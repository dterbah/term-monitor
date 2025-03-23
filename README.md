# term-monitor

## Overview

**term-monitor** is a lightweight command-line tool that provides real-time system monitoring information. It allows users to display various system metrics such as RAM usage, CPU usage, network latency (ping), and the number of running processes.

## Features

- 🖥️ **RAM Usage**: Displays the current memory usage.
- ⚙️ **CPU Usage**: Shows the current CPU utilization.
- 🌐 **Network Ping**: Measures the current ping to a predefined destination.
- 📊 **Process Count**: Displays the number of currently running processes.

## Installation

### Prerequisites

- Go 1.18 or later

### Clone the repository

```sh
git clone https://github.com/your-username/term-monitor.git
cd term-monitor
```

### Build the project

```sh
go build -o term-monitor
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
./term-monitor -ram -cpu
```

This command will display both RAM and CPU usage.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Author

Created by **Your Name**

