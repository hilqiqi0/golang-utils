package field

import (
	"io/ioutil"

	"github.com/hilqiqi0/golang-utils/tools/errs"
)

// 获取当前路径下所有文件名
func GetAllFiles(pathname string) (names []string) {
	rd, err := ioutil.ReadDir(pathname)
	errs.CheckCommonErr(err)
	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFiles(pathname + fi.Name() + "/")
		} else {
			names = append(names, fi.Name())
		}
	}
	return
}

// 获取当前路径所有文件夹名
func GetPathFolders(pathname string) (names []string) {
	rd, err := ioutil.ReadDir(pathname)
	errs.CheckCommonErr(err)
	for _, fi := range rd {
		if fi.IsDir() {
			names = append(names, fi.Name())
		}
	}
	return
}

// 获取当前路径所有文件夹名
func GetPathFiles(pathname string) (names []string) {
	rd, err := ioutil.ReadDir(pathname)
	errs.CheckCommonErr(err)
	for _, fi := range rd {
		if !fi.IsDir() {
			names = append(names, fi.Name())
		}
	}
	return
}
