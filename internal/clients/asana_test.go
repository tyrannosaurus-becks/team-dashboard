package clients

import (
	"fmt"
	"net/url"
	"os"
	"testing"
)

func TestAsana_SearchTasks(t *testing.T) {
	client := &Asana{PersonalAccessToken: os.Getenv("DASHBOARD_ASANA_PERSONAL_ACCESS_TOKEN")}
	query := url.Values{}
	query.Add("text", "hello")
	tasks, err := client.SearchTasks(os.Getenv("DASHBOARD_ASANA_WORKSPACE_GID"), query)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tasks)
}
