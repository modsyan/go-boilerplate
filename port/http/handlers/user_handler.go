package handlers

import (
	"company-name/constants/msgkey"
	"company-name/internal/user"
	"company-name/internal/user/dtos"
	"company-name/pkg/errors"
	loc "company-name/pkg/localization"
	"company-name/pkg/responses"
	"company-name/pkg/validators"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service   user.IUserService
	validator validators.IValidator
}

func NewUserHandler(service user.IUserService, validator validators.IValidator) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var request dtos.CreateUserRequest

	if !validators.BindJsonAndValidateRequest(c, &request, h.validator) {
		return
	}

	user, err := h.service.CreateUser(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Created(c, loc.L(msgkey.MsgResourceCreated, msgkey.MsgUserResource), user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var request dtos.UpdateUserRequest

	if !validators.BindJsonAndValidateRequest(c, &request, h.validator) {
		return
	}

	id := c.Param("id")
	if id != request.ID {
		responses.BadRequest(c, loc.L(msgkey.ErrUnMatchedId), nil)
		return
	}

	user, err := h.service.UpdateUser(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceUpdated, msgkey.MsgUserResource), user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var request = dtos.DeleteUserRequest{ID: c.Param("id")}

	if !validators.ValidateRequestOnly(c, &request, h.validator) {
		return
	}

	err := h.service.DeleteUser(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.NoContent(c, loc.L(msgkey.MsgResourceDeleted, msgkey.MsgUserResource))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var request dtos.GetPaginatedUsersRequest

	if !validators.BindQueryAndValidateRequest(c, &request, h.validator) {
		return
	}

	users, err := h.service.GetPaginatedUsers(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceFetched, msgkey.MsgUserResource), users)
}

func (h *UserHandler) GetDetailsUserByID(c *gin.Context) {
	var request = dtos.GetUserDetailsRequest{ID: c.Param("id")}

	if !validators.ValidateRequestOnly(c, &request, h.validator) {
		return
	}

	user, err := h.service.GetUserDetailsById(c, &request)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgResourceFetched, msgkey.MsgUserResource), user)
}
