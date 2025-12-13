# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [identity/identity_common.proto](#identity_identity_common-proto)
    - [Role](#-Role)
    - [User](#-User)
  
    - [Role.Status](#-Role-Status)
    - [User.Status](#-User-Status)
  
- [identity/role.proto](#identity_role-proto)
    - [AssignRolesRequest](#identity-AssignRolesRequest)
    - [AssignRolesResponse](#identity-AssignRolesResponse)
    - [CreateRoleRequest](#identity-CreateRoleRequest)
    - [CreateRoleResponse](#identity-CreateRoleResponse)
    - [DeleteRolesRequest](#identity-DeleteRolesRequest)
    - [DeleteRolesResponse](#identity-DeleteRolesResponse)
    - [GetRolesRequest](#identity-GetRolesRequest)
    - [GetRolesResponse](#identity-GetRolesResponse)
    - [UnassignRolesRequest](#identity-UnassignRolesRequest)
    - [UnassignRolesResponse](#identity-UnassignRolesResponse)
    - [UpdateRoleRequest](#identity-UpdateRoleRequest)
    - [UpdateRoleResponse](#identity-UpdateRoleResponse)
  
    - [RoleService](#identity-RoleService)
  
- [identity/user.proto](#identity_user-proto)
    - [AuthenticateRequest](#identity-AuthenticateRequest)
    - [AuthenticateResponse](#identity-AuthenticateResponse)
    - [CreateUserRequest](#identity-CreateUserRequest)
    - [CreateUserResponse](#identity-CreateUserResponse)
    - [GetUserRequest](#identity-GetUserRequest)
    - [GetUserResponse](#identity-GetUserResponse)
    - [LoginRequest](#identity-LoginRequest)
    - [LoginResponse](#identity-LoginResponse)
    - [UpdateUserPasswordRequest](#identity-UpdateUserPasswordRequest)
    - [UpdateUserPasswordResponse](#identity-UpdateUserPasswordResponse)
    - [UpdateUserStatusRequest](#identity-UpdateUserStatusRequest)
    - [UpdateUserStatusResponse](#identity-UpdateUserStatusResponse)
  
    - [UserService](#identity-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="identity_identity_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/identity_common.proto



<a name="-Role"></a>

### Role



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| status | [Role.Status](#Role-Status) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |






<a name="-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| email | [string](#string) |  |  |
| username | [string](#string) |  |  |
| status | [User.Status](#User-Status) |  |  |
| roles | [Role](#Role) | repeated |  |
| last_login_at | [int64](#int64) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |





 


<a name="-Role-Status"></a>

### Role.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| INACTIVE | 2 |  |
| DELETED | 3 |  |



<a name="-User-Status"></a>

### User.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| INACTIVE | 2 |  |
| DELETED | 3 |  |


 

 

 



<a name="identity_role-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/role.proto



<a name="identity-AssignRolesRequest"></a>

### AssignRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| roles | [Role](#Role) | repeated |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-AssignRolesResponse"></a>

### AssignRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-CreateRoleResponse"></a>

### CreateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [Role](#Role) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-DeleteRolesRequest"></a>

### DeleteRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [uint32](#uint32) | repeated |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-DeleteRolesResponse"></a>

### DeleteRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-GetRolesRequest"></a>

### GetRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [Role](#Role) | repeated |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-UnassignRolesRequest"></a>

### UnassignRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| roles | [Role](#Role) | repeated |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-UnassignRolesResponse"></a>

### UnassignRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-UpdateRoleResponse"></a>

### UpdateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [Role](#Role) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="identity-RoleService"></a>

### RoleService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRole | [CreateRoleRequest](#identity-CreateRoleRequest) | [CreateRoleResponse](#identity-CreateRoleResponse) | CreateRole creates a new role. |
| GetRoles | [GetRolesRequest](#identity-GetRolesRequest) | [GetRolesResponse](#identity-GetRolesResponse) | GetRoles gets all roles. |
| DeleteRoles | [DeleteRolesRequest](#identity-DeleteRolesRequest) | [DeleteRolesResponse](#identity-DeleteRolesResponse) | DeleteRoles deletes roles by IDs. |
| UpdateRole | [UpdateRoleRequest](#identity-UpdateRoleRequest) | [UpdateRoleResponse](#identity-UpdateRoleResponse) | UpdateRole updates a role by ID. |
| AssignRoles | [AssignRolesRequest](#identity-AssignRolesRequest) | [AssignRolesResponse](#identity-AssignRolesResponse) | AssignRoles assigns roles to a user. |
| UnassignRoles | [UnassignRolesRequest](#identity-UnassignRolesRequest) | [UnassignRolesResponse](#identity-UnassignRolesResponse) | UnassignRoles unassigns roles from a user. |

 



<a name="identity_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/user.proto



<a name="identity-AuthenticateRequest"></a>

### AuthenticateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-AuthenticateResponse"></a>

### AuthenticateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  |  |
| token | [string](#string) |  |  |
| token_expires_at | [int64](#int64) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-UpdateUserPasswordRequest"></a>

### UpdateUserPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| password | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-UpdateUserPasswordResponse"></a>

### UpdateUserPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="identity-UpdateUserStatusRequest"></a>

### UpdateUserStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| status | [User.Status](#User-Status) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="identity-UpdateUserStatusResponse"></a>

### UpdateUserStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="identity-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#identity-LoginRequest) | [LoginResponse](#identity-LoginResponse) | Login logs in a user. |
| Authenticate | [AuthenticateRequest](#identity-AuthenticateRequest) | [AuthenticateResponse](#identity-AuthenticateResponse) | Authenticate authenticates a user. |
| CreateUser | [CreateUserRequest](#identity-CreateUserRequest) | [CreateUserResponse](#identity-CreateUserResponse) | CreateUser creates a new user. |
| GetUser | [GetUserRequest](#identity-GetUserRequest) | [GetUserResponse](#identity-GetUserResponse) | GetUser gets a user by ID. |
| UpdateUserStatus | [UpdateUserStatusRequest](#identity-UpdateUserStatusRequest) | [UpdateUserStatusResponse](#identity-UpdateUserStatusResponse) | UpdateUserStatus updates the status of a user. |
| UpdateUserPassword | [UpdateUserPasswordRequest](#identity-UpdateUserPasswordRequest) | [UpdateUserPasswordResponse](#identity-UpdateUserPasswordResponse) | UpdateUserPassword updates the password of a user. |

 



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

