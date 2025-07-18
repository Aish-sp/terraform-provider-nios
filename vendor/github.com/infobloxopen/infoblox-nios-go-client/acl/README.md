# Go API client for acl

OpenAPI specification for Infoblox NIOS WAPI ACL objects

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.13.6
- Package version: 1.0.0
- Generator version: 7.5.0
- Build package: com.infoblox.codegen.NiosGoClientCodegen
For more information, please visit [https://www.infoblox.com](https://www.infoblox.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import acl "github.com/infobloxopen/infoblox-nios-go-client/acl"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `acl.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), acl.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `acl.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), acl.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `acl.ContextOperationServerIndices` and `acl.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), acl.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), acl.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/wapi/v2.13.6*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*NamedaclAPI* | [**Create**](docs/NamedaclAPI.md#create) | **Post** /namedacl | Create a namedacl object
*NamedaclAPI* | [**Delete**](docs/NamedaclAPI.md#delete) | **Delete** /namedacl/{reference} | Delete a namedacl object
*NamedaclAPI* | [**List**](docs/NamedaclAPI.md#list) | **Get** /namedacl | Retrieve namedacl objects
*NamedaclAPI* | [**Read**](docs/NamedaclAPI.md#read) | **Get** /namedacl/{reference} | Get a specific namedacl object
*NamedaclAPI* | [**Update**](docs/NamedaclAPI.md#update) | **Put** /namedacl/{reference} | Update a namedacl object


## Documentation For Models

 - [CreateNamedaclResponse](docs/CreateNamedaclResponse.md)
 - [CreateNamedaclResponseAsObject](docs/CreateNamedaclResponseAsObject.md)
 - [ExtAttrs](docs/ExtAttrs.md)
 - [GetNamedaclResponse](docs/GetNamedaclResponse.md)
 - [GetNamedaclResponseObjectAsResult](docs/GetNamedaclResponseObjectAsResult.md)
 - [ListNamedaclResponse](docs/ListNamedaclResponse.md)
 - [ListNamedaclResponseObject](docs/ListNamedaclResponseObject.md)
 - [Namedacl](docs/Namedacl.md)
 - [NamedaclAccessList](docs/NamedaclAccessList.md)
 - [NamedaclExplodedAccessList](docs/NamedaclExplodedAccessList.md)
 - [UpdateNamedaclResponse](docs/UpdateNamedaclResponse.md)
 - [UpdateNamedaclResponseAsObject](docs/UpdateNamedaclResponseAsObject.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), acl.ContextBasicAuth, acl.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



