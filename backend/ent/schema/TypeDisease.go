package schema
import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)
type TypeDisease struct {
    ent.Schema
}
func (TypeDisease) Fields() []ent.Field {
    return []ent.Field{
		field.String("DiseaseName").
		NotEmpty().
		Unique(),
		
	}
}
func (TypeDisease) Edges() []ent.Edge{
    return []ent.Edge{
        edge.To("fortype", Spaciality.Type),
    }
}