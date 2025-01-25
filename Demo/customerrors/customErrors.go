package customerrors

import (
	"net/http"
)

type ClientError struct {
	ErrorCode uint16 `json:"code"`
	Message   string `json:"message"`
	Error     string `json:"description"`
}

type CustomError struct {
	ClientError    *ClientError
	HTTPStatusCode uint16
	Reason         string
	Err            error
}

var List map[string]*CustomError = map[string]*CustomError{
	"SOMETHING_WENT_WRONG": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 1000,
			Message:   "Oops, something is going wrong, please try again.",
			Error:     "Oops, something is going wrong, please try again.",
		},
		HTTPStatusCode: http.StatusInternalServerError,
	},
	"APIKeyNotWellFormatted": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 905,
			Message:   "API key is invalid",
			Error:     "The provided API key is invalid, please use a valid API key of your MojoAuth account.",
		},
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"APIKeyRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 904,
			Message:   "API key is required",
			Error:     "API key is required, please use a valid API key in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"TokenisRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 907,
			Message:   "Token is required",
			Error:     "Token key is required, please use a valid API key in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"TokenisInvalid": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 906,
			Message:   "Token is invalid",
			Error:     "Token is invalid, please use a valid token in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ProjectExist": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 901,
			Message:   "Project name already exist",
			Error:     "Project name already exist",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ProjectNameRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 900,
			Message:   "Project name is required",
			Error:     "Project name is required, please use a valid project name in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"APISecretRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 906,
			Message:   "API Secret is required",
			Error:     "API Secret key is required, please use a valid Secret key in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"APISecretInvalid": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 907,
			Message:   "API Secret is invalid",
			Error:     "The provided API Secret key is invalid, please use a valid secret key of your MojoAuth account.",
		},
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"ProjectIDInvalid": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 902,
			Message:   "Project ID is invalid",
			Error:     "Project ID is invalid, please use a valid project id in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ProjectNotExist": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 911,
			Message:   "Project does not exist",
			Error:     "Project does not exist",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ProjectIDRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 910,
			Message:   "Project ID is required",
			Error:     "Project ID is required, please use a valid project id in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"UserNotFound": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 910,
			Message:   "User not found",
			Error:     "User not found",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"CustomerIDRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 910,
			Message:   "Customer ID is required",
			Error:     "Customer ID is required, please use a valid customer id in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
	"DocumentNotAvailable": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 905,
			Message:   "No Record Found",
			Error:     "Document not Available.",
		},
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"EnvRequired": &CustomError{
		ClientError: &ClientError{
			ErrorCode: 912,
			Message:   "Environment value is required",
			Error:     "Environment value is required, please use a valid Environment value in order to process this request.",
		},
		HTTPStatusCode: http.StatusBadRequest,
	},
}

// ExtendCustomError extends defined err with reason and actual err
func ExtendCustomError(ce *CustomError, reason string, actualErr error) (err *CustomError) {
	customErr := *ce

	err = &CustomError{
		ClientError:    customErr.ClientError,
		HTTPStatusCode: customErr.HTTPStatusCode,
		Reason:         reason,
		Err:            actualErr,
	}
	return
}
