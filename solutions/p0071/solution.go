package solution

import "path"

func simplifyPath(p string) string {
	return path.Clean(p)
}
