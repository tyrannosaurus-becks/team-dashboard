package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://app.asana.com/api/1.0/"
	limit   = 20
)

type Asana struct {
	PersonalAccessToken string
}

// SearchTasks hits https://developers.asana.com/docs/search-tasks-in-a-workspace
func (a *Asana) SearchTasks(workspaceGid string, queryParams url.Values) ([]*ObjectMetadata, error) {
	endpoint := fmt.Sprintf("workspaces/%s/tasks/search", workspaceGid)

	var tasks []*ObjectMetadata
	queryParams.Add("limit", fmt.Sprintf("%d", limit))
	u := getUrl(endpoint) + "?" + queryParams.Encode()
	for {
		resp := &struct {
			Data     []*ObjectMetadata `json:"data"`
			NextPage *NextPage         `json:"next_page"`
		}{}
		if err := a.submitRequest(u, resp); err != nil {
			return nil, err
		}
		tasks = append(tasks, resp.Data...)
		if resp.NextPage == nil {
			break
		}
		u = resp.NextPage.URI
	}
	return tasks, nil
}

// GetTask hits https://developers.asana.com/docs/get-a-task
func (a *Asana) GetTask(taskGid string) (*Task, error) {
	endpoint := fmt.Sprintf("tasks/%s", taskGid)

	resp := &struct {
		Data *Task `json:"data"`
	}{}
	if err := a.submitRequest(getUrl(endpoint), resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// GetProject hits https://developers.asana.com/docs/get-a-project
func (a *Asana) GetProject(projectGid string) (*Project, error) {
	endpoint := fmt.Sprintf("projects/%s", projectGid)

	resp := &struct {
		Data *Project `json:"data"`
	}{}
	if err := a.submitRequest(getUrl(endpoint), resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (a *Asana) submitRequest(url string, ret interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func getUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}

type ObjectMetadata struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

type Task struct {
	CustomFields []map[string]interface{} `json:"custom_fields"`
	Projects     []*ObjectMetadata        `json:"projects"`
}

type Project struct {
	Team *ObjectMetadata `json:"team"`
}

type NextPage struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}
