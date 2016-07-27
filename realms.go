package azeroth

const (
	REALM_ENDPOINT = "/realm/status"
)

type RealmStatus struct {
	Realms []Realm `json:"realms"`
}

type Realm struct {
	Type            string   `json:"type"`
	Population      string   `json:"population"`
	Queue           bool     `json:"queue"`
	Status          bool     `json:"status"`
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Battlegroup     string   `json:"battlegroup"`
	Locale          string   `json:"locale"`
	Timezone        string   `json:"timezone"`
	ConnectedRealms []string `json:"connected_realms"`
}
