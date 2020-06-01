package gen

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
)

const (
	Total = 37 * 37 * 37
	//Total = 37 * 37 * 37 * 37 * 37
)

func Username() {
	initDir()

	// gen username with file.
	fs := afero.NewOsFs()
	currentFile := "./1/all"
	var err error
	var file afero.File

	// first file
	file, err = fs.Create(currentFile)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < Total; i++ {
		fileName, username := genUsernameAndFile(i)
		if fileName == "" {
			continue
		}
		if fileName != currentFile {
			_ = file.Close()
			currentFile = fileName
			exist, err := afero.Exists(fs, fileName)
			if err != nil {
				panic(err.Error())
			}
			if exist {
				file, err = fs.Open(fileName)
				if err != nil {
					panic(fmt.Errorf("open file %s fail.err: %v", fileName, err.Error()))
				}
			} else {
				file, err = fs.Create(fileName)
			}
		}
		_, _ = file.WriteString(username + "\n")
	}
	_ = file.Close()
}

// switchFile
// 0-9 equal 0-9
// 10-35 equal a-z
// 36 equal -
func genUsernameAndFile(file int) (fileName, username string) {
	for {
		letter := '-'
		rem := file % 37
		switch {
		case 0 <= rem && rem <= 9:
			letter = int32(rem + 48)
		case 10 <= rem && rem <= 35:
			letter = int32(rem + 97 - 10)
		default:
			letter = '-'
		}

		username = fmt.Sprintf("%s%s", string(letter), username)
		file /= 37
		if file <= 0 {
			break
		}
	}

	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return "", ""
	}

	switch len(username) {
	case 1:
		fileName = "./1/all"
	case 2:
		fileName = "./2/all"
	case 3:
		fileName = "./3/all"
	case 4:
		fileName = fmt.Sprintf("./4/%s", username[:1])
	default:
		return "", ""
	}
	return
}

func Clean() {
	// clean all of dir and files.
	fs := afero.NewOsFs()
	_ = fs.RemoveAll("./1")
	_ = fs.RemoveAll("./2")
	_ = fs.RemoveAll("./3")
	_ = fs.RemoveAll("./4")
}

func initDir() {
	fs := afero.NewOsFs()
	_ = fs.RemoveAll("./1")
	_ = fs.RemoveAll("./2")
	_ = fs.RemoveAll("./3")
	_ = fs.RemoveAll("./4")
	_ = fs.Mkdir("./1", os.ModeDir)
	_ = fs.Mkdir("./2", os.ModeDir)
	_ = fs.Mkdir("./3", os.ModeDir)
	_ = fs.Mkdir("./4", os.ModeDir)
}
