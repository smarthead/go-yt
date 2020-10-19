package issues

import (
	"github.com/qdimka/go-yt/rest"
	"github.com/qdimka/go-yt/utils"
)

const (
	DefaultFields = "id,idReadable,summary,project(id,shortName,name),customFields(name,$type,value(name,login))"
)

type CustomFieldValue struct {
	Name  string
	Login string
}

type CustomField struct {
	Name  string      `json:"name"`
	Type  string      `json:"$type"`
	Value interface{} `json:"value"`
}

type Issue struct {
	Summary      string        `json:"summary"`
	Description  string        `json:"description"`
	Project      interface{}   `json:"project"`
	CustomFields []CustomField `json:"customFields"`
}

type Service struct {
	client *rest.Client
}

func NewIssuesService(client *rest.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetIssues(query string, fields ...string) (*[]Issue, error) {
	issues := new([]Issue)
	err := s.client.Get("api/issues", utils.ConstructQuery(query, fields), nil, issues)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (s *Service) CreateIssue(issue Issue) error {
	err := s.client.Post("api/issues", issue, nil, nil)
	return err
}

func (s *Service) Comment(i int) error {
	panic("implement me")
}
