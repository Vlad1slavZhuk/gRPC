package constErr

import (
	"errors"
)

var (
	NoValidAcc                  = errors.New("No valid account.")
	NoSuchAccountInTheDatabase  = errors.New("No such account in the database.")
	FailedToGenerateAToken      = errors.New("Failed to generate a token.")
	NotAuthorizedOrEmptyToken   = errors.New("Not authorized or hasn't set a token.")
	TokenIsDeprecated           = errors.New("Token is deprecated.")
	InvalidID                   = errors.New("Invalid ID.")
	UnsuccessfulMarshal         = errors.New("Unsuccessful marshal.")
	AdBaseIsEmpty               = errors.New("Ad base is empty.")
	ErrorAddingAds              = errors.New("Error adding ads.")
	UnsuccessfulDeleteOperation = errors.New("Unsuccessful delete operation.")
	EmptyFields                 = errors.New("Empty fields.")
	FailedDataUpdate            = errors.New("Failed data update.")
	DecodeError                 = errors.New("Decode error.")
	AccountBaseIsEmpty          = errors.New("Base of accounts is empty.")
	TokenNotContain             = errors.New("Token not contain.")
	NoValidToken                = errors.New("No valid token.")
	YouRat                      = errors.New("...Ля ты крыса! (Comedy club).")
	ErrorToAddAcc               = errors.New("Error to add account to base.")
	ErrorDelete                 = errors.New("Error to delete.")
	ErrorUpdate                 = errors.New("Error to Update")
	NotFoundAd                  = errors.New("No Found ad.")
	AdIsNil                     = errors.New("Ad is nil.")
	NotFoundAcc                 = errors.New("Not found Account.")
	AccIsNil                    = errors.New("Account is nil.")
)
