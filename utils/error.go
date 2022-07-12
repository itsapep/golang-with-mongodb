package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    string
	Message string
	ErrType int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type:%d, code:%s, err:%s", e.ErrType, e.Code, e.Message)
}

func NewCastInterfaceError(model string) AppError {
	return AppError{
		Code:    "ERR_CAST_INTERFACE",
		Message: "Cannot cast interface to " + model + " model",
	}
}

func NewNotFoundError(props interface{}) AppError {
	return AppError{
		Code:    "ERR_NOT_FOUND",
		Message: fmt.Sprintf("Data %v not found", props),
	}
}

func NewInvalidTransTypeError(id string) AppError {
	return AppError{
		Code:    "ERR_INVALID_TRANS_TYPE_ID",
		Message: "Cannot use " + id + " as trans type id",
	}
}

func NewCustomerAlreadyMemberError(id uint) AppError {
	return AppError{
		Code:    "ERR_CUSTOMER_ALREADY_MEMBER",
		Message: fmt.Sprintf("Customer with id %d is already a member", id),
	}
}

func NewTableUnavailableError(id uint) AppError {
	return AppError{
		Code:    "ERR_TABLE_UNAVAILABLE",
		Message: fmt.Sprintf("Table %d is not available", id),
	}
}

func NewInvalidMenuToOrderError(menuPriceIds []uint) AppError {
	return AppError{
		Code:    "ERR_INVALID_MENU_TO_ORDER",
		Message: fmt.Sprintf("Menu with price ID %v are not exist", menuPriceIds),
	}
}

func RequiredError() error {
	return AppError{
		Code:    "ERR_EMPTY_FIELD",
		Message: "Input cant be empty",
		ErrType: http.StatusBadRequest,
	}
}

func Unauthorized() error {
	return AppError{
		Code:    "ERR-UNAUTHORISED",
		Message: "Unauthorised user",
		ErrType: http.StatusUnauthorized,
	}
}
