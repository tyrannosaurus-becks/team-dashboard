package clients

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"

	tutil "github.com/tyrannosaurus-becks/team-dashboard/internal/testing"
)

func TestAsana_SearchTasks(t *testing.T) {
	tutil.RequireAcceptanceTestFlag(t)

	client := &Asana{PersonalAccessToken: os.Getenv("DASHBOARD_ASANA_PERSONAL_ACCESS_TOKEN")}
	query := url.Values{}
	query.Add("text", "hello")
	tasks, err := client.SearchTasks(os.Getenv("DASHBOARD_ASANA_WORKSPACE_GID"), query)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.MarshalIndent(tasks, "", "  ")
	fmt.Printf("%s\n", b)
}
