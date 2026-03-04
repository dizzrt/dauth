# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [identity/types.proto](#identity_types-proto)
    - [User](#identity-User)
    - [UserExtend](#identity-UserExtend)
  
    - [UserStatus](#identity-UserStatus)
  
- [identity/user.proto](#identity_user-proto)
    - [CreateUserRequest](#identity-CreateUserRequest)
    - [CreateUserResponse](#identity-CreateUserResponse)
    - [GetUserRequest](#identity-GetUserRequest)
    - [GetUserResponse](#identity-GetUserResponse)
    - [ListUsersRequest](#identity-ListUsersRequest)
    - [ListUsersResponse](#identity-ListUsersResponse)
    - [UpdateUserStatusRequest](#identity-UpdateUserStatusRequest)
    - [UpdateUserStatusResponse](#identity-UpdateUserStatusResponse)
  
    - [UserService](#identity-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="identity_types-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/types.proto



<a name="identity-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [uint32](#uint32) | optional |  |
| username | [string](#string) | optional |  |
| status | [UserStatus](#identity-UserStatus) | optional |  |
| phone | [string](#string) | optional |  |
| email | [string](#string) | optional |  |
| nickname | [string](#string) | optional |  |
| avatar | [string](#string) | optional |  |
| extend | [UserExtend](#identity-UserExtend) | optional |  |
| last_login_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |






<a name="identity-UserExtend"></a>

### UserExtend






 


<a name="identity-UserStatus"></a>

### UserStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_STATUS_UNSPECIFIED | 0 |  |
| ENABLED | 1 | 启用 |
| DISABLED | 2 | 禁用 |
| LOCKED | 3 | 锁定（密码/MFA验证失败触发） |


 

 

 



<a name="identity_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/user.proto



<a name="identity-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| password | [string](#string) |  |  |
| username | [string](#string) | optional |  |
| phone | [string](#string) | optional |  |
| email | [string](#string) | optional |  |
| nickname | [string](#string) | optional |  |
| status | [UserStatus](#identity-UserStatus) | optional |  |
| extend | [UserExtend](#identity-UserExtend) | optional |  |
| base | [common.Base](#common-Base) |  |  |






<a name="identity-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#identity-User) |  |  |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |






<a name="identity-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [uint32](#uint32) |  |  |
| base | [common.Base](#common-Base) |  |  |






<a name="identity-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#identity-User) |  |  |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |






<a name="identity-ListUsersRequest"></a>

### ListUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [common.Pagination](#common-Pagination) |  |  |
| base | [common.Base](#common-Base) |  |  |






<a name="identity-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#identity-User) | repeated |  |
| pagination | [common.Pagination](#common-Pagination) |  |  |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |






<a name="identity-UpdateUserStatusRequest"></a>

### UpdateUserStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [uint32](#uint32) |  |  |
| status | [UserStatus](#identity-UserStatus) |  |  |
| base | [common.Base](#common-Base) |  |  |






<a name="identity-UpdateUserStatusResponse"></a>

### UpdateUserStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [UserStatus](#identity-UserStatus) |  | new status |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |





 

 

 


<a name="identity-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#identity-CreateUserRequest) | [CreateUserResponse](#identity-CreateUserResponse) | CreateUser creates a new user. |
| GetUser | [GetUserRequest](#identity-GetUserRequest) | [GetUserResponse](#identity-GetUserResponse) | GetUser gets a user by ID. |
| ListUsers | [ListUsersRequest](#identity-ListUsersRequest) | [ListUsersResponse](#identity-ListUsersResponse) | ListUsers lists all users. |
| UpdateUserStatus | [UpdateUserStatusRequest](#identity-UpdateUserStatusRequest) | [UpdateUserStatusResponse](#identity-UpdateUserStatusResponse) | UpdateUserStatus updates a user&#39;s status. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

