package iGitPusher

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func (pInst *CWJGitPusher) cloneRepository() error {
	r, err := git.PlainClone(pInst.localPath, false, &git.CloneOptions{
		Auth:     pInst.publicKey,
		URL:      pInst.reposUrl,
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		fmt.Println("PlainClone error: ")
		return err
	}
	pInst.gitRepos = r
	return nil
}
func (pInst *CWJGitPusher) openRepository() error {
	r, err := git.PlainOpen(pInst.localPath)
	if err != nil {
		fmt.Println("git.PlainOpen error: ")
		return err
	}
	pInst.gitRepos = r
	return nil
}

func (pInst *CWJGitPusher) initReposTree() error {
	w, err := pInst.gitRepos.Worktree()
	if err != nil {
		fmt.Println("Worktree error: ", err)
		return err
	}

	pInst.gitTree = w

	return nil
}
