# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [authn/types.proto](#authn_types-proto)
    - [AuthnResult](#authn-AuthnResult)
  
    - [AuthnAttemptStatus](#authn-AuthnAttemptStatus)
    - [AuthnStatus](#authn-AuthnStatus)
    - [AuthnType](#authn-AuthnType)
  
- [authn/authn.proto](#authn_authn-proto)
    - [CheckAuthnStatusRequest](#authn-CheckAuthnStatusRequest)
    - [CheckAuthnStatusResponse](#authn-CheckAuthnStatusResponse)
    - [LoginRequest](#authn-LoginRequest)
    - [LoginResponse](#authn-LoginResponse)
    - [LogoutRequest](#authn-LogoutRequest)
    - [LogoutResponse](#authn-LogoutResponse)
  
    - [AuthnService](#authn-AuthnService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="authn_types-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## authn/types.proto



<a name="authn-AuthnResult"></a>

### AuthnResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |
| uid | [uint32](#uint32) |  |  |
| session_id | [string](#string) | optional |  |
| remaining_pwd_attempts | [int32](#int32) | optional |  |
| remaining_mfa_attempts | [int32](#int32) | optional |  |
| user_locked | [bool](#bool) | optional |  |
| redirect_url | [string](#string) | optional |  |





 


<a name="authn-AuthnAttemptStatus"></a>

### AuthnAttemptStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHN_ATTEMPT_STATUS_UNSPECIFIED | 0 |  |
| SUCCESS | 1 |  |
| FAILED | 2 |  |
| BLOCKED | 3 | 拦截（风控等因素） |



<a name="authn-AuthnStatus"></a>

### AuthnStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHN_STATUS_UNSPECIFIED | 0 |  |
| VALID | 1 |  |
| INVALID | 2 |  |



<a name="authn-AuthnType"></a>

### AuthnType


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHN_TYPE_UNSPECIFIED | 0 |  |
| PASSWORD | 1 | 密码 |
| PASSWORD_MFA | 2 | 密码&#43;MFA |
| THIRD_PARTY | 3 | 第三方身份 |


 

 

 



<a name="authn_authn-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## authn/authn.proto



<a name="authn-CheckAuthnStatusRequest"></a>

### CheckAuthnStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| base | [common.Base](#common-Base) |  |  |






<a name="authn-CheckAuthnStatusResponse"></a>

### CheckAuthnStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [AuthnStatus](#authn-AuthnStatus) |  |  |
| uid | [uint32](#uint32) | optional | user id |
| sid | [string](#string) | optional | session id |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |






<a name="authn-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |
| client_ip | [string](#string) | optional |  |
| user_agent | [string](#string) | optional |  |
| did | [string](#string) | optional |  |
| client_id | [string](#string) | optional |  |
| base | [common.Base](#common-Base) |  |  |






<a name="authn-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [AuthnAttemptStatus](#authn-AuthnAttemptStatus) |  |  |
| token | [string](#string) | optional |  |
| user_locked | [bool](#bool) | optional |  |
| remaining_pwd_attempts | [int32](#int32) | optional |  |
| remaining_mfa_attempts | [int32](#int32) | optional |  |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |






<a name="authn-LogoutRequest"></a>

### LogoutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| client_id | [string](#string) | optional |  |
| base | [common.Base](#common-Base) |  |  |






<a name="authn-LogoutResponse"></a>

### LogoutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_resp | [common.BaseResp](#common-BaseResp) |  |  |





 

 

 


<a name="authn-AuthnService"></a>

### AuthnService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#authn-LoginRequest) | [LoginResponse](#authn-LoginResponse) | Login logs in a user with the given account and password. |
| Logout | [LogoutRequest](#authn-LogoutRequest) | [LogoutResponse](#authn-LogoutResponse) | Logout logs out a user with the given uid, session_id, and client_id. |
| CheckAuthnStatus | [CheckAuthnStatusRequest](#authn-CheckAuthnStatusRequest) | [CheckAuthnStatusResponse](#authn-CheckAuthnStatusResponse) | CheckAuthnStatus checks the authentication status of a user with the given token. |

 



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

