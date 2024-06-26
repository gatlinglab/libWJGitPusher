package wjGitPusher

import (
	"fmt"
	"testing"
)

const g_key = ``

func TestWJGP_NewGitPusher(t *testing.T) {
	err := WJGP_Initialize("")
	if err != nil {
		t.Errorf("git provider init error: " + err.Error())
		return
	}

	pusher, err := WJGP_AddRepository("test1", "git@github.com:aristotle88/test1.git")
	if err != nil {
		t.Errorf("git add repository failed: " + err.Error())
		return
	}

	err = pusher.Initialize(g_key)
	if err != nil {
		t.Errorf("repository init failed: " + err.Error())
		return
	}

	err = pusher.AddNewFileData("test1", []byte("test1 data version 2"))
	if err != nil {
		t.Errorf("add new file error: " + err.Error())
		return
	}
	err = pusher.CommitAndPush("admin", "admin@google.com", "modify file version 2")
	if err != nil {
		t.Errorf("commit and push failed: " + err.Error())
		return
	}

	filelist := pusher.ListTree()

	fmt.Println("start file list: ")
	for _, str1 := range filelist {
		fmt.Println(str1)
	}

}
