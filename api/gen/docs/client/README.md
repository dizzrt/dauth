# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [client/client_common.proto](#client_client_common-proto)
    - [Client](#-Client)
  
    - [Client.Status](#-Client-Status)
  
- [client/client.proto](#client_client-proto)
    - [CreateClientRequest](#client-CreateClientRequest)
    - [CreateClientResponse](#client-CreateClientResponse)
    - [ValidateClientRequest](#client-ValidateClientRequest)
    - [ValidateClientResponse](#client-ValidateClientResponse)
  
    - [ClientService](#client-ClientService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="client_client_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## client/client_common.proto



<a name="-Client"></a>

### Client



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |
| status | [Client.Status](#Client-Status) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |





 


<a name="-Client-Status"></a>

### Client.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| INACTIVE | 2 |  |
| DELETED | 3 |  |


 

 

 



<a name="client_client-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## client/client.proto



<a name="client-CreateClientRequest"></a>

### CreateClientRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| secret | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |
| scopes | [uint32](#uint32) | repeated |  |
| base | [base.Base](#base-Base) |  |  |






<a name="client-CreateClientResponse"></a>

### CreateClientResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [uint32](#uint32) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="client-ValidateClientRequest"></a>

### ValidateClientRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [uint32](#uint32) |  |  |
| scope | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="client-ValidateClientResponse"></a>

### ValidateClientResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_ok | [bool](#bool) |  |  |
| reason | [string](#string) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="client-ClientService"></a>

### ClientService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateClient | [CreateClientRequest](#client-CreateClientRequest) | [CreateClientResponse](#client-CreateClientResponse) | CreateClient creates a new client. |
| ValidateClient | [ValidateClientRequest](#client-ValidateClientRequest) | [ValidateClientResponse](#client-ValidateClientResponse) | ValidateClient validates the client and scope. |

 



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

