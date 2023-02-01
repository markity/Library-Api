package hotreload

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func Go(listenPath string, f func()) {
	println("welcome to Go")
	c := make(chan struct{})

	// 创建一个队列, 内核会传递一个事件对象(包括事件掩码和目标)
	notifyFD, _ := syscall.InotifyInit()

	filepath.WalkDir(listenPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			_, err := syscall.InotifyAddWatch(notifyFD, path, syscall.IN_CREATE)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if strings.HasSuffix(path, ".go") {
				_, err := syscall.InotifyAddWatch(notifyFD, path, syscall.IN_MODIFY)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		return nil
	})
	go f()

	go func() {
		for {
			eventBytes := make([]byte, syscall.SizeofInotifyEvent+syscall.PathMax+1)
			n, _ := syscall.Read(notifyFD, eventBytes)
			event := syscall.InotifyEvent{}
			buf := &bytes.Buffer{}
			toWrite := eventBytes[:n+1]
			buf.Write(toWrite)
			binary.Read(buf, binary.LittleEndian, &event)
			// 如果创建了新文件, 就重新扫描
			if event.Mask == syscall.IN_CREATE {
				syscall.Close(notifyFD)
				notifyFD, _ := syscall.InotifyInit()
				filepath.WalkDir(listenPath, func(path string, d fs.DirEntry, err error) error {
					if d.IsDir() {
						_, err := syscall.InotifyAddWatch(notifyFD, path, syscall.IN_CREATE)
						if err != nil {
							log.Fatal(err)
						}
					} else {
						if strings.HasSuffix(path, ".go") {
							_, err := syscall.InotifyAddWatch(notifyFD, path, syscall.IN_MODIFY)
							if err != nil {
								log.Fatal(err)
							}

						}
					}
					return nil
				})
			} else if event.Mask == syscall.IN_MODIFY {
				go func() {
					c <- struct{}{}
				}()
			}
		}
	}()

	for {
		<-c

		cmd := exec.Command("go", "build", "-o", "/tmp/markity-reload-main.tmp")
		buf := &bytes.Buffer{}
		cmd.Stderr = buf
		err := cmd.Run()
		if err != nil {
			b, _ := io.ReadAll(buf)
			fmt.Println("编译失败:", string(b))
			continue
		}
		syscall.Exec("/tmp/markity-reload-main.tmp", nil, syscall.Environ())
	}
}
