package users

import (
	"github.com/qdimka/go-yt/rest"
	"github.com/qdimka/go-yt/utils"
)

const DefaultFields = "$type,id,avatarUrl,fullName,jabberAccountName,ringId,name,login,banned,email,guest,online,tags(id,name,issues(idReadable)),savedQueries(name,issues(idReadable))"

type User struct {
	Id                string        `json:"id"`
	Login             string        `json:"login"`
	FullName          string        `json:"fullName"`
	Email             string        `json:"email"`
	JabberAccountName string        `json:"jabberAccountName"`
	RingId            string        `json:"ringId"`
	Guest             bool          `json:"guest"`
	Online            bool          `json:"online"`
	Banned            bool          `json:"banned"`
	Tags              []interface{} `json:"tags"`
	SavedQueries      []interface{} `json:"savedQueries"`
	AvatarUrl         string        `json:"avatarUrl"`
	Profiles          interface{}   `json:"profiles"`
	Type              string        `json:"$type"`
}

type Service struct {
	client *rest.Client
}

func NewUsersService(client *rest.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetUsers(fields ...string) (*[]User, error) {
	users := new([]User)

	err := s.client.Get("api/users", utils.ConstructQuery("", fields), nil, users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
