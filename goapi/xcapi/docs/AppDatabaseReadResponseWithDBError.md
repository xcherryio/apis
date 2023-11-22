# AppDatabaseReadResponseWithDBError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tables** | Pointer to [**[]AppDatabaseTableReadResponse**](AppDatabaseTableReadResponse.md) |  | [optional] 
**AppDBErrorType** | Pointer to [**ErrorSubType**](ErrorSubType.md) |  | [optional] 
**AppDBErrorCode** | Pointer to **string** | the error code from database driver | [optional] 
**AppDBErrorMessage** | Pointer to **string** | the error message from database driver | [optional] 
**AppDBErrorTableName** | Pointer to **string** | the first table that encounters the error to help SDK to throw the error in a friendly way  | [optional] 

## Methods

### NewAppDatabaseReadResponseWithDBError

`func NewAppDatabaseReadResponseWithDBError() *AppDatabaseReadResponseWithDBError`

NewAppDatabaseReadResponseWithDBError instantiates a new AppDatabaseReadResponseWithDBError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseReadResponseWithDBErrorWithDefaults

`func NewAppDatabaseReadResponseWithDBErrorWithDefaults() *AppDatabaseReadResponseWithDBError`

NewAppDatabaseReadResponseWithDBErrorWithDefaults instantiates a new AppDatabaseReadResponseWithDBError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTables

`func (o *AppDatabaseReadResponseWithDBError) GetTables() []AppDatabaseTableReadResponse`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *AppDatabaseReadResponseWithDBError) GetTablesOk() (*[]AppDatabaseTableReadResponse, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *AppDatabaseReadResponseWithDBError) SetTables(v []AppDatabaseTableReadResponse)`

SetTables sets Tables field to given value.

### HasTables

`func (o *AppDatabaseReadResponseWithDBError) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetAppDBErrorType

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorType() ErrorSubType`

GetAppDBErrorType returns the AppDBErrorType field if non-nil, zero value otherwise.

### GetAppDBErrorTypeOk

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorTypeOk() (*ErrorSubType, bool)`

GetAppDBErrorTypeOk returns a tuple with the AppDBErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorType

`func (o *AppDatabaseReadResponseWithDBError) SetAppDBErrorType(v ErrorSubType)`

SetAppDBErrorType sets AppDBErrorType field to given value.

### HasAppDBErrorType

`func (o *AppDatabaseReadResponseWithDBError) HasAppDBErrorType() bool`

HasAppDBErrorType returns a boolean if a field has been set.

### GetAppDBErrorCode

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorCode() string`

GetAppDBErrorCode returns the AppDBErrorCode field if non-nil, zero value otherwise.

### GetAppDBErrorCodeOk

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorCodeOk() (*string, bool)`

GetAppDBErrorCodeOk returns a tuple with the AppDBErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorCode

`func (o *AppDatabaseReadResponseWithDBError) SetAppDBErrorCode(v string)`

SetAppDBErrorCode sets AppDBErrorCode field to given value.

### HasAppDBErrorCode

`func (o *AppDatabaseReadResponseWithDBError) HasAppDBErrorCode() bool`

HasAppDBErrorCode returns a boolean if a field has been set.

### GetAppDBErrorMessage

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorMessage() string`

GetAppDBErrorMessage returns the AppDBErrorMessage field if non-nil, zero value otherwise.

### GetAppDBErrorMessageOk

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorMessageOk() (*string, bool)`

GetAppDBErrorMessageOk returns a tuple with the AppDBErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorMessage

`func (o *AppDatabaseReadResponseWithDBError) SetAppDBErrorMessage(v string)`

SetAppDBErrorMessage sets AppDBErrorMessage field to given value.

### HasAppDBErrorMessage

`func (o *AppDatabaseReadResponseWithDBError) HasAppDBErrorMessage() bool`

HasAppDBErrorMessage returns a boolean if a field has been set.

### GetAppDBErrorTableName

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorTableName() string`

GetAppDBErrorTableName returns the AppDBErrorTableName field if non-nil, zero value otherwise.

### GetAppDBErrorTableNameOk

`func (o *AppDatabaseReadResponseWithDBError) GetAppDBErrorTableNameOk() (*string, bool)`

GetAppDBErrorTableNameOk returns a tuple with the AppDBErrorTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorTableName

`func (o *AppDatabaseReadResponseWithDBError) SetAppDBErrorTableName(v string)`

SetAppDBErrorTableName sets AppDBErrorTableName field to given value.

### HasAppDBErrorTableName

`func (o *AppDatabaseReadResponseWithDBError) HasAppDBErrorTableName() bool`

HasAppDBErrorTableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


