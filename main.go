package main

import (
	"fmt"
	"os"

	"my-ls/core"
)

func main() {
	path := "."
	showHidden := false // -a
	sortByTime := false // -t
	reverse := false    // -r
	longFormat := false // -l 

	for _, arg := range os.Args[1:] {
    // Check if it's a flag (starts with '-' and isn't just '-')
    if len(arg) > 1 && arg[0] == '-' {
        // Loop through each character after the '-'
        for i := 1; i < len(arg); i++ {
            switch arg[i] {
            case 'a':
                showHidden = true
            case 'l':
                longFormat = true
            case 't':
                sortByTime = true
            case 'r':
                reverse = true
            default:
                fmt.Printf("my-ls: invalid option -- '%c'\n", arg[i])
                os.Exit(1)
            }
        }
    } else {
        // If it doesn't start with '-', it's our path
        path = arg
    }
}

	names, err := core.DirReader(path, showHidden)
	if err != nil {
		fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
		return
	}

	nodes, err := core.GetFileInfo(names, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
		return
	}

	core.SortNodes(nodes, sortByTime, reverse)

	if longFormat {
		displayLong(nodes)
	} else {
		displayShort(nodes)
	}
}

func displayShort(nodes []core.FileNode) {
	for i, node := range nodes {
		fmt.Print(node.Name)
		if i < len(nodes)-1 {
			fmt.Print("  ") // Two spaces between names
		}
	}
	fmt.Println()
}

func displayLong(nodes []core.FileNode) {
	var totalBlocks int64
	for _, node := range nodes {
		totalBlocks += node.Blocks
	}

	fmt.Printf("total %d\n", totalBlocks/2)

	for _, node := range nodes {
		fmt.Printf("%s %3d %-8s %-8s %8d %s %s\n",
			node.Mode.String(),
			node.Links,
			node.Owner,
			node.Group,
			node.Size,
			node.ModeTime.Format("Jan _2 15:04"),
			node.Name,
		)
	}
}
