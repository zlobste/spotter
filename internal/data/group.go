package data

type Group struct {
	Id          uint64 `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Level       uint64 `db:"level" json:"level"`
}

func (g Group) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"title":       g.Title,
		"description": g.Description,
		"level":       g.Level,
	}

	return result
}

func (g Group) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":          g.Id,
		"title":       g.Title,
		"description": g.Description,
		"level":       g.Level,
	}

	return result
}
