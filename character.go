package azeroth

// CHARACTER_ENDPOINT is the base endpoint for WoW characters. When referenced, it needs the realm and the character, in
// that order.
const (
	CHARACTER_ENDPOINT = "/wow/character/%s/%s"
)

// TODO add hunterPets endpoint (I don't have any hunter pets, lol)
// TODO export the character item fields - this is on hold because I am working on the items stuff first.

// Character is the primary way to access character information. This Character Profile API can be used to fetch a
// single character at a time through an HTTP GET request to a URL describing the character profile resource. By
// default, a basic dataset will be returned and with each request and zero or more additional fields can be retrieved.
// To access this API, craft a resource URL pointing to the character who's information is to be retrieved.
type Character struct {
	// LastModified is in UNIX time.
	LastModified        int64                 `json:"lastModified"`
	Name                string                `json:"name"`
	Realm               string                `json:"realm"`
	Battlegroup         string                `json:"battlegroup"`
	Class               int                   `json:"class"`
	Race                int                   `json:"race"`
	Gender              int                   `json:"gender"`
	Level               int                   `json:"level"`
	AchievementPoints   int                   `json:"achievementPoints"`
	Thumbnail           string                `json:"thumbnail"`
	CalcClass           string                `json:"calcClass"`
	Faction             int                   `json:"faction"`
	TotalHonorableKills int                   `json:"totalHonorableKills"`
	Achievements        CharacterAchievements `json:"achievements,omitempty"`
	Appearance          CharacterAppearance   `json:"appearance,omitempty"`
	Feed                []CharacterFeed       `json:"feed,omitempty"`
	Guild               CharacterGuild        `json:"guild,omitempty"`
	Pets                CharacterPets         `json:"pets,omitempty"`
	PetSlots            []CharacterPetSlots   `json:"petSlots,omitempty"`
	Professions         CharacterProfessions  `json:"professions,omitempty"`
	Progression         CharacterProgression  `json:"progression,omitempty"`
	PVP                 CharacterPvp          `json:"pvp,omitempty"`
	Quests              CharacterQuests       `json:"quests,omitempty"`
	Reputation          []CharacterReputation `json:"reputation,omitempty"`
	Statistics          CharacterStatistics   `json:"statistics,omitempty"`
	Stats               CharacterStats        `json:"stats,omitempty"`
	Talents             []CharacterTalents    `json:"talents,omitempty"`
	Titles              []CharacterTitles     `json:"titles,omitempty"`
	Audit               CharacterAudit        `json:"audit,omitempty"`
}

// CharacterAchievements is a map of achievement data including completion timestamps and criteria information.
type CharacterAchievements struct {
	AchievementsCompleted          []int         `json:"achievementsCompleted"`
	AchievementsCompletedTimestamp []int64       `json:"achievementsCompletedTimestamp"`
	Criteria                       []int         `json:"criteria"`
	CriteriaQuantity               []interface{} `json:"criteriaQuantity"`
	CriteriaTimestamp              []int64       `json:"criteriaTimestamp"`
	CriteriaCreated                []interface{} `json:"criteriaCreated"`
}

// CharacterAppearance is a map of a character's appearance settings such as which face texture they've selected and
// whether or not a healm is visible.
type CharacterAppearance struct {
	FaceVariation        int   `json:"faceVariation"`
	SkinColor            int   `json:"skinColor"`
	HairVariation        int   `json:"hairVariation"`
	HairColor            int   `json:"hairColor"`
	FeatureVariation     int   `json:"featureVariation"`
	ShowHelm             bool  `json:"showHelm"`
	ShowCloak            bool  `json:"showCloak"`
	CustomDisplayOptions []int `json:"customDisplayOptions"`
}

// CharacterFeed is the activity feed of the character.
type CharacterFeed struct {
	Type           string                        `json:"type"`
	Timestamp      int64                         `json:"timestamp"`
	Achievement    CharacterFeedAchievement      `json:"achievement"`
	FeatOfStrength bool                          `json:"featOfStrength"`
	Criteria       CharacterAchievementCriterium `json:"criteria,omitempty"`
	Quantity       int                           `json:"quantity,omitempty"`
	Name           string                        `json:"name,omitempty"`
}

// CharacterAchievement does not map all fields to the Achievements API, so it should be used separately.
type CharacterFeedAchievement struct {
	ID          int                             `json:"id"`
	Title       string                          `json:"title"`
	Points      int                             `json:"points"`
	Description string                          `json:"description"`
	RewardItems []interface{}                   `json:"rewardItems"`
	Icon        string                          `json:"icon"`
	Criteria    []CharacterAchievementCriterium `json:"criteria"`
	AccountWide bool                            `json:"accountWide"`
	FactionID   int                             `json:"factionId"`
}

type CharacterAchievementCriterium struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
	Max         int    `json:"max"`
}

// CharacterGuild is a summary of the guild that the character belongs to. If the character does not belong to a guild
// and this field is requested, this field will not be exposed. When a guild is requested, a struct is returned with
// a basic set of guild information. Note that the rank of the character is not included in this block as it describes
// a guild and not a membership of the guild. To retrieve the character's rank within the guild, you must specific a
// separate request to the guild profile resource.
type CharacterGuild struct {
	Name              string               `json:"name"`
	Realm             string               `json:"realm"`
	Battlegroup       string               `json:"battlegroup"`
	Members           int                  `json:"members"`
	AchievementPoints int                  `json:"achievementPoints"`
	Emblem            CharacterGuildEmblem `json:"emblem"`
}

type CharacterGuildEmblem struct {
	Icon              int    `json:"icon"`
	IconColor         string `json:"iconColor"`
	IconColorID       int    `json:"iconColorId"`
	Border            int    `json:"border"`
	BorderColor       string `json:"borderColor"`
	BorderColorID     int    `json:"borderColorId"`
	BackgroundColor   string `json:"backgroundColor"`
	BackgroundColorID int    `json:"backgroundColorId"`
}

// CharacterMount is a list of all of the mounts obtained by the character.
type CharacterMount struct {
	NumCollected    int                       `json:"numCollected"`
	NumNotCollected int                       `json:"numNotCollected"`
	Collected       []CharacterMountCollected `json:"collected"`
}

type CharacterMountCollected struct {
	Name       string `json:"name"`
	SpellID    int    `json:"spellId"`
	CreatureID int    `json:"creatureId"`
	ItemID     int    `json:"itemId"`
	QualityID  int    `json:"qualityId"`
	Icon       string `json:"icon"`
	IsGround   bool   `json:"isGround"`
	IsFlying   bool   `json:"isFlying"`
	IsAquatic  bool   `json:"isAquatic"`
	IsJumping  bool   `json:"isJumping"`
}

// CharacterPets is a list of the battle pets obtained by the character.
type CharacterPets struct {
	NumCollected    int                      `json:"numCollected"`
	NumNotCollected int                      `json:"numNotCollected"`
	Collected       []CharacterPetsCollected `json:"collected"`
}

type CharacterPetsCollected struct {
	Name                        string             `json:"name"`
	SpellID                     int                `json:"spellId"`
	CreatureID                  int                `json:"creatureId"`
	ItemID                      int                `json:"itemId"`
	QualityID                   int                `json:"qualityId"`
	Icon                        string             `json:"icon"`
	Stats                       CharacterPetsStats `json:"stats"`
	BattlePetGUID               string             `json:"battlePetGuid"`
	IsFavorite                  bool               `json:"isFavorite"`
	IsFirstAbilitySlotSelected  bool               `json:"isFirstAbilitySlotSelected"`
	IsSecondAbilitySlotSelected bool               `json:"isSecondAbilitySlotSelected"`
	IsThirdAbilitySlotSelected  bool               `json:"isThirdAbilitySlotSelected"`
	CreatureName                string             `json:"creatureName"`
	CanBattle                   bool               `json:"canBattle"`
}

type CharacterPetsStats struct {
	SpeciesID    int `json:"speciesId"`
	BreedID      int `json:"breedId"`
	PetQualityID int `json:"petQualityId"`
	Level        int `json:"level"`
	Health       int `json:"health"`
	Power        int `json:"power"`
	Speed        int `json:"speed"`
}

// CharacterPetSlots is some data about the current battle pet slots on this characters account. The pet slot contains which slot it is and whether the slot is empty or locked. We also include the battlePetId which is unique for this character and can be used to match a battlePetId in the pets field for this character. The ability list is the list of 3 active abilities on that pet. If the pet is not high enough level than it will always be the first three abilities that the pet has.
type CharacterPetSlots struct {
	Slot          int    `json:"slot"`
	BattlePetGUID string `json:"battlePetGuid"`
	IsEmpty       bool   `json:"isEmpty"`
	IsLocked      bool   `json:"isLocked"`
	Abilities     []int  `json:"abilities"`
}

// CharacterProfessions is a list of the character's professions. Does not include class professions.
type CharacterProfessions struct {
	Primary   []CharacterProfessionDetails `json:"primary"`
	Secondary []CharacterProfessionDetails `json:"secondary"`
}

type CharacterProfessionDetails struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Rank    int    `json:"rank"`
	Max     int    `json:"max"`
	Recipes []int  `json:"recipes"`
}

// CharacterProgression is a list of raids and bosses indicating raid progression and completeness.
type CharacterProgression struct {
	Raids []CharacterRaids `json:"raids"`
}

type CharacterRaids struct {
	Name   string                `json:"name"`
	Lfr    int                   `json:"lfr"`
	Normal int                   `json:"normal"`
	Heroic int                   `json:"heroic"`
	Mythic int                   `json:"mythic"`
	ID     int                   `json:"id"`
	Bosses []CharacterRaidBosses `json:"bosses"`
}

type CharacterRaidBosses struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	NormalKills     int    `json:"normalKills"`
	NormalTimestamp int64  `json:"normalTimestamp"`
}

// CharacterPvp is a map of pvp information including arena team membership and rated battlegrounds information.
// TODO find more players who PVP to help build this modelling.
type CharacterPvp struct {
	Brackets struct {
		ARENABRACKET2V2 struct {
			Slug         string `json:"slug"`
			Rating       int    `json:"rating"`
			WeeklyPlayed int    `json:"weeklyPlayed"`
			WeeklyWon    int    `json:"weeklyWon"`
			WeeklyLost   int    `json:"weeklyLost"`
			SeasonPlayed int    `json:"seasonPlayed"`
			SeasonWon    int    `json:"seasonWon"`
			SeasonLost   int    `json:"seasonLost"`
		} `json:"ARENA_BRACKET_2v2"`
		ARENABRACKET3V3 struct {
			Slug         string `json:"slug"`
			Rating       int    `json:"rating"`
			WeeklyPlayed int    `json:"weeklyPlayed"`
			WeeklyWon    int    `json:"weeklyWon"`
			WeeklyLost   int    `json:"weeklyLost"`
			SeasonPlayed int    `json:"seasonPlayed"`
			SeasonWon    int    `json:"seasonWon"`
			SeasonLost   int    `json:"seasonLost"`
		} `json:"ARENA_BRACKET_3v3"`
		ARENABRACKETRBG struct {
			Slug         string `json:"slug"`
			Rating       int    `json:"rating"`
			WeeklyPlayed int    `json:"weeklyPlayed"`
			WeeklyWon    int    `json:"weeklyWon"`
			WeeklyLost   int    `json:"weeklyLost"`
			SeasonPlayed int    `json:"seasonPlayed"`
			SeasonWon    int    `json:"seasonWon"`
			SeasonLost   int    `json:"seasonLost"`
		} `json:"ARENA_BRACKET_RBG"`
		ARENABRACKET2V2SKIRMISH struct {
			Slug         string `json:"slug"`
			Rating       int    `json:"rating"`
			WeeklyPlayed int    `json:"weeklyPlayed"`
			WeeklyWon    int    `json:"weeklyWon"`
			WeeklyLost   int    `json:"weeklyLost"`
			SeasonPlayed int    `json:"seasonPlayed"`
			SeasonWon    int    `json:"seasonWon"`
			SeasonLost   int    `json:"seasonLost"`
		} `json:"ARENA_BRACKET_2v2_SKIRMISH"`
		UNKNOWN struct {
			Slug         string `json:"slug"`
			Rating       int    `json:"rating"`
			WeeklyPlayed int    `json:"weeklyPlayed"`
			WeeklyWon    int    `json:"weeklyWon"`
			WeeklyLost   int    `json:"weeklyLost"`
			SeasonPlayed int    `json:"seasonPlayed"`
			SeasonWon    int    `json:"seasonWon"`
			SeasonLost   int    `json:"seasonLost"`
		} `json:"UNKNOWN"`
	} `json:"brackets"`
}

// CharacterQuests is a list of completed quests by a character.
type CharacterQuests struct {
	Quests []int `json:"quests,omitempty"`
}

// CharacterReputation is a list of the factions that the character has an associated reputation with.
type CharacterReputation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Standing int    `json:"standing"`
	Value    int    `json:"value"`
	Max      int    `json:"max"`
}

// CharacterStatistics is a map of character statistics.
type CharacterStatistics struct {
	ID            int                                `json:"id"`
	Name          string                             `json:"name"`
	SubCategories []CharacterStatisticsSubCategories `json:"subCategories"`
}

type CharacterStatisticsSubCategories struct {
	ID            int                      `json:"id"`
	Name          string                   `json:"name"`
	Statistics    []CharacterSubStatistics `json:"statistics"`
	SubCategories []struct {
		ID         int                      `json:"id"`
		Name       string                   `json:"name"`
		Statistics []CharacterSubStatistics `json:"statistics"`
	}
}

type CharacterSubStatistics struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	LastUpdated int64  `json:"lastUpdated"`
	Money       bool   `json:"money"`
}

// CharacterStats is a map of character attributes and stats.
type CharacterStats struct {
	Health                      int     `json:"health"`
	PowerType                   string  `json:"powerType"`
	Power                       int     `json:"power"`
	Str                         int     `json:"str"`
	Agi                         int     `json:"agi"`
	Int                         int     `json:"int"`
	Sta                         int     `json:"sta"`
	SpeedRating                 float64 `json:"speedRating"`
	SpeedRatingBonus            int     `json:"speedRatingBonus"`
	Crit                        float64 `json:"crit"`
	CritRating                  int     `json:"critRating"`
	Haste                       float64 `json:"haste"`
	HasteRating                 int     `json:"hasteRating"`
	HasteRatingPercent          float64 `json:"hasteRatingPercent"`
	Mastery                     float64 `json:"mastery"`
	MasteryRating               int     `json:"masteryRating"`
	Leech                       float64 `json:"leech"`
	LeechRating                 float64 `json:"leechRating"`
	LeechRatingBonus            float64 `json:"leechRatingBonus"`
	Versatility                 int     `json:"versatility"`
	VersatilityDamageDoneBonus  float64 `json:"versatilityDamageDoneBonus"`
	VersatilityHealingDoneBonus float64 `json:"versatilityHealingDoneBonus"`
	VersatilityDamageTakenBonus float64 `json:"versatilityDamageTakenBonus"`
	AvoidanceRating             float64 `json:"avoidanceRating"`
	AvoidanceRatingBonus        int     `json:"avoidanceRatingBonus"`
	SpellPen                    int     `json:"spellPen"`
	SpellCrit                   float64 `json:"spellCrit"`
	SpellCritRating             int     `json:"spellCritRating"`
	Mana5                       int     `json:"mana5"`
	Mana5Combat                 int     `json:"mana5Combat"`
	Armor                       int     `json:"armor"`
	Dodge                       int     `json:"dodge"`
	DodgeRating                 int     `json:"dodgeRating"`
	Parry                       int     `json:"parry"`
	ParryRating                 int     `json:"parryRating"`
	Block                       int     `json:"block"`
	BlockRating                 int     `json:"blockRating"`
	MainHandDmgMin              int     `json:"mainHandDmgMin"`
	MainHandDmgMax              int     `json:"mainHandDmgMax"`
	MainHandSpeed               float64 `json:"mainHandSpeed"`
	MainHandDps                 float64 `json:"mainHandDps"`
	OffHandDmgMin               int     `json:"offHandDmgMin"`
	OffHandDmgMax               int     `json:"offHandDmgMax"`
	OffHandSpeed                float64 `json:"offHandSpeed"`
	OffHandDps                  float64 `json:"offHandDps"`
	RangedDmgMin                int     `json:"rangedDmgMin"`
	RangedDmgMax                int     `json:"rangedDmgMax"`
	RangedSpeed                 int     `json:"rangedSpeed"`
	RangedDps                   int     `json:"rangedDps"`
}

// CharacterTalents is a list of talent structures.
type CharacterTalents struct {
	Selected   bool                      `json:"selected,omitempty"`
	Talents    []CharacterSpecificTalent `json:"talents"`
	Spec       CharacterSpec             `json:"spec,omitempty"`
	CalcTalent string                    `json:"calcTalent"`
	CalcSpec   string                    `json:"calcSpec"`
}

type CharacterSpecificTalent struct {
	Tier   int            `json:"tier"`
	Column int            `json:"column"`
	Spell  CharacterSpell `json:"spell"`
	Spec   CharacterSpec  `json:"spec,omitempty"`
}

type CharacterSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

// TODO this may be related to the spells API.
type CharacterSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

type CharacterTitles struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected,omitempty"`
}

// TODO need more characters to validate this structure.
type CharacterAudit struct {
	NumberOfIssues int `json:"numberOfIssues"`
	Slots          struct {
		Num2 int `json:"2"`
		Num4 int `json:"4"`
		Num5 int `json:"5"`
		Num6 int `json:"6"`
		Num7 int `json:"7"`
		Num8 int `json:"8"`
		Num9 int `json:"9"`
	} `json:"slots"`
	EmptyGlyphSlots     int  `json:"emptyGlyphSlots"`
	UnspentTalentPoints int  `json:"unspentTalentPoints"`
	NoSpec              bool `json:"noSpec"`
	UnenchantedItems    struct {
		Num2 int `json:"2"`
		Num4 int `json:"4"`
		Num6 int `json:"6"`
		Num7 int `json:"7"`
		Num8 int `json:"8"`
		Num9 int `json:"9"`
	} `json:"unenchantedItems"`
	EmptySockets          int `json:"emptySockets"`
	ItemsWithEmptySockets struct {
	} `json:"itemsWithEmptySockets"`
	AppropriateArmorType   int `json:"appropriateArmorType"`
	InappropriateArmorType struct {
	} `json:"inappropriateArmorType"`
	LowLevelItems struct {
	} `json:"lowLevelItems"`
	LowLevelThreshold   int `json:"lowLevelThreshold"`
	MissingExtraSockets struct {
		Num5 int `json:"5"`
	} `json:"missingExtraSockets"`
	RecommendedBeltBuckle struct {
		ID          int           `json:"id"`
		Description string        `json:"description"`
		Name        string        `json:"name"`
		Icon        string        `json:"icon"`
		Stackable   int           `json:"stackable"`
		ItemBind    int           `json:"itemBind"`
		BonusStats  []interface{} `json:"bonusStats"`
		ItemSpells  []struct {
			SpellID int `json:"spellId"`
			Spell   struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			NCharges   int    `json:"nCharges"`
			Consumable bool   `json:"consumable"`
			CategoryID int    `json:"categoryId"`
			Trigger    string `json:"trigger"`
		} `json:"itemSpells"`
		BuyPrice          int  `json:"buyPrice"`
		ItemClass         int  `json:"itemClass"`
		ItemSubClass      int  `json:"itemSubClass"`
		ContainerSlots    int  `json:"containerSlots"`
		InventoryType     int  `json:"inventoryType"`
		Equippable        bool `json:"equippable"`
		ItemLevel         int  `json:"itemLevel"`
		MaxCount          int  `json:"maxCount"`
		MaxDurability     int  `json:"maxDurability"`
		MinFactionID      int  `json:"minFactionId"`
		MinReputation     int  `json:"minReputation"`
		Quality           int  `json:"quality"`
		SellPrice         int  `json:"sellPrice"`
		RequiredSkill     int  `json:"requiredSkill"`
		RequiredLevel     int  `json:"requiredLevel"`
		RequiredSkillRank int  `json:"requiredSkillRank"`
		ItemSource        struct {
			SourceID   int    `json:"sourceId"`
			SourceType string `json:"sourceType"`
		} `json:"itemSource"`
		BaseArmor            int           `json:"baseArmor"`
		HasSockets           bool          `json:"hasSockets"`
		IsAuctionable        bool          `json:"isAuctionable"`
		Armor                int           `json:"armor"`
		DisplayInfoID        int           `json:"displayInfoId"`
		NameDescription      string        `json:"nameDescription"`
		NameDescriptionColor string        `json:"nameDescriptionColor"`
		Upgradable           bool          `json:"upgradable"`
		HeroicTooltip        bool          `json:"heroicTooltip"`
		Context              string        `json:"context"`
		BonusLists           []interface{} `json:"bonusLists"`
		AvailableContexts    []string      `json:"availableContexts"`
		BonusSummary         struct {
			DefaultBonusLists []interface{} `json:"defaultBonusLists"`
			ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
			BonusChances      []interface{} `json:"bonusChances"`
		} `json:"bonusSummary"`
		ArtifactID int `json:"artifactId"`
	} `json:"recommendedBeltBuckle"`
	MissingBlacksmithSockets struct {
	} `json:"missingBlacksmithSockets"`
	MissingEnchanterEnchants struct {
	} `json:"missingEnchanterEnchants"`
	MissingEngineerEnchants struct {
	} `json:"missingEngineerEnchants"`
	MissingScribeEnchants struct {
	} `json:"missingScribeEnchants"`
	NMissingJewelcrafterGems     int `json:"nMissingJewelcrafterGems"`
	MissingLeatherworkerEnchants struct {
	} `json:"missingLeatherworkerEnchants"`
}
