package tomes

// BOSS_MASTER_ENDPOINT is the base Boss endpoint. When calling a specific boss, use BOSS_ENDPOINT with the boss ID.
const (
	BOSS_MASTER_ENDPOINT = "/wow/boss"
	BOSS_ENDPOINT = "/wow/boss/%d"
)

type BossesMaster struct {
	Bosses []Boss `json:"bosses"`
}

type Boss struct {
	ID                    int `json:"id"`
	Name                  string `json:"name"`
	URLSlug               string `json:"urlSlug"`
	Description           string `json:"description,omitempty"`
	ZoneID                int `json:"zoneId"`
	AvailableInNormalMode bool `json:"availableInNormalMode"`
	AvailableInHeroicMode bool `json:"availableInHeroicMode"`
	Health                int `json:"health"`
	HeroicHealth          int `json:"heroicHealth"`
	Level                 int `json:"level"`
	HeroicLevel           int `json:"heroicLevel"`
	JournalID             int `json:"journalId,omitempty"`
	Npcs                  []BossNPC `json:"npcs"`
	Location              BossLocation `json:"location,omitempty"`
}

type BossNPC struct {
	ID      int `json:"id"`
	Name    string `json:"name"`
	URLSlug string `json:"urlSlug"`
}

type BossLocation struct {
	ID   int `json:"id"`
	Name string `json:"name"`
}