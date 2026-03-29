package display

import (
	"fmt"

	"my-ls/core"
)

func PrintShortFormat(nodes []core.FileNode) {
	for _, node := range nodes {
		fmt.Printf("%s ", node.Name)
	}
	fmt.Println()
}

func PrintLongFormat(nodes []core.FileNode) {
	var totalBlocks int64
	for _, node := range nodes {
		totalBlocks += node.Blocks
	}
	fmt.Printf("total %d\n", totalBlocks)

	for _, node := range nodes {
		fmt.Printf("%s %d %d %d %8d %s %s\n",
			node.Mode.String(),
			node.Links,
			node.UID,
			node.GID,
			node.Size,
			node.ModeTime.Format("Jan _2 15:04"),
			node.Name,
		)
	}
}
