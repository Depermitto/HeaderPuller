package hp

import "strings"

func FilepathSplit(filepath string) (dirname string, filename string) {
	i := strings.LastIndex(filepath, PathSep)
	if i == -1 {
		dirname, filename = IncludeDir, filepath
	} else {
		dirname, filename = filepath[:i], filepath[i+1:]
	}
	return dirname, filename
}

func FileFmt(pathParts ...string) (filename string) {
	for i, s := range pathParts {
		filename += s
		if i < len(pathParts)-1 {
			filename += PathSep
		}
	}
	return filename
}

func IsRepoLink(s string) bool {
	return strings.HasPrefix(s, "http")
}
