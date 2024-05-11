package models

// on OPT the arc and chapter are splitted into seperated arrays (arcs and entries)
// therefore map them into a single struct to have Arc and entries together.
type Arc struct {
	Name    string     `json:"name"`
	Entries []OPTEntry `json:"entries"`
}
