# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [sp/sp_common.proto](#sp_sp_common-proto)
    - [ServiceProvider](#-ServiceProvider)
  
    - [ServiceProvider.Status](#-ServiceProvider-Status)
  
- [sp/sp.proto](#sp_sp-proto)
    - [CreateServiceProviderRequest](#sp-CreateServiceProviderRequest)
    - [CreateServiceProviderResponse](#sp-CreateServiceProviderResponse)
    - [GetServiceProviderRequest](#sp-GetServiceProviderRequest)
    - [GetServiceProviderResponse](#sp-GetServiceProviderResponse)
    - [ListServiceProviderRequest](#sp-ListServiceProviderRequest)
    - [ListServiceProviderResponse](#sp-ListServiceProviderResponse)
    - [ValidateServiceProviderRequest](#sp-ValidateServiceProviderRequest)
    - [ValidateServiceProviderResponse](#sp-ValidateServiceProviderResponse)
  
    - [ServiceProviderService](#sp-ServiceProviderService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="sp_sp_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## sp/sp_common.proto



<a name="-ServiceProvider"></a>

### ServiceProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |
| status | [ServiceProvider.Status](#ServiceProvider-Status) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |





 


<a name="-ServiceProvider-Status"></a>

### ServiceProvider.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| INACTIVE | 2 |  |
| DELETED | 3 |  |


 

 

 



<a name="sp_sp-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## sp/sp.proto



<a name="sp-CreateServiceProviderRequest"></a>

### CreateServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| secret | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |
| scopes | [uint32](#uint32) | repeated |  |
| base | [base.Base](#base-Base) |  |  |






<a name="sp-CreateServiceProviderResponse"></a>

### CreateServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sp_id | [uint32](#uint32) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="sp-GetServiceProviderRequest"></a>

### GetServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sp_id | [uint32](#uint32) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="sp-GetServiceProviderResponse"></a>

### GetServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sp | [ServiceProvider](#ServiceProvider) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="sp-ListServiceProviderRequest"></a>

### ListServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [uint32](#uint32) |  |  |
| page_size | [uint32](#uint32) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="sp-ListServiceProviderResponse"></a>

### ListServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sp_list | [ServiceProvider](#ServiceProvider) | repeated |  |
| total | [uint32](#uint32) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |






<a name="sp-ValidateServiceProviderRequest"></a>

### ValidateServiceProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sp_id | [uint32](#uint32) |  |  |
| scope | [string](#string) |  |  |
| base | [base.Base](#base-Base) |  |  |






<a name="sp-ValidateServiceProviderResponse"></a>

### ValidateServiceProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_ok | [bool](#bool) |  |  |
| reason | [string](#string) |  |  |
| base_resp | [base.BaseResp](#base-BaseResp) |  |  |





 

 

 


<a name="sp-ServiceProviderService"></a>

### ServiceProviderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateServiceProvider | [CreateServiceProviderRequest](#sp-CreateServiceProviderRequest) | [CreateServiceProviderResponse](#sp-CreateServiceProviderResponse) | CreateServiceProvider creates a new service provider. |
| GetServiceProvider | [GetServiceProviderRequest](#sp-GetServiceProviderRequest) | [GetServiceProviderResponse](#sp-GetServiceProviderResponse) | GetServiceProvider gets the service provider by id. |
| ListServiceProvider | [ListServiceProviderRequest](#sp-ListServiceProviderRequest) | [ListServiceProviderResponse](#sp-ListServiceProviderResponse) | ListServiceProvider lists the service providers. |
| ValidateServiceProvider | [ValidateServiceProviderRequest](#sp-ValidateServiceProviderRequest) | [ValidateServiceProviderResponse](#sp-ValidateServiceProviderResponse) | ValidateServiceProvider validates the service provider and scope. |

 



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

