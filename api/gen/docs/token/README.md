# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [token/token_common.proto](#token_token_common-proto)
    - [Token](#token-Token)
  
    - [Token.TokenType](#token-Token-TokenType)
  
- [token/token.proto](#token_token-proto)
    - [IssueRequest](#token-IssueRequest)
    - [IssueResponse](#token-IssueResponse)
    - [IssueSSOTokenRequest](#token-IssueSSOTokenRequest)
    - [IssueSSOTokenResponse](#token-IssueSSOTokenResponse)
    - [RevokeRequest](#token-RevokeRequest)
    - [RevokeResponse](#token-RevokeResponse)
    - [ValidateRequest](#token-ValidateRequest)
    - [ValidateResponse](#token-ValidateResponse)
  
    - [TokenService](#token-TokenService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="token_token_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/token_common.proto



<a name="token-Token"></a>

### Token
map to domain/token/entity/token.go BaseToken


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| issuer | [string](#string) |  |  |
| subject | [string](#string) |  |  |
| audience | [string](#string) | repeated |  |
| issued_at | [int64](#int64) |  |  |
| not_before | [int64](#int64) |  |  |
| expires_at | [int64](#int64) |  |  |
| uid | [uint32](#uint32) |  |  |
| type | [Token.TokenType](#token-Token-TokenType) |  |  |





 


<a name="token-Token-TokenType"></a>

### Token.TokenType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TokenType_UNKNOWN | 0 |  |
| TokenType_SSO | 1 |  |
| TokenType_ID | 2 |  |
| TokenType_ACCESS | 3 |  |
| TokenType_REFRESH | 4 |  |


 

 

 



<a name="token_token-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/token.proto



<a name="token-IssueRequest"></a>

### IssueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [uint32](#uint32) |  |  |
| client_id | [uint32](#uint32) |  |  |
| scope | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="token-IssueResponse"></a>

### IssueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |
| access_expire_at | [int64](#int64) |  |  |
| refresh_expire_at | [int64](#int64) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="token-IssueSSOTokenRequest"></a>

### IssueSSOTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [uint32](#uint32) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="token-IssueSSOTokenResponse"></a>

### IssueSSOTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| expires_at | [int64](#int64) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="token-RevokeRequest"></a>

### RevokeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="token-RevokeResponse"></a>

### RevokeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_success | [bool](#bool) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="token-ValidateRequest"></a>

### ValidateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| type | [Token.TokenType](#token-Token-TokenType) |  |  |
| client_id | [uint32](#uint32) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="token-ValidateResponse"></a>

### ValidateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#token-Token) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="token-TokenService"></a>

### TokenService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IssueSSOToken | [IssueSSOTokenRequest](#token-IssueSSOTokenRequest) | [IssueSSOTokenResponse](#token-IssueSSOTokenResponse) | IssueSSOToken issues a SSO token for the given uid. |
| Issue | [IssueRequest](#token-IssueRequest) | [IssueResponse](#token-IssueResponse) | Issue issues a token for the given uid and client_id. |
| Validate | [ValidateRequest](#token-ValidateRequest) | [ValidateResponse](#token-ValidateResponse) | Validate validates a token. |
| Revoke | [RevokeRequest](#token-RevokeRequest) | [RevokeResponse](#token-RevokeResponse) | Revoke revokes a token. |

 



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

