package gen

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/afero"
)

var (
	fs        = afero.NewOsFs()
	numberMap sync.Map
)

func Username() {
	initDir()

	gen1()
	gen2()
	gen3()
	gen4()
}

// switchFile
// 0 equal -
// 1-10 equal 0-9
// 11-36 equal a-z
func number2Letter(nums ...int32) (name string) {
	for _, v := range nums {
		value, _ := numberMap.Load(v)
		name = fmt.Sprintf("%s%s", name, string(value.(int32)))
	}
	if !validUsername(name) {
		return ""
	}
	return
}

func gen1() {
	f1, err := fs.Create("./1/all")
	if err != nil {
		panic(err.Error())
	}
	var i int32
	for i = 0; i < 37; i++ {
		s := number2Letter(i)
		if s == "" {
			continue
		}
		_, _ = f1.WriteString(s + "\n")
	}
}

func gen2() {
	f1, err := fs.Create("./2/all")
	if err != nil {
		panic(err.Error())
	}
	var i, j int32
	for i = 0; i < 37; i++ {
		for j = 0; j < 37; j++ {
			s := number2Letter(i, j)
			if s == "" {
				continue
			}
			_, _ = f1.WriteString(s + "\n")
		}
	}
}

func gen3() {
	f1, err := fs.Create("./3/all")
	if err != nil {
		panic(err.Error())
	}
	var i, j, k int32
	for i = 0; i < 37; i++ {
		for j = 0; j < 37; j++ {
			for k = 0; k < 37; k++ {
				s := number2Letter(i, j, k)
				if s == "" {
					continue
				}
				_, _ = f1.WriteString(s + "\n")
			}
		}
	}
}

func gen4() {
	wg := sync.WaitGroup{}
	var i int32
	for i = 1; i < 37; i++ {
		wg.Add(1)
		go func(i int32) {
			value, _ := numberMap.Load(i)
			f4, err := fs.Create(fmt.Sprintf("./4/%s.txt", string(value.(int32))))
			if err != nil {
				panic(err.Error())
			}
			count := 0
			var j, k, l int32
			for j = 0; j < 37; j++ {
				for k = 0; k < 37; k++ {
					for l = 0; l < 37; l++ {
						count++
						s := number2Letter(i, j, k, l)
						if s == "" {
							continue
						}
						_, err = f4.WriteString(s + "\n")
						if err != nil {
							panic(err.Error())
						}
					}
				}
			}
			fmt.Printf("file: %v,count: %v\n", i, count)
			_ = f4.Close()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func validUsername(username string) bool {
	if len(username) < 1 {
		return false
	}
	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return false
	}
	return true
}

func Clean() {
	// clean all of dir and files.
	_ = fs.RemoveAll("./1")
	_ = fs.RemoveAll("./2")
	_ = fs.RemoveAll("./3")
	_ = fs.RemoveAll("./4")
}

func initDir() {
	_ = fs.RemoveAll("./1")
	_ = fs.RemoveAll("./2")
	_ = fs.RemoveAll("./3")
	_ = fs.RemoveAll("./4")
	_ = fs.Mkdir("./1", os.ModeDir)
	_ = fs.Mkdir("./2", os.ModeDir)
	_ = fs.Mkdir("./3", os.ModeDir)
	_ = fs.Mkdir("./4", os.ModeDir)

	// switchFile
	// 0 equal -
	// 1-10 equal 0-9
	// 11-36 equal a-z
	var rem int32
	for rem = 0; rem < 37; rem++ {
		switch {
		case 0 == rem:
			numberMap.Store(rem, '-')
		case 1 <= rem && rem <= 10:
			numberMap.Store(rem, rem+48-1)
		default:
			numberMap.Store(rem, rem+97-11)
		}
	}
}
