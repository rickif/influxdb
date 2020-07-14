package tenant

import (
	"fmt"
	"strings"

	"github.com/influxdata/influxdb/v2"
)

var (
	// ErrNameisEmpty is when a name is empty
	ErrNameisEmpty = &influxdb.Error{
		Code: influxdb.EInvalid,
		Msg:  "name is empty",
	}

	// NotUniqueIDError is used when attempting to create an org or bucket that already
	// exists.
	NotUniqueIDError = &influxdb.Error{
		Code: influxdb.EConflict,
		Msg:  "ID already exists",
	}

	// ErrFailureGeneratingID occurs ony when the random number generator
	// cannot generate an ID in MaxIDGenerationN times.
	ErrFailureGeneratingID = &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  "unable to generate valid id",
	}

	// ErrOnboardingNotAllowed occurs when request to onboard comes in and we are not allowing this request
	ErrOnboardingNotAllowed = &influxdb.Error{
		Code: influxdb.EConflict,
		Msg:  "onboarding has already been completed",
	}

	ErrOnboardInvalid = &influxdb.Error{
		Code: influxdb.EEmptyValue,
		Msg:  "onboard failed, missing value",
	}

	ErrNotFound = &influxdb.Error{
		Code: influxdb.ENotFound,
		Msg:  "not found",
	}
)

// ErrInternalServiceError is used when the error comes from an internal system.
func ErrInternalServiceError(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Err:  err,
	}
}

type errSlice []error

func (e errSlice) Error() string {
	l := len(e)
	sb := strings.Builder{}
	for i, err := range e {
		if i > 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(fmt.Sprintf("error %d/%d: %s", i+1, l, err.Error()))
	}
	return sb.String()
}

// AggregateError enables composing multiple errors.
// This is ideal in the case that you are applying functions with side effects to a slice of elements.
// E.g., deleting/updating a slice of resources.
type AggregateError struct {
	errs errSlice
}

// NewAggregateError returns a new AggregateError.
func NewAggregateError() *AggregateError {
	return &AggregateError{
		errs: make([]error, 0),
	}
}

// Add adds an error to the aggregate.
func (e *AggregateError) Add(err error) {
	if err == nil {
		return
	}
	e.errs = append(e.errs, err)
}

// Err returns a proper error from this aggregate error.
func (e *AggregateError) Err() error {
	if len(e.errs) > 0 {
		return e.errs
	}
	return nil
}

var (
	// ErrOrgNotFound is used when the user is not found.
	ErrOrgNotFound = &influxdb.Error{
		Msg:  "organization not found",
		Code: influxdb.ENotFound,
	}
)

// OrgAlreadyExistsError is used when creating a new organization with
// a name that has already been used. Organization names must be unique.
func OrgAlreadyExistsError(name string) error {
	return &influxdb.Error{
		Code: influxdb.EConflict,
		Msg:  fmt.Sprintf("organization with name %s already exists", name),
	}
}

func OrgNotFoundByName(name string) error {
	return &influxdb.Error{
		Code: influxdb.ENotFound,
		Op:   influxdb.OpFindOrganizations,
		Msg:  fmt.Sprintf("organization name \"%s\" not found", name),
	}
}

// ErrCorruptOrg is used when the user cannot be unmarshalled from the bytes
// stored in the kv.
func ErrCorruptOrg(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  "user could not be unmarshalled",
		Err:  err,
		Op:   "kv/UnmarshalOrg",
	}
}

// ErrUnprocessableOrg is used when a org is not able to be processed.
func ErrUnprocessableOrg(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EUnprocessableEntity,
		Msg:  "user could not be marshalled",
		Err:  err,
		Op:   "kv/MarshalOrg",
	}
}

// InvalidOrgIDError is used when a service was provided an invalid ID.
// This is some sort of internal server error.
func InvalidOrgIDError(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInvalid,
		Msg:  "org id provided is invalid",
		Err:  err,
	}
}

var (
	invalidBucketListRequest = &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  "invalid bucket list action, call should be GetBucketByName",
		Op:   "kv/listBucket",
	}

	errRenameSystemBucket = &influxdb.Error{
		Code: influxdb.EInvalid,
		Msg:  "system buckets cannot be renamed",
	}

	errDeleteSystemBucket = &influxdb.Error{
		Code: influxdb.EInvalid,
		Msg:  "system buckets cannot be deleted",
	}

	ErrBucketNotFound = &influxdb.Error{
		Code: influxdb.ENotFound,
		Msg:  "bucket not found",
	}

	ErrBucketNameNotUnique = &influxdb.Error{
		Code: influxdb.EConflict,
		Msg:  "bucket name is not unique",
	}
)

// ErrBucketNotFoundByName is used when the user is not found.
func ErrBucketNotFoundByName(n string) *influxdb.Error {
	return &influxdb.Error{
		Msg:  fmt.Sprintf("bucket %q not found", n),
		Code: influxdb.ENotFound,
	}
}

// ErrCorruptBucket is used when the user cannot be unmarshalled from the bytes
// stored in the kv.
func ErrCorruptBucket(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  "user could not be unmarshalled",
		Err:  err,
		Op:   "kv/UnmarshalBucket",
	}
}

// BucketAlreadyExistsError is used when attempting to create a user with a name
// that already exists.
func BucketAlreadyExistsError(n string) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EConflict,
		Msg:  fmt.Sprintf("bucket with name %s already exists", n),
	}
}

// ErrUnprocessableBucket is used when a org is not able to be processed.
func ErrUnprocessableBucket(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EUnprocessableEntity,
		Msg:  "user could not be marshalled",
		Err:  err,
		Op:   "kv/MarshalBucket",
	}
}
