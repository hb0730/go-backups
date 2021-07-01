package uploads

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	util "github.com/hb0730/go-backups/utils"
	"github.com/mritd/logger"
	"os"
	"time"
)

//GitUpload git upload
type GitUpload struct {
	//URL git url
	URL string
	//Username username
	Username string
	//Email git email
	Email string
	//Token git token
	Token string
	//DirPath git repository local path
	DirPath    string
	repository *git.Repository
	worktree   *git.Worktree
}

//NewGitUpload new GitUpload
func NewGitUpload(url, username, email, token, dirPath string) *GitUpload {
	g := new(GitUpload)
	g.URL = url
	g.Username = username
	g.Email = email
	g.Token = token
	g.DirPath = dirPath
	return g
}

//Close remove git repository
func (g *GitUpload) Close() error {
	return os.RemoveAll(g.DirPath)
}

func (g *GitUpload) auth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: g.Username,
		Password: g.Token,
	}
}
func (g *GitUpload) add(path string) (err error) {
	return g.worktree.AddWithOptions(&git.AddOptions{All: true, Path: path})
}
func (g *GitUpload) commit(description string) (err error) {
	_, err = g.worktree.Commit(
		description,
		&git.CommitOptions{
			All:    true,
			Author: &object.Signature{Name: g.Username, Email: g.Email, When: time.Now()},
		},
	)
	return
}
func (g *GitUpload) push() error {
	return g.repository.Push(&git.PushOptions{Auth: g.auth()})
}

func (g *GitUpload) commitAndPush(description string) error {
	err := g.commit(description)
	if err != nil {
		logger.Error("[git]", "git commit error", err.Error())
		return err
	}
	return g.push()
}

//Clone git clone
func (g *GitUpload) Clone() (err error) {
	g.repository, err = git.PlainClone(g.DirPath, false, &git.CloneOptions{
		URL:  g.URL,
		Auth: g.auth(),
	})
	if err != nil {
		logger.Error("[git]", "git clone error", err.Error())
		return err
	}
	g.worktree, err = g.repository.Worktree()
	return
}

//UploadDir git upload compress file
//dir compress dir
//filename file name
//description git commit description
func (g *GitUpload) UploadDir(dir, filename string, description string) error {
	z := util.NewZipUtils()
	defer z.Close()
	newFilename := g.getNewFilename(filename)
	err := z.CompressDir(dir, newFilename)
	if err != nil {
		logger.Error("[git]", "compress files error", err.Error())
		return err
	}
	err = g.add(newFilename)
	if err != nil {
		logger.Error("[git]", "git add file error", err.Error())
		return err
	}
	return g.commitAndPush(description)
}

//UploadDirs git upload compress file
//dir compress dir
//filename file name
//description git commit description
func (g *GitUpload) UploadDirs(dirs []string, filename string, description string) error {
	z := util.NewZipUtils()
	defer z.Close()
	newFilename := g.getNewFilename(filename)
	err := z.CompressDirs(dirs, newFilename)
	if err != nil {
		logger.Error("[git]", "compress files error", err.Error())
		return err
	}
	err = g.add(newFilename)
	if err != nil {
		logger.Error("[git]", "git add file error", err.Error())
		return err
	}
	return g.commitAndPush(description)
}

func (g *GitUpload) getNewFilename(filename string) string {
	return g.DirPath +
		string(os.PathSeparator) +
		time.Now().Format("2006-01-02") +
		string(os.PathSeparator) +
		filename
}
