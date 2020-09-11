// Copyright (c) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
        "fmt"
        //"io/ioutil"
        //"os/exec"
	"os"
        //"path/filepath"
        //"reflect"
        //"strings"
	//"strconv"
        "testing"

        "github.com/stretchr/testify/assert"
	//"github.com/prometheus/procfs"
	//"github.com/pkg/errors"
)

func TestNewProc(t *testing.T) {
	assert := assert.New(t)

	pid := os.Getpid()
	p, err := NewProc(pid)
	assert.NotNil(p)
	assert.NoError(err)

	p, err = NewProc(0)
	assert.Nil(p)
	assert.Error(err)

	//cmd1 := exec.Command("/bin/sleep","30")
	//cmd1.Start()
	//cmd2 := exec.Command("/bin/sleep", "30")
	//cmd2.Start()
	//cmd3 := exec.Command("/bin/sleep","30")
	//cmd3.Start()
	//cmd2.Process.Kill()
	//cmd.Process.Release()
	//cmd, _ := os.StartProcess("/bin/sleep", 10, "")
	/*pid = os.Getpid()
	ppid := os.Getppid()
	fmt.Println("pid: ", pid)
	fmt.Println("ppid:", ppid)
	cmd1 := exec.Command("/bin/sleep","300")
	cmd1.Run()
	pid = os.Getpid()
        ppid = os.Getppid()
        fmt.Println("pid: ", pid)
        fmt.Println("ppid:", ppid)
*/
}

func TestChildren(t *testing.T) {
	assert := assert.New(t)

	pid := os.Getpid()
	ppid := os.Getppid()

	parent, err := NewProc(ppid)
	children, err := parent.Children()
	assert.NoError(err)
	assert.NotNil(children)

	//c := strconv.Itoa(child.PID)
	//c1 := strconv.Itoa(children[0].PID)
	//fmt.Println(c1)
	//fmt.Println(c)
	//fmt.Println(uint(pid))

	var childrenPIDs []int
	for _, child := range children {
		childrenPIDs = append(childrenPIDs, child.PID)
	}
	fmt.Println(childrenPIDs)
	fmt.Println(pid)
	fmt.Println(ppid)
	//assert.Contains(childrenPIDs, pid)
}

//func TestChildren(t * testing.T) {
//	assert := assert.New(t)

/*	p, err := NewProc(1)
	assert.NoError(err)
	parent := strconv.Itoa(p.PID)
	fmt.Println(parent)
	infos, err := ioutil.ReadDir(filepath.Join(procfs.DefaultMountPoint, parent, taskPath))
	fmt.Println(infos[0])
	fmt.Println(err)
	l := len(infos)
	fmt.Println("length", l)
	if !infos[0].IsDir() {
		fmt.Println("Infos[0] is not dir")
	}
	if infos[0].Name() == parent {
		fmt.Println("Infos[0] name is the same as parent")
	}
	pid, err := strconv.Atoi(infos[0].Name())
	assert.NoError(err)
	fmt.Println("PID: ", pid)
	child, err := NewProc(pid)
	assert.NoError(err)
	fmt.Println("child", child)
	var children []*Proc
	children = append(children, child)
	fmt.Println(children)*/
	//children, err := p.Children()
	//assert.Nil(children)
	//fmt.Println(children)
	//fmt.Println(err)
//	p, _ := NewProc(1)
//	children, _ := p.testChildren()
//	assert.NotNil(children)
//}
/*
func (p *Proc) testChildren() ([]*Proc, error) {
        parent := strconv.Itoa(p.PID)
	fmt.Println("parent", parent)
	fmt.Println("taskPath", taskPath)
        infos, err := ioutil.ReadDir(filepath.Join(procfs.DefaultMountPoint, parent, taskPath))
        if err != nil {
                return nil, errors.Wrapf(err, "Fail to read pid %v proc task dir", p.PID)
        }

        var children []*Proc
        for _, info := range infos {
                //if !info.IsDir() || info.Name() == parent {
                if !info.IsDir() {
			fmt.Println("not dir or is parent")
                        continue
                }
		fmt.Println("before strconv for pid")
                pid, err := strconv.Atoi(info.Name())
		fmt.Println("pid", pid)
                if err != nil {
			fmt.Println("error is not nil")
                        return nil, errors.Wrapf(err, "Invalid child pid %v", info.Name())
                }
		fmt.Println("before NewProc")
                child, err := NewProc(pid)
		fmt.Println("child", child)
                if err != nil {
                        return nil, err
                }
                children = append(children, child)
        }

        return children, nil
}*/
