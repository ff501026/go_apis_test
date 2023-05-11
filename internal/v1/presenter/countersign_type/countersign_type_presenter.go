package countersign_type

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/countersign_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary 新增會簽類型
// @description 新增會簽類型
// @Tags CountersignType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body countersign_type.Created true "會簽類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/countersign_type [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &countersign_type.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋會簽類型
// @description 條件搜尋會簽類型
// @Tags CountersignType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param CtID query string false "會簽類型織ID"
// @param CountersignName query string false "會簽類型名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=countersign_type.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/countersign_type [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &countersign_type.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CountersignTypeResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一會簽類型
// @description 取得單一會簽類型
// @Tags CountersignType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param CtID path string true "會簽類型ID"
// @success 200 object code.SuccessfulMessage{body=countersign_type.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/countersign_type/{countersign_typeID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	ctId := ctx.Param("ctID")
	input := &countersign_type.Field{}
	input.CtID = ctId
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一會簽類型
// @description 刪除單一會簽類型
// @Tags CountersignType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param CtID path string true "會簽類型ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/countersign_type/{countersign_typeID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ctId := ctx.Param("ctID")
	input := &countersign_type.Updated{}
	input.CtID = ctId
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一會簽類型
// @description 更新單一會簽類型
// @Tags CountersignType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param CtID path string true "會簽類型ID"
// @param * body countersign_type.Updated true "更新會簽類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/countersign_type/{countersign_typeID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ctId := ctx.Param("ctID")
	input := &countersign_type.Updated{}
	input.CtID = ctId
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
