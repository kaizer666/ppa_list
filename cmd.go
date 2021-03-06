package main

import (
	"os/exec"
	"bytes"
	"os"
	"strings"
	"strconv"
)

func ExecCmd(name string, cmds []string) (string,error){
	cmd := exec.Command(name, cmds...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	var b2 bytes.Buffer=bytes.Buffer{}

	b2.ReadFrom(stdout)
	return string(b2.Bytes()),nil
}

type AppInfo struct {
	Pid uint32
	Path string
}

func GetPid() []AppInfo {
	thisApp:=os.Args[0]
	thisPid:=uint32(os.Getpid())
	thisAppLen:=len(thisApp)
	s,_:=ExecCmd("ps",[]string{"ax","-o", "pid,command"})
	a:=strings.Split(s,"\n")
	result:=make([]AppInfo,0)
	for _,p:=range a[1:]{
		p=strings.Trim(p," ")
		indx:=strings.Index(p," ")
		if p=="" || indx<0{
			continue
		}
		pid, _ := strconv.ParseUint(p[0:indx], 10, 32)
		pch:=string(strings.Trim(p[indx:]," "))
		if len(pch)>=thisAppLen && pch[0:thisAppLen]==thisApp && thisPid!=uint32(pid){
			result=append(result,AppInfo{Pid:uint32(pid),Path:pch})
		}
	}
	return result
}
