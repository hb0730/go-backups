package uploads

import (
	"errors"
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
	BaseUpload
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
func NewGitUpload(url, username, email, token, dirPath, compress string) *GitUpload {
	g := new(GitUpload)
	g.URL = url
	g.Username = username
	g.Email = email
	g.Token = token
	g.DirPath = dirPath
	g.CompressType = compress
	return g
}

//Close remove git repository
func (g *GitUpload) Close() error {
	return util.RemoveAll(g.DirPath)
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
	err = g.validate()
	if err != nil {
		return err
	}
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
//filename File with name only
//description git commit description
func (g *GitUpload) UploadDir(dir, filename string, description string) error {
	if dir == "" || filename == "" {
		return nil
	}
	compress, err := g.before()
	if err != nil {
		return err
	}
	defer compress.Close()
	newFilename := g.getNewFilename(filename)
	err = compress.CompressDir(&newFilename, dir)
	if err != nil {
		logger.Error("[git]", "compress files error", err.Error())
		return err
	}
	return g.after(newFilename, description)
}

//UploadDirs git upload compress file
//dir compress dir
//filename File with name only
//description git commit description
func (g *GitUpload) UploadDirs(dirs []string, filename string, description string) error {
	if len(dirs) == 0 || filename == "" {
		return nil
	}
	compress, err := g.before()
	if err != nil {
		return err
	}
	defer compress.Close()
	newFilename := g.getNewFilename(filename)
	err = compress.CompressDirs(&newFilename, dirs)
	if err != nil {
		logger.Error("[git]", "compress files error", err.Error())
		return err
	}
	return g.after(newFilename, description)

}
func (g *GitUpload) before() (util.Compress, error) {
	err := g.validate()
	if err != nil {
		return nil, err
	}
	if g.repository == nil {
		err = g.Clone()
		if err != nil {
			return nil, err
		}
	}
	return g.GetCompress()
}

func (g *GitUpload) after(filename, description string) (err error) {
	err = g.add(filename)
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

func (g *GitUpload) validate() error {
	if g.URL == "" || g.Username == "" || g.Token == "" || g.DirPath == "" {
		return errors.New("field Required")
	}
	return nil
}
