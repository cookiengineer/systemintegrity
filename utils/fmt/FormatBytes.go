package fmt

import "fmt"

func FormatBytes(bytes uint64) string {

	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}

	value := float64(bytes)
	unit := 0

	for value >= 1000 && unit < len(units)-1 {
		value /= 1000
		unit++
	}

	if unit == 0 {
		return fmt.Sprintf("%d B", bytes)
	}

	return fmt.Sprintf("%.2f %s", value, units[unit])

}
