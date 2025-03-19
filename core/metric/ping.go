package metric

import (
	"os/exec"
	"strconv"
	"strings"
)

const HOST = "google.com"

func GetPing() (float64, error) {
	cmd := exec.Command("ping", "-c", "1", HOST)
	output, err := cmd.Output()

	if err != nil {
		return 0.0, err
	}

	result := string(output)

	lines := strings.Split(result, "\n")
	for _, line := range lines {
		if strings.Contains(line, "time=") {
			contents := strings.Split(line, " ")
			for _, content := range contents {
				if strings.Contains(content, "time=") {
					ping := strings.Split(content, "=")[1]
					return strconv.ParseFloat(ping, 64)
				}
			}
		}
	}

	return 0, err
}
