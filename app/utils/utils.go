package utils

// GetPages 获取分页参数
func GetPages(page, size int64) (offset int64) {
	offset = (page - 1) * size
	return
}
