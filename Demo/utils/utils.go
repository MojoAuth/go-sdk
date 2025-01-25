package utils

import (
	"encoding/json"
	"mojogodemo/customerrors"
	"net/http"
)

func WriteErrorJSON(r *http.Request, w http.ResponseWriter, cErr *customerrors.CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(cErr.HTTPStatusCode))

	byteArr, mErr := json.Marshal(cErr.ClientError)
	if mErr != nil {
		sErr := customerrors.ExtendCustomError(
			customerrors.List["SOMETHING_WENT_WRONG"],
			"WriteErrorJSON() : Error while marshaling.",
			mErr,
		)
		WriteErrorJSON(r, w, sErr)
	} else {
		_, mErr = w.Write(byteArr)
		if mErr != nil {
			sErr := customerrors.ExtendCustomError(
				customerrors.List["SOMETHING_WENT_WRONG"],
				"WriteErrorJSON() : Error while writing the response.",
				mErr,
			)
			WriteErrorJSON(r, w, sErr)
		}
	}
	return
}
func WriteResponseJSON(r *http.Request, w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	byteArr, mErr := json.Marshal(body)
	if mErr != nil {
		sErr := customerrors.ExtendCustomError(
			customerrors.List["SOMETHING_WENT_WRONG"],
			"WriteResponseJSON() : Error while marshaling.",
			mErr,
		)
		WriteErrorJSON(r, w, sErr)
	} else {
		_, mErr = w.Write(byteArr)
		if mErr != nil {
			sErr := customerrors.ExtendCustomError(
				customerrors.List["SOMETHING_WENT_WRONG"],
				"WriteResponseJSON() : Error while writing the response.",
				mErr,
			)
			WriteErrorJSON(r, w, sErr)
		}
	}

	return
}
