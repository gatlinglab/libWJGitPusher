package iGitPusher

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type CWJGitPusher struct {
	reposName string
	reposUrl  string
	localPath string
	//reposKey    string
	gitProvider string
	publicKey   *ssh.PublicKeys
	gitRepos    *git.Repository
	gitTree     *git.Worktree
}

func newGitPusher(provider, reposName, localPath, remoteUrl string) *CWJGitPusher {
	return &CWJGitPusher{gitProvider: provider, reposName: reposName, localPath: localPath, reposUrl: remoteUrl}
}

func (pInst *CWJGitPusher) Initialize(key string) error {
	if pInst.reposName == "" {
		return fmt.Errorf("git repository name is empty")
	}

	publicKey, err := ssh.NewPublicKeys("git", []byte(key), "")
	if err != nil {
		fmt.Println("generate public key error: ")
		return err
	}

	pInst.publicKey = publicKey

	// check the .git path
	checkPath := filepath.Join(pInst.localPath, ".git")
	// clone bare only. then open local git;
	if _, err1 := os.Stat(checkPath); os.IsNotExist(err1) {
		err = pInst.cloneRepository()
		if err != nil {
			fmt.Println("clone repositoryBare error: ", err)
			return err
		}
	} else {
		err = pInst.openRepository()
		if err != nil {
			fmt.Println("open local repository error: ", err)
			return err
		}

	}

	err = pInst.initReposTree()
	if err != nil {
		return err
	}

	return nil
}
func (pInst *CWJGitPusher) AddNewFileData(filename string, data []byte) error {
	file1 := filepath.Join(pInst.localPath, filename)
	err := os.WriteFile(file1, data, 0644)
	if err != nil {
		fmt.Println("WriteFile error where add file to repository: ", err)
		return err
	}
	_, err = pInst.gitTree.Add(filename)
	if err != nil {
		fmt.Println("git tree add file error: ", err)
		return err
	}
	_, err = pInst.gitTree.Status()
	if err != nil {
		fmt.Println("git tree status error: ", err)
		return err
	}
	return nil
}
func (pInst *CWJGitPusher) AddNewFile(filename string) error {
	_, err := pInst.gitTree.Add(filename)
	if err != nil {
		fmt.Println("git tree add file error: ", err)
		return err
	}
	_, err = pInst.gitTree.Status()
	if err != nil {
		fmt.Println("git tree status error: ", err)
		return err
	}
	return nil
}
func (pInst *CWJGitPusher) CommitAndPush(authorName, authorEmail, comment string) error {
	commit, err := pInst.gitTree.Commit(comment, &git.CommitOptions{
		Author: &object.Signature{
			Name:  authorName,
			Email: authorEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println("commit generate error: ", err)
		return err
	}

	_, err = pInst.gitRepos.CommitObject(commit)
	if err != nil {
		fmt.Println("commit to repository error: ", err)
		return err
	}

	err = pInst.gitRepos.Push(&git.PushOptions{
		Auth: pInst.publicKey,
	})
	if err != nil {
		fmt.Println("Push error: ", err)
		return err
	}

	return nil
}
func (pInst *CWJGitPusher) DeleteLocalFile(filename string) {
	file1 := filepath.Join(pInst.localPath, filename)
	os.Remove(file1)
}
func (pInst *CWJGitPusher) ListTree() []string {
	var files []string
	ref, err := pInst.gitRepos.Head()
	if err != nil {
		fmt.Println("git head error: ", err)
		return files
	}
	commit, err := pInst.gitRepos.CommitObject(ref.Hash())
	if err != nil {
		fmt.Println("git get commit error: ", err)
		return files
	}
	tree, err := commit.Tree()
	if err != nil {
		fmt.Println("git get tree error: ", err)
		return files
	}
	tree.Files().ForEach(func(f *object.File) error {
		files = append(files, f.Name)
		return nil
	})

	return files

}
