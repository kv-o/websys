package filesys

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"main/auth"
)

const (
	O_READ     = 4
	O_WRITE    = 2
	O_EXEC     = 1
	O_USRSHIFT = 6
	O_GRPSHIFT = 3
	O_OTHSHIFT = 0

	O_USR_R = O_READ << O_USRSHIFT
	O_USR_W = O_WRITE << O_USRSHIFT
	O_USR_X = O_EXEC << O_USRSHIFT

	O_GRP_R = O_READ << O_GRPSHIFT
	O_GRP_W = O_WRITE << O_GRPSHIFT
	O_GRP_X = O_EXEC << O_GRPSHIFT

	O_OTH_R = O_READ << O_OTHSHIFT
	O_OTH_W = O_WRITE << O_OTHSHIFT
	O_OTH_X = O_EXEC << O_OTHSHIFT

	EXABYTE  = 1000000000000000000
	EXBIBYTE = 1152921504606846976
	GIBIBYTE = 1073741824
	GIGABYTE = 1000000000
	KIBIBYTE = 1024
	KILOBYTE = 1000
	MEBIBYTE = 1048576
	MEGABYTE = 1000000
	PEBIBYTE = 1125899906842624
	PETABYTE = 1000000000000000
	TEBIBYTE = 1099511627776
	TERABYTE = 1000000000000
)

func permissions(mode fs.FileMode) string {
	xbit := byte('x')
	perms := []byte("--- --- ---")

	if mode.IsDir() {
		xbit = byte('l')
	}
	mode = mode.Perm()

	if mode&O_USR_R != 0 {
		perms[0] = byte('r')
	}
	if mode&O_USR_W != 0 {
		perms[1] = byte('w')
	}
	if mode&O_USR_X != 0 {
		perms[2] = xbit
	}

	if mode&O_GRP_R != 0 {
		perms[4] = byte('r')
	}
	if mode&O_GRP_W != 0 {
		perms[5] = byte('w')
	}
	if mode&O_GRP_X != 0 {
		perms[6] = xbit
	}

	if mode&O_OTH_R != 0 {
		perms[8] = byte('r')
	}
	if mode&O_OTH_W != 0 {
		perms[9] = byte('w')
	}
	if mode&O_OTH_X != 0 {
		perms[10] = xbit
	}

	return string(perms)
}

// BUG: errors are returned immediately, disallowing 16.4k+ notation
func size(path string, info fs.FileInfo) (int64, error) {
	var sz int64
	if !info.IsDir() {
		return info.Size(), nil
	}
	dir, err := os.ReadDir(path)
	if err != nil {
		return sz, err
	}
	for _, node := range dir {
		name := filepath.Join(path, node.Name())
		attrs, err := node.Info()
		if err != nil {
			return sz, err
		}
		nodeSize, err := size(name, attrs)
		if err != nil {
			return sz, err
		}
		sz += nodeSize
	}
	return sz, nil
}

type FileInfo struct {
	Perms   string
	Owner   string
	Group   string
	Size    string
	ModTime string
	Name    string
	Next    *FileInfo
}

func readDir(w io.Writer, path string) error {
	var maxOwner, maxGroup int
	var infos FileInfo
	info := &infos

	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range dir {
		node, err := entry.Info()
		if err != nil {
			return err
		}

		// permissions
		info.Perms = permissions(node.Mode())

		// owner
		// BUG: Go does not provide node.Owner
		info.Owner = "none"
		if len(info.Owner) > maxOwner {
			maxOwner = len(info.Owner)
		}

		// group
		// BUG: Go does not provide node.Group
		info.Group = "none"
		if len(info.Group) > maxGroup {
			maxGroup = len(info.Group)
		}

		// size
		sz, err := size(filepath.Join(path, node.Name()), node)
		if err != nil {
			return err
		}
		if sz < KILOBYTE {
			info.Size = fmt.Sprintf("%d", sz)
		} else if sz > EXABYTE {
			info.Size = fmt.Sprintf("%.1fE", float64(sz)/EXABYTE)
		} else if sz > PETABYTE {
			info.Size = fmt.Sprintf("%.1fP", float64(sz)/PETABYTE)
		} else if sz > TERABYTE {
			info.Size = fmt.Sprintf("%.1fT", float64(sz)/TERABYTE)
		} else if sz > GIGABYTE {
			info.Size = fmt.Sprintf("%.1fG", float64(sz)/GIGABYTE)
		} else if sz > MEGABYTE {
			info.Size = fmt.Sprintf("%.1fM", float64(sz)/MEGABYTE)
		} else if sz > KILOBYTE {
			info.Size = fmt.Sprintf("%.1fk", float64(sz)/KILOBYTE)
		}

		// modtime
		info.ModTime = node.ModTime().Format("2006/01/02 15:04")

		// name
		info.Name = node.Name()

		newInfo := new(FileInfo)
		info.Next = newInfo
		info = info.Next
	}

	for info = &infos; info.Next != nil; info = info.Next {
		ownerPad := maxOwner - len(info.Owner)
		groupPad := maxGroup - len(info.Group)
		sizePad := 6 - len(info.Size)

		fmt.Fprintf(
			w,
			"%s    %s    %s%s    %s%s    %s    %s\n",
			info.Perms,
			info.Owner,
			strings.Repeat(" ", ownerPad),
			info.Group,
			strings.Repeat(" ", groupPad+sizePad),
			info.Size,
			info.ModTime,
			info.Name,
		)
	}

	return nil
}

func Asset(w http.ResponseWriter, r *http.Request) {
	_, err := auth.LookupUser(r.Header.Get("Cookie"))
	if err != nil {
		w.Header().Set("Location", "/signin?redirect=" + url.QueryEscape(r.URL.String()))
		w.WriteHeader(302)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/fs")
	node, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}
	if node.IsDir() {
		err = readDir(w, path)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return
		}
		io.Copy(w, file)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
