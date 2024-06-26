package iGitPusher

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CWjGitProvider struct {
	rootPath string
}

var G_singleGitProvider *CWjGitProvider = &CWjGitProvider{}

func IGP_GetGitProvider() *CWjGitProvider {
	return G_singleGitProvider
}

func (pInst *CWjGitProvider) Initialize(rootPath string) error {
	err := pInst.initPath(rootPath)
	if err != nil {
		return err
	}
	err = pInst.initKnowHost()
	if err != nil {
		return err
	}

	return nil
}

func (pInst *CWjGitProvider) AddRepository(reposName, remoteUrl string) (*CWJGitPusher, error) {
	provider := pInst.analyseProvider(remoteUrl)
	if provider == "" {
		return nil, errors.New("don't know the git provider")
	}
	localPath := filepath.Join(pInst.rootPath, reposName)
	err := os.MkdirAll(localPath, 0755)
	if err != nil {
		return nil, err
	}

	return newGitPusher(provider, reposName, localPath, remoteUrl), nil
}

func (pInst *CWjGitProvider) initPath(path string) error {
	if path != "" {
		pInst.rootPath = path
		return nil
	}
	file, _ := exec.LookPath(os.Args[0])
	path1, _ := filepath.Abs(file)
	index := strings.LastIndex(path1, string(os.PathSeparator))
	pInst.rootPath = filepath.Join(path1[:index+1], "gitroot")
	err := os.MkdirAll(pInst.rootPath, 0755)

	return err
}

func (pInst *CWjGitProvider) initKnowHost() error {
	knowhostFile := filepath.Join(pInst.rootPath, "known_hosts")
	err := os.WriteFile(knowhostFile, []byte(CG_KnowHost), 0644)
	if err != nil {
		fmt.Println("init know host file error: ", err)
		return err
	}
	os.Setenv("SSH_KNOWN_HOSTS", knowhostFile)

	return nil
}

func (pInst *CWjGitProvider) analyseProvider(url string) string {
	index := strings.Index(url, "github.com")
	if index > -1 {
		return "GITHUB"
	}
	index = strings.Index(url, "gitlab.com")
	if index > -1 {
		return "GITLAB"
	}
	index = strings.Index(url, "bitbucket.org")
	if index > -1 {
		return "BITBUCKET"
	}

	return ""
}
