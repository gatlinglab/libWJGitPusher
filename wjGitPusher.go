package wjGitPusher

import iGitPusher "github.com/gatlinglab/libWJGitPusher/internel"

//	type IWJGitProvider interface {
//		Initialize(rootpath string) error
//		AddRepository(reposName, localPath, remoteUrl, key string) IWJGitPusher
//	}
type IWJGitPusher interface {
	Initialize(key string) error
	AddNewFileData(filename string, data []byte) error
	AddNewFile(filename string) error
	CommitAndPush(authorName, authorEmail, comment string) error
	DeleteLocalFile(filename string)
	ListTree() []string
}

func WJGP_Initialize(rootPath string) error {
	return iGitPusher.G_singleGitProvider.Initialize(rootPath)
}

func WJGP_AddRepository(reposName, remoteUrl string) (IWJGitPusher, error) {
	return iGitPusher.G_singleGitProvider.AddRepository(reposName, remoteUrl)
}
