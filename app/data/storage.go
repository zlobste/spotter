package data

type Storage interface {
	Users() UsersStorage
	Groups() GroupsStorage
	UserGroups() UserGroupsStorage
	Timers() TimersStorage
	Confirmations() ConfirmationsStorage
	Votings() VotingsStorage
	Votes() VotesStorage
}
