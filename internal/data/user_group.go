package data

type UserGroup struct {
	GroupId uint64 `db:"group_id"`
	UserId  uint64 `db:"user_id"`
}

type UserGroupsStorage interface {}