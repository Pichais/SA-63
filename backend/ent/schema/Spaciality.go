package schema
import (
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	
)
type Spaciality struct {
    ent.Schema
}
func (Spaciality) Fields() []ent.Field {
    return nil
}

func (Spaciality) Edges() []ent.Edge {
    return []ent.Edge{
		edge.From("user",Physician.Type).
			Ref("foruser").
			Unique(),
		edge.From("ogan",Organ.Type).
			Ref("forogan").
			Unique(),
		edge.From("type",TypeDisease.Type).
			Ref("fortype").
			Unique(),
	}
}