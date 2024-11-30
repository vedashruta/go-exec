package cpu

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MemoryStats struct {
	Total     string `json:"total"`
	Used      string `json:"used"`
	Free      string `json:"free"`
	Shared    string `json:"shared"`
	BuffCache string `json:"buff_cache"`
	Available string `json:"available"`
}

type SwapStats struct {
	Total string `json:"total"`
	Used  string `json:"used"`
	Free  string `json:"free"`
}

type CPUStats struct {
	Usage       string `json:"usage"`
	UserTime    string `json:"user_time"`
	SystemTime  string `json:"system_time"`
	IdleTime    string `json:"idle_time"`
	LoadAverage string `json:"load_average"`
}

type ProcessStats struct {
	TotalProcesses int `json:"total_processes"`
	Running        int `json:"running"`
	Sleeping       int `json:"sleeping"`
	Stopped        int `json:"stopped"`
	Zombie         int `json:"zombie"`
}

func memory(c *fiber.Ctx) (err error) {
	command := exec.Command("free", "-h")
	stdout, err := command.Output()
	if err != nil {
		return c.JSON(err)
	}
	lines := bytes.Split(stdout, []byte("\n"))
	if len(lines) < 3 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unexpected output format",
		})
	}
	memoryFields := strings.Fields(string(lines[1]))
	if len(memoryFields) >= 7 {
		memoryStats := MemoryStats{
			Total:     memoryFields[1],
			Used:      memoryFields[2],
			Free:      memoryFields[3],
			Shared:    memoryFields[4],
			BuffCache: memoryFields[5],
			Available: memoryFields[6],
		}
		swapFields := strings.Fields(string(lines[2]))
		var swapStats SwapStats
		if len(swapFields) >= 4 {
			swapStats = SwapStats{
				Total: swapFields[1],
				Used:  swapFields[2],
				Free:  swapFields[3],
			}
		}
		data := fiber.Map{
			"time":          time.Now().Format(time.RFC3339),
			"ram_available": memoryStats,
			"swap":          swapStats,
		}
		return c.JSON(data)
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "failed to parse memory stats",
	})
}

func system(c *fiber.Ctx) error {
	cpuCommand := exec.Command("top", "-bn1")
	cpuOutput, err := cpuCommand.Output()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to execute top command",
		})
	}
	psCommand := exec.Command("ps", "-e", "-o", "state")
	psOutput, err := psCommand.Output()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to execute ps command",
		})
	}
	lines := strings.Split(string(cpuOutput), "\n")
	var cpuStats CPUStats
	for _, line := range lines {
		if strings.Contains(line, "%Cpu(s):") {
			fields := strings.Fields(line)
			if len(fields) >= 8 {
				cpuStats = CPUStats{
					Usage:       fields[1] + "%",
					UserTime:    fields[1] + "%",
					SystemTime:  fields[3] + "%",
					IdleTime:    fields[7] + "%",
					LoadAverage: lines[len(lines)-1], // last line in top contains load average
				}
			}
			break
		}
	}
	var processStats ProcessStats
	processStates := strings.Fields(string(psOutput))
	stateCount := make(map[string]int)
	for _, state := range processStates {
		stateCount[state]++
	}
	processStats = ProcessStats{
		TotalProcesses: len(processStates),
		Running:        stateCount["R"],
		Sleeping:       stateCount["S"],
		Stopped:        stateCount["T"],
		Zombie:         stateCount["Z"],
	}
	data := fiber.Map{
		"time":          time.Now().Format(time.RFC3339),
		"cpu_stats":     cpuStats,
		"process_stats": processStats,
	}
	return c.JSON(data)
}
