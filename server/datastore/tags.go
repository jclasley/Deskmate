package datastore

// LoadTags returns an array of map[string]interfaces that
// contain all of the tags current stored in the `tags`
// table of the database
func LoadTags() {}

// CreateTag takes the metadata needed for saving a tag and
// saves the data to the `tags` table in the database
func CreateTag() {}

// RemoveTag removes the specified tag from the database so that
// it isn't loaded on the next restart
func RemoveTag() {}

// UpdateTag updates the specified tag with the provided metadata
func UpdateTag() {}
