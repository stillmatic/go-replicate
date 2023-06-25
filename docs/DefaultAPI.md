# \DefaultAPI

All URIs are relative to *https://api.replicate.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionsGet**](DefaultAPI.md#CollectionsGet) | **Get** /v1/collections/{collection_slug} | Get a collection of models
[**CollectionsList**](DefaultAPI.md#CollectionsList) | **Get** /v1/collections | List collections of models
[**ModelsGet**](DefaultAPI.md#ModelsGet) | **Get** /v1/models/{model_owner}/{model_name} | Get a model
[**ModelsVersionsDelete**](DefaultAPI.md#ModelsVersionsDelete) | **Delete** /v1/models/{model_owner}/{model_name}/versions/{version_id} | Delete a model version
[**ModelsVersionsGet**](DefaultAPI.md#ModelsVersionsGet) | **Get** /v1/models/{model_owner}/{model_name}/versions/{version_id} | Get a model version
[**ModelsVersionsList**](DefaultAPI.md#ModelsVersionsList) | **Get** /v1/models/{model_owner}/{model_name}/versions | List model versions
[**PredictionsCancel**](DefaultAPI.md#PredictionsCancel) | **Post** /v1/predictions/{prediction_id}/cancel | Cancel a prediction
[**PredictionsCreate**](DefaultAPI.md#PredictionsCreate) | **Post** /v1/predictions | Create a prediction
[**PredictionsGet**](DefaultAPI.md#PredictionsGet) | **Get** /v1/predictions/{prediction_id} | Get a prediction
[**PredictionsList**](DefaultAPI.md#PredictionsList) | **Get** /v1/predictions | List predictions
[**TrainingsCancel**](DefaultAPI.md#TrainingsCancel) | **Post** /v1/trainings/{training_id}/cancel | Cancel a training
[**TrainingsCreate**](DefaultAPI.md#TrainingsCreate) | **Post** /v1/models/{model_owner}/{model_name}/versions/{version_id}/trainings | Create a training
[**TrainingsGet**](DefaultAPI.md#TrainingsGet) | **Get** /v1/trainings/{training_id} | Get a training
[**TrainingsList**](DefaultAPI.md#TrainingsList) | **Get** /v1/trainings | List trainings



## CollectionsGet

> CollectionsGet(ctx, collectionSlug).Execute()

Get a collection of models



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    collectionSlug := TODO // interface{} | The slug of the collection, like `super-resolution` or `image-restoration`. See [replicate.com/collections](https://replicate.com/collections) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.CollectionsGet(context.Background(), collectionSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CollectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionSlug** | [**interface{}**](.md) | The slug of the collection, like &#x60;super-resolution&#x60; or &#x60;image-restoration&#x60;. See [replicate.com/collections](https://replicate.com/collections)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsList

> CollectionsList(ctx).Execute()

List collections of models



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.CollectionsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsGet

> ModelsGet(ctx, modelOwner, modelName).Execute()

Get a model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    modelOwner := TODO // interface{} | The name of the user or organization that owns the model. 
    modelName := TODO // interface{} | The name of the model. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ModelsGet(context.Background(), modelOwner, modelName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ModelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelOwner** | [**interface{}**](.md) | The name of the user or organization that owns the model.  | 
**modelName** | [**interface{}**](.md) | The name of the model.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsVersionsDelete

> ModelsVersionsDelete(ctx, modelOwner, modelName, versionId).Execute()

Delete a model version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    modelOwner := TODO // interface{} | The name of the user or organization that owns the model. 
    modelName := TODO // interface{} | The name of the model. 
    versionId := TODO // interface{} | The ID of the version. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ModelsVersionsDelete(context.Background(), modelOwner, modelName, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ModelsVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelOwner** | [**interface{}**](.md) | The name of the user or organization that owns the model.  | 
**modelName** | [**interface{}**](.md) | The name of the model.  | 
**versionId** | [**interface{}**](.md) | The ID of the version.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelsVersionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsVersionsGet

> ModelsVersionsGet(ctx, modelOwner, modelName, versionId).Execute()

Get a model version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    modelOwner := TODO // interface{} | The name of the user or organization that owns the model. 
    modelName := TODO // interface{} | The name of the model. 
    versionId := TODO // interface{} | The ID of the version. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ModelsVersionsGet(context.Background(), modelOwner, modelName, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ModelsVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelOwner** | [**interface{}**](.md) | The name of the user or organization that owns the model.  | 
**modelName** | [**interface{}**](.md) | The name of the model.  | 
**versionId** | [**interface{}**](.md) | The ID of the version.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelsVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsVersionsList

> ModelsVersionsList(ctx, modelOwner, modelName).Execute()

List model versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    modelOwner := TODO // interface{} | The name of the user or organization that owns the model. 
    modelName := TODO // interface{} | The name of the model. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ModelsVersionsList(context.Background(), modelOwner, modelName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ModelsVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelOwner** | [**interface{}**](.md) | The name of the user or organization that owns the model.  | 
**modelName** | [**interface{}**](.md) | The name of the model.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelsVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredictionsCancel

> PredictionsCancel(ctx, predictionId).Execute()

Cancel a prediction

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    predictionId := TODO // interface{} | The ID of the prediction you want to cancel. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PredictionsCancel(context.Background(), predictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PredictionsCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**predictionId** | [**interface{}**](.md) | The ID of the prediction you want to cancel.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPredictionsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredictionsCreate

> PredictionsCreate(ctx).PredictionsCreateRequest(predictionsCreateRequest).Execute()

Create a prediction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    predictionsCreateRequest := *openapiclient.NewPredictionsCreateRequest() // PredictionsCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PredictionsCreate(context.Background()).PredictionsCreateRequest(predictionsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PredictionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredictionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **predictionsCreateRequest** | [**PredictionsCreateRequest**](PredictionsCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredictionsGet

> PredictionsGet(ctx, predictionId).Execute()

Get a prediction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    predictionId := TODO // interface{} | The ID of the prediction you want to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PredictionsGet(context.Background(), predictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PredictionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**predictionId** | [**interface{}**](.md) | The ID of the prediction you want to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPredictionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredictionsList

> PredictionsList(ctx).Execute()

List predictions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PredictionsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PredictionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPredictionsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainingsCancel

> TrainingsCancel(ctx, trainingId).Execute()

Cancel a training

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    trainingId := TODO // interface{} | The ID of the training you want to cancel. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.TrainingsCancel(context.Background(), trainingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TrainingsCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trainingId** | [**interface{}**](.md) | The ID of the training you want to cancel.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrainingsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainingsCreate

> TrainingsCreate(ctx, modelOwner, modelName, versionId).TrainingsCreateRequest(trainingsCreateRequest).Execute()

Create a training



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    modelOwner := TODO // interface{} | The name of the user or organization that owns the model. 
    modelName := TODO // interface{} | The name of the model. 
    versionId := TODO // interface{} | The ID of the version. 
    trainingsCreateRequest := *openapiclient.NewTrainingsCreateRequest() // TrainingsCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.TrainingsCreate(context.Background(), modelOwner, modelName, versionId).TrainingsCreateRequest(trainingsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TrainingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelOwner** | [**interface{}**](.md) | The name of the user or organization that owns the model.  | 
**modelName** | [**interface{}**](.md) | The name of the model.  | 
**versionId** | [**interface{}**](.md) | The ID of the version.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrainingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **trainingsCreateRequest** | [**TrainingsCreateRequest**](TrainingsCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainingsGet

> TrainingsGet(ctx, trainingId).Execute()

Get a training



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {
    trainingId := TODO // interface{} | The ID of the training you want to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.TrainingsGet(context.Background(), trainingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TrainingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trainingId** | [**interface{}**](.md) | The ID of the training you want to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrainingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainingsList

> TrainingsList(ctx).Execute()

List trainings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stillmatic/go-replicate"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.TrainingsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TrainingsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTrainingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

