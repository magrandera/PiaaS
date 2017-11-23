// +build linux,!arm

package models

// Buildpack : stores the different buildpacks
type Buildpack struct {
	Name string
}

// InitBuildpacks : Initializes the database for buildpacks
func InitBuildpacks() {
	records, _ := db.ReadAll("buildpack")
	if len(records) != 0 {
		return
	}
	if err := db.Write("buildpack", "python3", Buildpack{Name: "apihackers/python3"}); err != nil {
		return
	}
}
