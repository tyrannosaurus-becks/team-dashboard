package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://app.asana.com/api/1.0/"
	limit   = 20 // results per page
)

type Asana struct {
	PersonalAccessToken string
}

// ListProjects hits https://developers.asana.com/docs/get-a-teams-projects
func (a *Asana) ListProjects(teamGid string) ([]*Project, error) {
	endpoint := fmt.Sprintf("teams/%s/projects", teamGid)
	method := http.MethodGet

	resp := &struct {
		Data     []*Project `json:"data"`
		NextPage *NextPage  `json:"next_page"`
	}{}
	if err := a.submitRequest(endpoint, method, resp); err != nil {
		return nil, err
	}
	// TODO paginate.
	return resp.Data, nil
}

// ListTasks hits https://developers.asana.com/docs/get-tasks-from-a-project
func (a *Asana) ListTasks(projectGid string) ([]*Task, error) {
	endpoint := fmt.Sprintf("projects/%s/tasks", projectGid)
	method := http.MethodGet

	resp := &struct {
		Data     []*Task   `json:"data"`
		NextPage *NextPage `json:"next_page"`
	}{}
	if err := a.submitRequest(endpoint, method, resp); err != nil {
		return nil, err
	}
	// TODO paginate.
	return resp.Data, nil
}

func (a *Asana) submitRequest(endpoint, method string, ret interface{}) error {
	u := fmt.Sprintf("%s%s?limit=%d", baseURL, endpoint, limit)
	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+a.PersonalAccessToken)
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("received %d: %s", resp.StatusCode, b)
	}
	if ret == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(ret)
}

type Project struct {
	// TODO
}

type Task struct {
	// TODO
}

type NextPage struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}
