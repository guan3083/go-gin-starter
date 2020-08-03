package util

import (
	"go-gin-starter/pkg/setting"
	"go-gin-starter/request"
)

func GetPaginationParams(startPage int, pageSize int) (int, int) {
	limit := defaultLimitPagination(pageSize, setting.AppSetting.PageSize)
	offset := defaultOffsetPagination(startPage, limit, 0)

	return offset, limit
}

func GetPaginationByCommon(page request.ReqCommonPage) request.ReqCommonPage {
	limit := defaultLimitPagination(page.PageSize, setting.AppSetting.PageSize)
	offset := defaultOffsetPagination(page.PageNo, limit, 0)

	return request.ReqCommonPage{PageNo: offset, PageSize: limit}
}

func defaultOffsetPagination(startPage int, limit int, defaultValue int) int {
	offset := defaultValue

	if startPage > 0 {
		offset = (startPage - 1) * limit
	}

	return offset
}

func defaultLimitPagination(pageSize int, defaultValue int) int {
	if pageSize <= 0 {
		return defaultValue
	}

	limit := pageSize

	sizeLimit := setting.AppSetting.PageSizeLimit
	if limit > sizeLimit {
		return sizeLimit
	}

	return limit
}
