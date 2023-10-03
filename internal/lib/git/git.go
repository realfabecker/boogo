package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/realfabecker/boogo/internal/lib/pdir"
	"github.com/realfabecker/boogo/internal/lib/spin"
)

// Clone recupera um repositório a partir do github
func Clone(repository string, directory string) error {
	cmd := exec.Command("git", "clone", repository, directory)
	return cmd.Run()
}

// Checkout recupera um projeto e o Clona no diretório
func Checkout(url string, dir string) error {
	label := fmt.Sprintf("git clone %s %s", url, dir)
	if err := spin.NewSpinner(label).WrapStart(func() error {
		return Clone(url, dir)
	}); err != nil {
		return err
	}
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := pdir.RmRf(
		filepath.Join(workdir, dir, ".git"),
	); err != nil {
		return err
	}
	return nil
}
