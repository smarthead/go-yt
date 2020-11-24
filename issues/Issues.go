package issues

import (
	"fmt"
	"github.com/qdimka/go-yt/rest"
	"github.com/qdimka/go-yt/users"
	"github.com/qdimka/go-yt/utils"
)

const (
	DefaultFields = "id,$type,idReadable,summary,project(id,shortName,name),customFields(name,$type,value(name,login))"
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
	Id           string        `json:"id,omitempty"`
	Summary      string        `json:"summary"`
	Description  string        `json:"description,omitempty"`
	Type         string        `json:"$type,omitempty"`
	Project      interface{}   `json:"project"`
	CustomFields []CustomField `json:"customFields,omitempty"`
}

type IssueResult struct {
	Id              string `json:"id"`
	Type            string `json:"$type"`
	NumberInProject int    `json:"numberInProject"`
}

type IssueComment struct {
	Text         string        `json:"text"`
	UsesMarkdown bool          `json:"usesMarkdown,omitempty"`
	TextPreview  string        `json:"textPreview,omitempty"`
	Created      string        `json:"created,omitempty"`
	Updated      string        `json:"updated,omitempty"`
	Author       users.User    `json:"author,omitempty"`
	Issue        Issue         `json:"issue,omitempty"`
	Attachments  []interface{} `json:"attachments,omitempty"`
	Visibility   interface{}   `json:"visibility,omitempty"`
	Deleted      bool          `json:"deleted,omitempty"`
	Type         string        `json:"$type,omitempty"`
}

type Service struct {
	client *rest.Client
}

func NewIssuesService(client *rest.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetIssues(query string, fields ...string) (*[]Issue, error) {
	issues := new([]Issue)
	err := s.client.Get("api/issues",
		utils.ConstructQuery(map[string]string{
			"query": query,
		}, fields), nil, issues)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (s *Service) CreateIssue(issue *Issue) (*IssueResult, error) {
	result := new(IssueResult)

	if err := s.client.Post("api/issues",
		utils.ConstructQuery(nil, []string{"id", "$type", "numberInProject"}),
		issue, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) CommentIssue(issue *Issue, comment *IssueComment) (*IssueResult, error) {
	result := new(IssueResult)

	if err := s.client.Post(fmt.Sprintf("api/issues/%s/comments", issue.Id),
		utils.ConstructQuery(nil, []string{"id", "$type", "numberInProject"}),
		comment, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}
