package tomes

// AUCTION_ENDPOINT is the base Auction House endpoint. When calling, it requires the realm slug.
const (
	AUCTION_ENDPOINT = "/wow/auction/data/%s"
)

// AuctionAPI is the parent type for the Auction House API. When calling, the Auction House API returns the
// AuctionFileEndpoint type, which contains basic information regarding where to download the Auction House file
// containing the gzipped auctions.
type AuctionAPI struct {
	Files []AuctionFileEndpoint `json:"files"`
}

type AuctionFileEndpoint struct {
	URL          string `json:"url"`

	// LastModified contains the UNIX time signature of when the last modification time of the Auction House data
	// for the specific realm request.
	LastModified int64 `json:"lastModified"`
}

type AuctionFile struct {
	Realms   []AuctionRealm `json:"realms"`
	Auctions []Auction `json:"auctions"`
}

type AuctionRealm struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Auction reflects the actual structure of each individual auction within a realm.
type Auction struct {
	Auc          int `json:"auc"`

	// Item is the corresponding item ID from the Item API.
	Item         int `json:"item"`

	// Owner may contain certain special characters.
	Owner        string `json:"owner"`
	OwnerRealm   string `json:"ownerRealm"`
	Bid          int `json:"bid"`
	Buyout       int `json:"buyout"`
	Quantity     int `json:"quantity"`

	// TimeLeft will be either SHORT, MEDIUM, LONG, or VERY_LONG
	TimeLeft     string `json:"timeLeft"`
	Rand         int `json:"rand"`
	Seed         int `json:"seed"`
	Context      int `json:"context"`
	Modifiers    []AuctionModifier `json:"modifiers,omitempty"`
	PetSpeciesID int `json:"petSpeciesId,omitempty"`
	PetBreedID   int `json:"petBreedId,omitempty"`
	PetLevel     int `json:"petLevel,omitempty"`
	PetQualityID int `json:"petQualityId,omitempty"`
	BonusLists   []AuctionBonus
}

type AuctionModifier struct {
	Type  int `json:"type"`
	Value int `json:"value"`
}

type AuctionBonus struct {
	BonusListID int `json:"bonusListId"`
}