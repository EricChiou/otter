package router

import (
	"log"
	"regexp"
	"strings"
)

func addRoute(method string, tree *node, path string, run func(*Context)) {
	if !checkFormat(method, path) {
		return
	}

	if checkDuplicate(tree, "", path[1:]) {
		log.Fatalln("path duplicated, " + method + ": '" + path)
		return
	}

	if !addPath(tree, "", path[1:], run) {
		log.Fatalln("add path fail, " + method + ": '" + path)
	}
}

func checkFormat(method, path string) bool {
	// check path is valid
	if len(path) == 0 {
		log.Fatalln("path error, path can not be empty.")
		return false
	}

	if path[0:1] != "/" {
		log.Fatalln("path error, " + method + ": '" + path + "', path must begin with '/'.")
		return false
	}

	if len(path) == 1 {
		return true
	}

	paths := strings.Split(path[1:], "/")
	for _, p := range paths {
		if len(p) > 0 && p[0:1] == ":" {
			p = p[1:]
		}
		if len(p) == 0 {
			log.Fatalln("path error, " + method + ": '" + path + "', path has wrong format.")
			return false
		}
		match, _ := regexp.MatchString("^[0-9a-zA-Z]+$", p)
		if !match {
			log.Fatalln("path error, " + method + ": '" + path + "', path has invalid character, only accept 0-9, a-z, A-Z.")
			return false
		}
	}
	return true
}

func checkDuplicate(tree *node, path, pathSeg string) bool {
	if len(pathSeg) == 0 {
		if tree.run != nil {
			return true
		}
	} else {
		path, pathSeg = filterPath(pathSeg)
		wildChild := (path[0:1] == ":")
		if wildChild {
			path = path[1:]
		}

		for _, child := range tree.children {
			if path == child.path || wildChild || child.wildChild {
				if checkDuplicate(child, path, pathSeg) {
					return true
				}
			}
		}
	}

	return false
}

func addPath(tree *node, path, pathSeg string, run func(*Context)) bool {
	if len(pathSeg) == 0 {
		tree.run = &run
		return true
	}

	path, pathSeg = filterPath(pathSeg)
	wildChild := (path[0:1] == ":")
	if wildChild {
		path = path[1:]
	}

	for _, child := range tree.children {
		if path == child.path && wildChild == child.wildChild {
			if addPath(child, path, pathSeg, run) {
				return true
			}
		}
	}

	newChild := node{path: path, wildChild: wildChild, run: nil, children: []*node{}}
	tree.children = append(tree.children, &newChild)
	if addPath(&newChild, path, pathSeg, run) {
		return true
	}

	return false
}

func filterPath(path string) (string, string) {
	for i, c := range path {
		if c == 47 {
			return path[:i], path[i+1:]
		}
	}
	return path, ""
}
