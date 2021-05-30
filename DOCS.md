# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [urlshortener/service.proto](#urlshortener/service.proto)
    - [Key](#proto.Key)
    - [SetShortenedUrlInput](#proto.SetShortenedUrlInput)
    - [Url](#proto.Url)
    - [Void](#proto.Void)
  
    - [UrlShortenerService](#proto.UrlShortenerService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="urlshortener/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## urlshortener/service.proto



<a name="proto.Key"></a>

### Key
Represents a shortened URL key (a sequence of n length of random letters)
Is used as suffix for shortened URLs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  | Represents the key of a shortened URL (EX. localhost:40/&lt;lkjhbalsjdhbl&gt;, key between &lt;&gt;) |






<a name="proto.SetShortenedUrlInput"></a>

### SetShortenedUrlInput
input used for SetShortenedUrl rpc method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Url | [string](#string) |  | URL to be shortened |
| ExpiryMinutes | [int64](#int64) |  | Expiry time in minutes for entry If a 0 value is given, the entry will be persistent. |






<a name="proto.Url"></a>

### Url
Represents an accessible URL.
Used as either input or output of rpc calls which receive/return URLs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Url | [string](#string) |  | Full URL (Ex. https://facebook.com) |






<a name="proto.Void"></a>

### Void
Represents an empty return





 

 

 


<a name="proto.UrlShortenerService"></a>

### UrlShortenerService
UrlShortenerService represents the service for handling url shortening (CRUD).
Service uses Redis to cache and store results.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetShortenedUrl | [Key](#proto.Key) | [Url](#proto.Url) | Gets an already existing shortened URL by it&#39;s key (EX. localhost:40/&lt;lkjhbalsjdhbl&gt;, key between &lt;&gt;) |
| SetShortenedUrl | [SetShortenedUrlInput](#proto.SetShortenedUrlInput) | [Url](#proto.Url) | Creates a new shortened URL and returns it as an accessible path |
| DeleteShortenedUrl | [Key](#proto.Key) | [Void](#proto.Void) | Deletes an already existing shortened URL by it&#39;s key (EX. localhost:40/&lt;lkjhbalsjdhbl&gt;, key between &lt;&gt;) |

 



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

