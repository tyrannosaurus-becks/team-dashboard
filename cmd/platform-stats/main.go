package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

const timeFormat = "2006-01-02"

var (
	startFlag = flag.String("start", "", "What day to start, defaults to 30 days ago")
	endFlag   = flag.String("end", "", "The last day end day thru which to go, defaults to today")
)

func main() {
	flag.Parse()
	now := time.Now().UTC()
	if startFlag == nil || *startFlag == "" {
		// Default to 30 days ago.
		thirtyDaysAgo := now.Add(30 * 24 * time.Hour * -1).Format(timeFormat)
		startFlag = &thirtyDaysAgo
	}
	if endFlag == nil || *endFlag == "" {
		today := now.Format(timeFormat)
		endFlag = &today
	}

	var config models.Config
	if err := envconfig.Process("dashboard", &config); err != nil {
		log.Fatal(err)
	}

	asana := clients.Asana{
		PersonalAccessToken: config.AsanaPersonalAccessToken,
	}

	// How many tickets were opened during this time period?
	queryParams := &url.Values{}
	queryParams.Add("teams.any", config.AsanaPlatformTeamGid)
	queryParams.Add("created_on.after", *startFlag)
	queryParams.Add("created_on.before", *endFlag)

	newTasks, err := asana.SearchTasks(config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d Platform tasks were created between %s and %s.", len(newTasks), *startFlag, *endFlag)

	var numTasksWithDays int
	var totalDaysEstimated float64
	for _, newTask := range newTasks {
		task, err := asana.GetTask(newTask.GID)
		if err != nil {
			log.Fatal(err)
		}
		days := numDays(task)
		if days == nil {
			// This task didn't have an estimated number of days.
			continue
		}
		numTasksWithDays++
		totalDaysEstimated += *days
	}

	// What's the average number of days per task?
	avgDaysPerTask := totalDaysEstimated / float64(numTasksWithDays)
	log.Printf("%f was the average day per task", avgDaysPerTask)
}

// The int returned will be nil if a number of days isn't given.
func numDays(task *clients.Task) *float64 {
	for _, customField := range task.CustomFields {
		if fmt.Sprintf("%s", customField["name"]) != "Days" {
			continue
		}
		daysIfc, ok := customField["number_value"]
		if !ok {
			return nil
		}
		days, ok := daysIfc.(float64)
		if !ok {
			return nil
		}
		return &days
	}
	return nil
}
