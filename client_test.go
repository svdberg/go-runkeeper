package runkeeper

import (
	"fmt"
	"testing"
	"time"
)

func TestPostNewFitnessActivity(t *testing.T) {
	activity := FitnessActivity{}
	activity.StartTime = Time(time.Now().Add(-60 * time.Minute))
	activity.Type = "Running"
	activity.Comment = "Test for Sander"
	client := NewClient("7ad9ba823346491cad21a9c0ad85763a")
	uri, err := client.PostNewFitnessActivity(&activity)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("URI: %s", uri)
}
