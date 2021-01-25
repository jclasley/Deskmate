package tags

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/tylerconlee/Deskmate/server/datastore"
)

// GetAllTagsHandler recieves the request for getting all tags loaded and
// returns a JSON encoded tag object
func GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	loadTags()
	t, err := json.Marshal(T)
	if err != nil {
		fmt.Println("Error marshalling JSON for tags")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(t)

}

// PostTagHandler receives the request of tag data and uses it to save
// the tag data to the database
func PostTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t Tag
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	tag := map[string]interface{}{
		"tag":              t.Tag,
		"slackID":          t.SlackID,
		"groupID":          t.GroupID,
		"channel":          t.Channel,
		"notificationType": t.NotificationType,
		"added":            time.Now(),
	}

	datastore.CreateTag(tag)

	m, err := json.Marshal(T)
	if err != nil {
		fmt.Println("Error marshalling JSON for tags")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)

}

// UpdateTagHandler receives the request with tag data and uses it to update
// the existing tag data in the database
func UpdateTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t Tag
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	ID, err := strconv.Atoi(r.RequestURI)
	if err != nil {
		fmt.Println(err)
	}
	tag := map[string]interface{}{
		"id":               ID,
		"tag":              t.Tag,
		"slackID":          t.SlackID,
		"groupID":          t.GroupID,
		"channel":          t.Channel,
		"notificationType": t.NotificationType,
		"added":            time.Now(),
	}
	datastore.UpdateTag(tag)

}

// DeleteTagHandler receives the DELETE request for the specified tag and removes that tag from the database
func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	t := path.Base(r.RequestURI)
	fmt.Println("Deleting tag ID: ", t)
	ID, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(err)
	}
	removeTag(ID)
	datastore.RemoveTag(ID)
}
