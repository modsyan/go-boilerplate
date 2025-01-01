package msgkey

const (

	// "-----------Start---------0": "-----------Start---------0",

	MsgSuccess = "success"
	MsgError   = "error"

	// "1-------------------------": "1-------------------------",
	// "-------Http_Errors--------": "-------Http_Errors--------",
	// "-------------------------1": "-------------------------1",

	ErrNotFound            = "not_found"
	ErrInternalServerError = "internal_server_error"
	ErrBadRequest          = "bad_request"
	ErrForbidden           = "forbidden"
	ErrUnauthorized        = "unauthorized"
	ErrUnprocessableEntity = "unprocessable_entity"
	ErrConflict            = "conflict"
	ErrValidationFailed    = "validation_failed"

	// "2-------------------------": "2-------------------------",
	// "---------2.General--------": "---------2.General--------",
	// "-------------------------2": "-------------------------2",

	MsgResourceCreated  = "resource_created"
	MsgResourceUpdated  = "resource_updated"
	MsgResourceDeleted  = "resource_deleted"
	MsgResourceFetched  = "resource_fetched"
	MsgResourceNotFound = "resource_not_found"

	ErrResourceCreated  = "resource_created"
	ErrResourceUpdated  = "err_resource_updated"
	ErrResourceDeleted  = "err_resource_deleted"
	ErrResourceFetched  = "err_resource_fetched"
	ErrResourceNotFound = "err_resource_not_found"
	ErrUnMatchedId      = "err_unmatched_id"

	// "3-------------------------": "3-------------------------",
	// "-------3.Resources--------": "-------3.Resources--------",
	// "-------------------------3": "-------------------------3",

	MsgUserResource         = "user_resource"
	MsgContentBlockResource = "content_block_resource"

	// "4-------------------------": "4-------------------------",
	// "-----------4.Auth---------": "-----------4.Auth---------",
	// "-------------------------4": "-------------------------4",

	MsgUserRegistered     = "user_registered"
	MsgTokenGenerated     = "token_generated"
	MsgInvalidCredentials = "invalid_credentials"
	MsgUserLoggedIn       = "user_logged_in"
	MsgLoginSuccessful    = "login_successful"
	MsgEmailVerified      = "email_verified"
)
