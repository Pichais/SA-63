package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

type Organ struct {
    ent.Schema
}
func (Organ) Fields() []ent.Field {
    return []ent.Field{
		field.String("OrganName").
		NotEmpty(),		
	}
}
func (Organ) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("forogan", Spaciality.Type),
    }
}