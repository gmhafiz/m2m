// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/league"
	"entgo.io/ent/dialect/sql"
)

// League is the model entity for the League schema.
type League struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LeagueQuery when eager-loading is set.
	Edges LeagueEdges `json:"edges"`
}

// LeagueEdges holds the relations/edges for other nodes in the graph.
type LeagueEdges struct {
	// LeagueCertificateType holds the value of the league_certificate_type edge.
	LeagueCertificateType []*LeagueCertificateType `json:"league_certificate_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LeagueCertificateTypeOrErr returns the LeagueCertificateType value or an error if the edge
// was not loaded in eager-loading.
func (e LeagueEdges) LeagueCertificateTypeOrErr() ([]*LeagueCertificateType, error) {
	if e.loadedTypes[0] {
		return e.LeagueCertificateType, nil
	}
	return nil, &NotLoadedError{edge: "league_certificate_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*League) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case league.FieldID:
			values[i] = new(sql.NullInt64)
		case league.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type League", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the League fields.
func (l *League) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case league.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = uint64(value.Int64)
		case league.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		}
	}
	return nil
}

// QueryLeagueCertificateType queries the "league_certificate_type" edge of the League entity.
func (l *League) QueryLeagueCertificateType() *LeagueCertificateTypeQuery {
	return (&LeagueClient{config: l.config}).QueryLeagueCertificateType(l)
}

// Update returns a builder for updating this League.
// Note that you need to call League.Unwrap() before calling this method if this League
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *League) Update() *LeagueUpdateOne {
	return (&LeagueClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the League entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *League) Unwrap() *League {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: League is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *League) String() string {
	var builder strings.Builder
	builder.WriteString("League(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Leagues is a parsable slice of League.
type Leagues []*League

func (l Leagues) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
