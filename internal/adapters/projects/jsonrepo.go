package projects

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/entities"
	"github.com/realfabecker/bogo/internal/core/ports"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// JsonProjectRepository memory project repository struct
type JsonProjectRepository struct {
	ports.Logger
}

// NewJsonProjectRepository JsonProjectRepository construtor
func NewJsonProjectRepository(logger ports.Logger) ports.ProjectRepository {
	return &JsonProjectRepository{logger}
}

// Get return a repository by its name
func (m JsonProjectRepository) Get(name string) (*domain.Project, error) {
	p, err := m.List()
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	for _, v := range p {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("%s is not a valid repository", name)
}

// List return a list of repositories
func (m JsonProjectRepository) List() ([]domain.Project, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	d, err := os.ReadFile(filepath.Join(h, ".bogo", "repositories.json"))
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	r := make([]domain.Project, 0)
	if err := json.Unmarshal(d, &r); err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	return r, nil
}

// Sync obtain repository configration
func (m JsonProjectRepository) Sync(url string) error {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	r, err := client.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			m.Logger.Error(err.Error())
		}
	}()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errors.New("unable to download repo config")
	}

	if err := m.Validate(data); err != nil {
		return err
	}

	h, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(h, ".bogo", "repositories.json"), data, 0644)
}

// Validate method validate definition
func (m JsonProjectRepository) Validate(data []byte) error {
	loader := gojsonschema.NewStringLoader(entities.ProjectListSchema)

	s := gojsonschema.NewSchemaLoader()
	schema, err := s.Compile(loader)
	if err != nil {
		return fmt.Errorf("compile: %w", err)
	}

	document := gojsonschema.NewStringLoader(string(data))
	result, err := schema.Validate(document)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	if !result.Valid() {
		var e string
		for _, v := range result.Errors() {
			e = e + v.String() + ";"
		}
		return errors.New(e)
	}
	return nil
}
