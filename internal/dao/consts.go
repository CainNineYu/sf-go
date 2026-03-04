package dao

// 顺序："asc"   so != "desc"
type OrderBy string

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RolePartner Role = "partner"
	RoleLeader  Role = "leader"
)

type OrderActionsType string

const (
	OrderActionsTypeOpen          OrderActionsType = "open"
	OrderActionsTypeCloseAll      OrderActionsType = "closeAll"
	OrderActionsTypeCancelTpSl    OrderActionsType = "cancelTpSl"
	OrderActionsTypeSetTpSl       OrderActionsType = "setTpSl"
	OrderActionsTypeClose         OrderActionsType = "close"
	OrderActionsTypeCancelPending OrderActionsType = "cancelPending"
)
