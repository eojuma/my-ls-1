package core

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

type FileNode struct {
	Name     string
	Path     string
	Mode     os.FileMode
	ModeTime time.Time
	IsDir    bool
	// long listing
	Links  uint64
	Size   int64
	UID    uint32
	GID    uint32
	Blocks int64
	Owner  string
	Group  string
}

func GetFileInfo(names []string, parentPath string) ([]FileNode, error) {
	nodes := make([]FileNode, 0, len(names))
	for _, name := range names {

		fullPath := parentPath + "/" + name

		if parentPath == "." || parentPath == "" {
			fullPath = name
		}

		info, err := os.Lstat(fullPath)
		if err != nil {
			continue
		}

		stat := info.Sys().(*syscall.Stat_t)

		uidStr := strconv.FormatUint(uint64(stat.Uid), 10)
		gidStr := strconv.FormatUint(uint64(stat.Gid), 10)

		ownerName := uidStr
		u, err := user.LookupId(uidStr)
		if err == nil {
			ownerName = u.Username
		}

		groupName := gidStr 
		g, err := user.LookupGroupId(gidStr)
		if err == nil {
			groupName = g.Name
		}

		


		node := FileNode{
			Name:     name,
			Path:     fullPath,
			Size:     info.Size(),
			Mode:     info.Mode(),
			ModeTime: info.ModTime(),
			IsDir:    info.IsDir(),
			Links:    uint64(stat.Nlink),
			UID:      stat.Uid,
			GID:      stat.Gid,
			Blocks:   stat.Blocks,
		}
		node.Owner = ownerName
		node.Group = groupName
		nodes = append(nodes, node)
	}

	return nodes, nil
}
