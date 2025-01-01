package handlers

import (
	"company-name/constants/msgkey"
	"company-name/internal/content-blocks"
	"company-name/internal/content-blocks/dtos"
	"company-name/pkg/errors"
	loc "company-name/pkg/localization"
	"company-name/pkg/responses"
	"company-name/pkg/validators"
	"github.com/gin-gonic/gin"
)

type ContentBlocksHandler struct {
	service   blocks.IContentBlocksService
	validator validators.IValidator
}

func NewContentBlocksHandler(service blocks.IContentBlocksService, validator validators.IValidator) *ContentBlocksHandler {
	return &ContentBlocksHandler{
		service:   service,
		validator: validator,
	}
}

func (h *ContentBlocksHandler) CreateContentBlock(c *gin.Context) {
	var request dtos.CreateContentBlockRequest

	if !validators.BindJsonAndValidateRequest(c, &request, h.validator) {
		return
	}

	contentBlock, err := h.service.CreateBlock(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Created(c, loc.L(msgkey.MsgResourceCreated, msgkey.MsgContentBlockResource), contentBlock)
}

func (h *ContentBlocksHandler) UpdateContentBlock(c *gin.Context) {
	var request dtos.UpdateContentBlockRequest

	if !validators.BindJsonAndValidateRequest(c, &request, h.validator) {
		return
	}

	contentBlock, err := h.service.UpdateBlock(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceUpdated, msgkey.MsgContentBlockResource), contentBlock)
}

func (h *ContentBlocksHandler) GetContentBlock(c *gin.Context) {
	var request dtos.GetContentBlockRequest

	if !validators.BindQueryAndValidateRequest(c, &request, h.validator) {
		return
	}

	contentBlock, err := h.service.GetBlock(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceFetched, msgkey.MsgContentBlockResource), contentBlock)
}

func (h *ContentBlocksHandler) GetPageContentBlocks(c *gin.Context) {
	request := dtos.GetPageContentBlocksRequest{Page: c.Param("name")}

	if !validators.ValidateRequestOnly(c, &request, h.validator) {
		return
	}

	contentBlocks, err := h.service.GetPage(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceFetched, msgkey.MsgContentBlockResource), contentBlocks)
}

func (h *ContentBlocksHandler) DeleteContentBlock(c *gin.Context) {
	var request dtos.DeleteBlockRequest

	if !validators.BindQueryAndValidateRequest(c, &request, h.validator) {
		return
	}

	err := h.service.DeleteBlock(c, &request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	responses.NoContent(c, loc.L(msgkey.MsgResourceDeleted, msgkey.MsgContentBlockResource))
}
