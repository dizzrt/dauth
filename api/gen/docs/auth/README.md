# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/auht_common.proto](#auth_auht_common-proto)
    - [AuthorizationCode](#-AuthorizationCode)
  
- [auth/auth.proto](#auth_auth-proto)
    - [ExchangeTokenRequest](#auth-ExchangeTokenRequest)
    - [ExchangeTokenResponse](#auth-ExchangeTokenResponse)
    - [GenerateAuthorizationCodeRequest](#auth-GenerateAuthorizationCodeRequest)
    - [GenerateAuthorizationCodeResponse](#auth-GenerateAuthorizationCodeResponse)
  
    - [AuthService](#auth-AuthService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auth_auht_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/auht_common.proto



<a name="-AuthorizationCode"></a>

### AuthorizationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| code | [string](#string) |  |  |
| user_id | [uint32](#uint32) |  |  |
| client_id | [uint32](#uint32) |  |  |
| redirect_uri | [string](#string) |  |  |
| scope | [string](#string) |  |  |
| issue_at | [int64](#int64) |  |  |
| expires_at | [int64](#int64) |  |  |
| used | [bool](#bool) |  |  |





 

 

 

 



<a name="auth_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/auth.proto



<a name="auth-ExchangeTokenRequest"></a>

### ExchangeTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| client_id | [uint32](#uint32) |  |  |
| client_secret | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="auth-ExchangeTokenResponse"></a>

### ExchangeTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |
| access_expire_at | [int64](#int64) |  |  |
| refresh_expire_at | [int64](#int64) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="auth-GenerateAuthorizationCodeRequest"></a>

### GenerateAuthorizationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| client_id | [uint32](#uint32) |  |  |
| redirect_uri | [string](#string) |  |  |
| scope | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="auth-GenerateAuthorizationCodeResponse"></a>

### GenerateAuthorizationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="auth-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateAuthorizationCode | [GenerateAuthorizationCodeRequest](#auth-GenerateAuthorizationCodeRequest) | [GenerateAuthorizationCodeResponse](#auth-GenerateAuthorizationCodeResponse) | GenerateAuthorizationCode generates an authorization code. |
| ExchangeToken | [ExchangeTokenRequest](#auth-ExchangeTokenRequest) | [ExchangeTokenResponse](#auth-ExchangeTokenResponse) | ExchangeToken exchanges an authorization code for an access token. |

 



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

