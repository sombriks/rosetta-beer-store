// Code generated by entc, DO NOT EDIT.

package media

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "idmedia"
	// FieldCreationdatemedia holds the string denoting the creationdatemedia field in the database.
	FieldCreationdatemedia = "creationdatemedia"
	// FieldDatamedia holds the string denoting the datamedia field in the database.
	FieldDatamedia = "datamedia"
	// FieldNomemedia holds the string denoting the nomemedia field in the database.
	FieldNomemedia = "nomemedia"
	// FieldMimemedia holds the string denoting the mimemedia field in the database.
	FieldMimemedia = "mimemedia"
	// Table holds the table name of the media in the database.
	Table = "media"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldCreationdatemedia,
	FieldDatamedia,
	FieldNomemedia,
	FieldMimemedia,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
