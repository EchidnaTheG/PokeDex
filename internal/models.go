package internal

type LocationAreaData struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             LocationResult        `json:"location"`
	Names                []AreaName            `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

// Encounter method rate structure
type EncounterMethodRate struct {
	EncounterMethod LocationResult        `json:"encounter_method"`
	VersionDetails  []MethodVersionDetail `json:"version_details"`
}

// Version details for encounter methods
type MethodVersionDetail struct {
	Rate    int            `json:"rate"`
	Version LocationResult `json:"version"`
}

// Area name structure
type AreaName struct {
	Name     string         `json:"name"`
	Language LocationResult `json:"language"`
}

// Pokemon encounter structure
type PokemonEncounter struct {
	Pokemon        LocationResult                  `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

// Version details for pokemon encounters
type PokemonEncounterVersionDetail struct {
	Version          LocationResult    `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

// Individual encounter detail
type EncounterDetail struct {
	MinLevel        int              `json:"min_level"`
	MaxLevel        int              `json:"max_level"`
	ConditionValues []LocationResult `json:"condition_values"`
	Chance          int              `json:"chance"`
	Method          LocationResult   `json:"method"`
}

// Individual location result
type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Main response structure
type LocationData struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`     // pointer because it can be null
	Previous *string          `json:"previous"` // pointer because it can be null
	Results  []LocationResult `json:"results"`  // array of location results
}

type Config struct {
	Next     *string
	Previous *string
}

type Pokemon struct{
	ID                     int                 `json:"id"`
    Name                   string              `json:"name"`
    BaseExperience         int                 `json:"base_experience"`
    Height                 int                 `json:"height"`
    IsDefault              bool                `json:"is_default"`
    Order                  int                 `json:"order"`
    Weight                 int                 `json:"weight"`
    Abilities              []PokemonAbility    `json:"abilities"`
    Forms                  []LocationResult    `json:"forms"`
    GameIndices            []GameIndex         `json:"game_indices"`
    HeldItems              []HeldItem          `json:"held_items"`
    LocationAreaEncounters string              `json:"location_area_encounters"`
    Moves                  []PokemonMove       `json:"moves"`
    Species                LocationResult      `json:"species"`
    Sprites                PokemonSprites      `json:"sprites"`
    Cries                  PokemonCries        `json:"cries"`
    Stats                  []PokemonStat       `json:"stats"`
    Types                  []PokemonType       `json:"types"`
    PastTypes              []PastType          `json:"past_types"`
    PastAbilities          []PastAbility       `json:"past_abilities"`
}

// Pokemon ability structure
type PokemonAbility struct {
    IsHidden bool           `json:"is_hidden"`
    Slot     int            `json:"slot"`
    Ability  LocationResult `json:"ability"`
}

// Game index structure
type GameIndex struct {
    GameIndex int            `json:"game_index"`
    Version   LocationResult `json:"version"`
}

// Held item structure
type HeldItem struct {
    Item           LocationResult           `json:"item"`
    VersionDetails []HeldItemVersionDetail `json:"version_details"`
}

type HeldItemVersionDetail struct {
    Rarity  int            `json:"rarity"`
    Version LocationResult `json:"version"`
}

// Pokemon move structure
type PokemonMove struct {
    Move                LocationResult             `json:"move"`
    VersionGroupDetails []MoveVersionGroupDetail   `json:"version_group_details"`
}

type MoveVersionGroupDetail struct {
    LevelLearnedAt   int            `json:"level_learned_at"`
    VersionGroup     LocationResult `json:"version_group"`
    MoveLearnMethod  LocationResult `json:"move_learn_method"`
    Order            int            `json:"order"`
}

// Pokemon sprites (simplified - you can expand this if needed)
type PokemonSprites struct {
    BackDefault      *string `json:"back_default"`
    BackFemale       *string `json:"back_female"`
    BackShiny        *string `json:"back_shiny"`
    BackShinyFemale  *string `json:"back_shiny_female"`
    FrontDefault     *string `json:"front_default"`
    FrontFemale      *string `json:"front_female"`
    FrontShiny       *string `json:"front_shiny"`
    FrontShinyFemale *string `json:"front_shiny_female"`
    // Note: Omitting the complex nested "other" and "versions" objects for simplicity
    // You can add them if needed
}

// Pokemon cries
type PokemonCries struct {
    Latest string `json:"latest"`
    Legacy string `json:"legacy"`
}

// Pokemon stat structure
type PokemonStat struct {
    BaseStat int            `json:"base_stat"`
    Effort   int            `json:"effort"`
    Stat     LocationResult `json:"stat"`
}

// Pokemon type structure
type PokemonType struct {
    Slot int            `json:"slot"`
    Type LocationResult `json:"type"`
}

// Past type structure
type PastType struct {
    Generation LocationResult  `json:"generation"`
    Types      []PokemonType   `json:"types"`
}

// Past ability structure
type PastAbility struct {
    Generation LocationResult     `json:"generation"`
    Abilities  []PokemonAbility   `json:"abilities"`
}