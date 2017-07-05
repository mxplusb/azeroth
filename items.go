package azeroth

import (
	"github.com/jinzhu/gorm"
)

// ITEM_ENDPOINT is the base endpoint for all WoW items. When called, it accepts an integer.
// ITEMSET_ENDPOINT is the base endpoint for all WoW item sets. When called, it accepts an integer.
const (
	ITEM_ENDPOINT    = "/wow/item/"
	ITEMSET_ENDPOINT = "/wow/item/set/"
)

type Item struct {
	gorm.Model
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
}
