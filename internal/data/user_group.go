package data

type UserGroup struct {
	GroupId uint64 `db:"group_id"`
	UserId  uint64 `db:"user_id"`
}

func (u UserGroup) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"group_id": u.UserId,
		"user_id":  u.GroupId,
	}

	return result
}
