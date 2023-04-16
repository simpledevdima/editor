package editor

// OptionValue passed value
type OptionValue string

// OptionLabel display title
type OptionLabel string

// SelectOptions select field options
type SelectOptions map[OptionValue]OptionLabel

// Select data type for select field information
type Select struct {
	ShowPlaceHolder bool
	PlaceHolder     string
	Options         SelectOptions
	Default         OptionValue
}
