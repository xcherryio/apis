# AppDatabaseReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tables** | Pointer to [**[]AppDatabaseTableReadResponse**](AppDatabaseTableReadResponse.md) |  | [optional] 
**AppDBErrorType** | Pointer to [**ErrorSubType**](ErrorSubType.md) |  | [optional] 
**AppDBErrorCode** | Pointer to **string** | the error code from database driver | [optional] 
**AppDBErrorMessage** | Pointer to **string** | the error message from database driver | [optional] 
**AppDBErrorTableName** | Pointer to **string** | the first table that encounters the error to help SDK to throw the error in a friendly way  | [optional] 

## Methods

### NewAppDatabaseReadResponse

`func NewAppDatabaseReadResponse() *AppDatabaseReadResponse`

NewAppDatabaseReadResponse instantiates a new AppDatabaseReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseReadResponseWithDefaults

`func NewAppDatabaseReadResponseWithDefaults() *AppDatabaseReadResponse`

NewAppDatabaseReadResponseWithDefaults instantiates a new AppDatabaseReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTables

`func (o *AppDatabaseReadResponse) GetTables() []AppDatabaseTableReadResponse`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *AppDatabaseReadResponse) GetTablesOk() (*[]AppDatabaseTableReadResponse, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *AppDatabaseReadResponse) SetTables(v []AppDatabaseTableReadResponse)`

SetTables sets Tables field to given value.

### HasTables

`func (o *AppDatabaseReadResponse) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetAppDBErrorType

`func (o *AppDatabaseReadResponse) GetAppDBErrorType() ErrorSubType`

GetAppDBErrorType returns the AppDBErrorType field if non-nil, zero value otherwise.

### GetAppDBErrorTypeOk

`func (o *AppDatabaseReadResponse) GetAppDBErrorTypeOk() (*ErrorSubType, bool)`

GetAppDBErrorTypeOk returns a tuple with the AppDBErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorType

`func (o *AppDatabaseReadResponse) SetAppDBErrorType(v ErrorSubType)`

SetAppDBErrorType sets AppDBErrorType field to given value.

### HasAppDBErrorType

`func (o *AppDatabaseReadResponse) HasAppDBErrorType() bool`

HasAppDBErrorType returns a boolean if a field has been set.

### GetAppDBErrorCode

`func (o *AppDatabaseReadResponse) GetAppDBErrorCode() string`

GetAppDBErrorCode returns the AppDBErrorCode field if non-nil, zero value otherwise.

### GetAppDBErrorCodeOk

`func (o *AppDatabaseReadResponse) GetAppDBErrorCodeOk() (*string, bool)`

GetAppDBErrorCodeOk returns a tuple with the AppDBErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorCode

`func (o *AppDatabaseReadResponse) SetAppDBErrorCode(v string)`

SetAppDBErrorCode sets AppDBErrorCode field to given value.

### HasAppDBErrorCode

`func (o *AppDatabaseReadResponse) HasAppDBErrorCode() bool`

HasAppDBErrorCode returns a boolean if a field has been set.

### GetAppDBErrorMessage

`func (o *AppDatabaseReadResponse) GetAppDBErrorMessage() string`

GetAppDBErrorMessage returns the AppDBErrorMessage field if non-nil, zero value otherwise.

### GetAppDBErrorMessageOk

`func (o *AppDatabaseReadResponse) GetAppDBErrorMessageOk() (*string, bool)`

GetAppDBErrorMessageOk returns a tuple with the AppDBErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorMessage

`func (o *AppDatabaseReadResponse) SetAppDBErrorMessage(v string)`

SetAppDBErrorMessage sets AppDBErrorMessage field to given value.

### HasAppDBErrorMessage

`func (o *AppDatabaseReadResponse) HasAppDBErrorMessage() bool`

HasAppDBErrorMessage returns a boolean if a field has been set.

### GetAppDBErrorTableName

`func (o *AppDatabaseReadResponse) GetAppDBErrorTableName() string`

GetAppDBErrorTableName returns the AppDBErrorTableName field if non-nil, zero value otherwise.

### GetAppDBErrorTableNameOk

`func (o *AppDatabaseReadResponse) GetAppDBErrorTableNameOk() (*string, bool)`

GetAppDBErrorTableNameOk returns a tuple with the AppDBErrorTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDBErrorTableName

`func (o *AppDatabaseReadResponse) SetAppDBErrorTableName(v string)`

SetAppDBErrorTableName sets AppDBErrorTableName field to given value.

### HasAppDBErrorTableName

`func (o *AppDatabaseReadResponse) HasAppDBErrorTableName() bool`

HasAppDBErrorTableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


