package util

func GetPageStart(page int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * AppSetting.PageSize
	}
	return result
}
