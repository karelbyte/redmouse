package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}

// Fields of the User.
func (User) Fields()  []ent.Field {
	return []ent.Field{
		field.String("names"),
		field.String("email").Unique(),
		field.String("password"),
		field.String("description"),
		field.Bool("status"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
