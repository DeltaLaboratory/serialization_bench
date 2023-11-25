package serialization_bench

import "fmt"

func formatByteSize(size int) string {
	table := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB"}
	i := 0
	for size > 1024 {
		size /= 1024
		i++
	}
	return fmt.Sprintf("%d%s", size, table[i])
}
