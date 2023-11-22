# AppDatabaseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDBErrorType** | [**ErrorSubType**](ErrorSubType.md) |  | 
**AppDBErrorCode** | **string** | the error code from database driver | 
**AppDBErrorMessage** | Pointer to **string** | the error message from database driver | [optional] 
**AppDBErrorTableName** | Pointer to **string** | the first table that encounters the error to help SDK to throw the error in a friendly way  | [optional] 

## Methods

### NewAppDatabaseError

`func NewAppDatabaseError(appDBErrorType ErrorSubType, appDBErrorCode string, ) *AppDatabaseError`

NewAppDatabaseError instantiates a new AppDatabaseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseErrorWithDefaults

`func NewAppDatabaseErrorWithDefaults() *AppDatabaseError`

NewAppDatabaseErrorWithDefaults instantiates a new AppDatabaseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDBErrorType

`func (o *AppDatabaseError) GetAppDBErrorType() ErrorSubType`

GetAppDBErrorType returns the AppDBErrorType field if non-nil, zero value otherwise.

### GetAppDBErrorTypeOk

`func (o *AppDatabaseError) GetAppDBErrorTypeOk() (*ErrorSubType, bool)`

GetAppDBErrorTypeOk returns a tuple with the AppDBErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorType

`func (o *AppDatabaseError) SetAppDBErrorType(v ErrorSubType)`

SetAppDBErrorType sets AppDBErrorType field to given value.


### GetAppDBErrorCode

`func (o *AppDatabaseError) GetAppDBErrorCode() string`

GetAppDBErrorCode returns the AppDBErrorCode field if non-nil, zero value otherwise.

### GetAppDBErrorCodeOk

`func (o *AppDatabaseError) GetAppDBErrorCodeOk() (*string, bool)`

GetAppDBErrorCodeOk returns a tuple with the AppDBErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorCode

`func (o *AppDatabaseError) SetAppDBErrorCode(v string)`

SetAppDBErrorCode sets AppDBErrorCode field to given value.


### GetAppDBErrorMessage

`func (o *AppDatabaseError) GetAppDBErrorMessage() string`

GetAppDBErrorMessage returns the AppDBErrorMessage field if non-nil, zero value otherwise.

### GetAppDBErrorMessageOk

`func (o *AppDatabaseError) GetAppDBErrorMessageOk() (*string, bool)`

GetAppDBErrorMessageOk returns a tuple with the AppDBErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorMessage

`func (o *AppDatabaseError) SetAppDBErrorMessage(v string)`

SetAppDBErrorMessage sets AppDBErrorMessage field to given value.

### HasAppDBErrorMessage

`func (o *AppDatabaseError) HasAppDBErrorMessage() bool`

HasAppDBErrorMessage returns a boolean if a field has been set.

### GetAppDBErrorTableName

`func (o *AppDatabaseError) GetAppDBErrorTableName() string`

GetAppDBErrorTableName returns the AppDBErrorTableName field if non-nil, zero value otherwise.

### GetAppDBErrorTableNameOk

`func (o *AppDatabaseError) GetAppDBErrorTableNameOk() (*string, bool)`

GetAppDBErrorTableNameOk returns a tuple with the AppDBErrorTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorTableName

`func (o *AppDatabaseError) SetAppDBErrorTableName(v string)`

SetAppDBErrorTableName sets AppDBErrorTableName field to given value.

### HasAppDBErrorTableName

`func (o *AppDatabaseError) HasAppDBErrorTableName() bool`

HasAppDBErrorTableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


