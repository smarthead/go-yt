package projects

import (
	"github.com/qdimka/go-yt/rest"
	"github.com/qdimka/go-yt/utils"
)

const (
	DefaultFields = "id,shortName,name"
)

type Project struct {
	Id          string `json:"id"`
	ShortName   string `json:"shortName"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Service struct {
	client *rest.Client
}

func NewProjectsService(client *rest.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetProjects(query string, fields ...string) (*[]Project, error) {
	projects := new([]Project)

	err := s.client.Get("api/admin/projects", utils.ConstructQuery(query, fields), nil, projects)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *Service) GetProjectById(projectId string, fields ...string) (*Project, error) {
	project := new(Project)

	err := s.client.Get("api/admin/projects/"+projectId, utils.ConstructQuery("", fields), nil, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}
