package tomes

// ACHIEVEMENT_ENDPOINT is the base Achievement endpoint. When calling, it requires the achievement ID.
const (
	ACHIEVEMENT_ENDPOINT = "/wow/achievement/%d"
)

// Achievements describes the base achievement API.
type Achievements struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Points      int `json:"points"`
	Description string `json:"description"`
	Reward      string `json:"reward"`
	RewardItems []Rewards `json:"rewardItems"`
	Icon        string `json:"icon"`
	Criteria    []AchievementCriterium `json:"criteria"`
	AccountWide bool `json:"accountWide"`
	FactionID   int `json:"factionId"`
}

type Rewards struct {
	ID                   int `json:"id"`
	Name                 string `json:"name"`
	Icon                 string `json:"icon"`
	Quality              int `json:"quality"`
	ItemLevel            int `json:"itemLevel"`
	TooltipParams        AcheivementTooltip `json:"tooltipParams"`
	Stats                []interface{} `json:"stats"`
	Armor                int `json:"armor"`
	Context              string `json:"context"`
	BonusLists           []interface{} `json:"bonusLists"`
	ArtifactID           int `json:"artifactId"`
	DisplayInfoID        int `json:"displayInfoId"`
	ArtifactAppearanceID int `json:"artifactAppearanceId"`
	ArtifactTraits       []interface{} `json:"artifactTraits"`
	Relics               []interface{} `json:"relics"`
	Appearance           AchievementAppearance `json:"appearance"`
}

type AchievementCriterium struct {
	ID          int `json:"id"`
	Description string `json:"description"`
	OrderIndex  int `json:"orderIndex"`
	Max         int `json:"max"`
}

type AcheivementTooltip struct {
	 TimewalkerLevel int `json:"timewalkerLevel"`
}

type AchievementAppearance struct {

}