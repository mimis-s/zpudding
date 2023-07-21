package file

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

func DirCopy(src string, dest string) error {
	os.Mkdir(dest, os.ModePerm)
	// 遍历原文件夹内部所有item
	items, _ := ioutil.ReadDir(src)
	for _, item := range items {

		// 文件
		if !item.IsDir() {
			FileCopy(path.Join(src, item.Name()), path.Join(dest, item.Name()))
			continue
		}

		// // 目录
		// os.Mkdir(path.Join(dest, item.Name()), os.ModePerm)
		// 递归
		DirCopy(path.Join(src, item.Name()), path.Join(dest, item.Name()))
	}

	return nil
}

func FileCopy(src, dest string) error {
	// open src readonly
	srcFp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	// create dest
	dstFp, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dstFp.Close()

	// copy
	_, err = io.Copy(dstFp, srcFp)
	return err
}
