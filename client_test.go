package runkeeper

import (
	"fmt"
	"testing"
	"time"
)

//sample data:
/*
{
   "type": "Running",
   "equipment": "None",
   "start_time": "Sat, 1 Jan 2011 00:00:00",
   "notes": "My first late-night run",
   "path": [
   {
      "timestamp": 0,
      "altitude": 0,
      "longitude": -70.95182336425782,
      "latitude": 42.312620297384676,
      "type": "start"
   },
   {
      "timestamp": 8,
      "altitude": 0,
      "longitude": -70.95255292510987,
      "latitude": 42.31230294498018,
      "type": "end"
   }
   ],
   "post_to_facebook": true,
   "post_to_twitter": true
 }
*/

func TestPostNewFitnessActivity(t *testing.T) {
	activity := CreateNewFitnessActivity("Test for Sander", 200)
	activity.StartTime = Time(time.Now().Add(-60 * time.Minute))
	activity.Type = "Running"
	client := NewClient("7ad9ba823346491cad21a9c0ad85763a")
	uri, err := client.PostNewFitnessActivity(activity)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("URI: %s", uri)
}
