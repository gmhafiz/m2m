// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/certificatetype"
	"entgo.io/ent/dialect/sql"
)

// CertificateType is the model entity for the CertificateType schema.
type CertificateType struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CertificateTypeQuery when eager-loading is set.
	Edges CertificateTypeEdges `json:"edges"`
}

// CertificateTypeEdges holds the relations/edges for other nodes in the graph.
type CertificateTypeEdges struct {
	// LeagueCertificateTypeLeagueTypeID holds the value of the league_certificate_type_league_type_id edge.
	LeagueCertificateTypeLeagueTypeID []*LeagueCertificateType `json:"league_certificate_type_league_type_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LeagueCertificateTypeLeagueTypeIDOrErr returns the LeagueCertificateTypeLeagueTypeID value or an error if the edge
// was not loaded in eager-loading.
func (e CertificateTypeEdges) LeagueCertificateTypeLeagueTypeIDOrErr() ([]*LeagueCertificateType, error) {
	if e.loadedTypes[0] {
		return e.LeagueCertificateTypeLeagueTypeID, nil
	}
	return nil, &NotLoadedError{edge: "league_certificate_type_league_type_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CertificateType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certificatetype.FieldID:
			values[i] = new(sql.NullInt64)
		case certificatetype.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CertificateType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertificateType fields.
func (ct *CertificateType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certificatetype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = uint64(value.Int64)
		case certificatetype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ct.Name = value.String
			}
		}
	}
	return nil
}

// QueryLeagueCertificateTypeLeagueTypeID queries the "league_certificate_type_league_type_id" edge of the CertificateType entity.
func (ct *CertificateType) QueryLeagueCertificateTypeLeagueTypeID() *LeagueCertificateTypeQuery {
	return (&CertificateTypeClient{config: ct.config}).QueryLeagueCertificateTypeLeagueTypeID(ct)
}

// Update returns a builder for updating this CertificateType.
// Note that you need to call CertificateType.Unwrap() before calling this method if this CertificateType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CertificateType) Update() *CertificateTypeUpdateOne {
	return (&CertificateTypeClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the CertificateType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CertificateType) Unwrap() *CertificateType {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: CertificateType is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CertificateType) String() string {
	var builder strings.Builder
	builder.WriteString("CertificateType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("name=")
	builder.WriteString(ct.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CertificateTypes is a parsable slice of CertificateType.
type CertificateTypes []*CertificateType

func (ct CertificateTypes) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
