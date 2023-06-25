package data

type Pagination struct {
	Page  int64 `form:"page" binding:"required,numeric" msg:"列表分页参数必需 [page]"`
	Limit int64 `form:"limit" binding:"required,numeric" msg:"列表分页数量必需 [limit]"`
}
