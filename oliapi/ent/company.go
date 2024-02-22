// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"oliapi/ent/company"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL string `json:"logo_url,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ArchivedAt holds the value of the "archived_at" field.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges        CompanyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// Bots holds the value of the bots edge.
	Bots []*Bot `json:"bots,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BotsOrErr returns the Bots value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) BotsOrErr() ([]*Bot, error) {
	if e.loadedTypes[0] {
		return e.Bots, nil
	}
	return nil, &NotLoadedError{edge: "bots"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldName, company.FieldLogoURL, company.FieldEmail:
			values[i] = new(sql.NullString)
		case company.FieldCreatedAt, company.FieldUpdatedAt, company.FieldArchivedAt:
			values[i] = new(sql.NullTime)
		case company.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case company.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				c.LogoURL = value.String
			}
		case company.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case company.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case company.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case company.FieldArchivedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field archived_at", values[i])
			} else if value.Valid {
				c.ArchivedAt = new(time.Time)
				*c.ArchivedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Company.
// This includes values selected through modifiers, order, etc.
func (c *Company) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryBots queries the "bots" edge of the Company entity.
func (c *Company) QueryBots() *BotQuery {
	return NewCompanyClient(c.config).QueryBots(c)
}

// QueryUsers queries the "users" edge of the Company entity.
func (c *Company) QueryUsers() *UserQuery {
	return NewCompanyClient(c.config).QueryUsers(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return NewCompanyClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(c.LogoURL)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.ArchivedAt; v != nil {
		builder.WriteString("archived_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company
