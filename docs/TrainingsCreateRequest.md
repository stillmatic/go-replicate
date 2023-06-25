# TrainingsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **interface{}** | A string representing the desired model to push to in the format &#x60;{destination_model_owner}/{destination_model_name}&#x60;. This should be an existing model owned by the user or organization making the API request. If the destination is invalid, the server returns an appropriate 4XX response. | [optional] 
**Input** | Pointer to **interface{}** | An object containing inputs to the Cog model&#39;s &#x60;train()&#x60; function. | [optional] 
**Webhook** | Pointer to **interface{}** | An HTTPS URL for receiving a webhook when the training completes. The webhook will be a POST request where the request body is the same as the response body of the [get training](#trainings.get) operation. If there are network problems, we will retry the webhook a few times, so make sure it can be safely called more than once. | [optional] 

## Methods

### NewTrainingsCreateRequest

`func NewTrainingsCreateRequest() *TrainingsCreateRequest`

NewTrainingsCreateRequest instantiates a new TrainingsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingsCreateRequestWithDefaults

`func NewTrainingsCreateRequestWithDefaults() *TrainingsCreateRequest`

NewTrainingsCreateRequestWithDefaults instantiates a new TrainingsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *TrainingsCreateRequest) GetDestination() interface{}`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TrainingsCreateRequest) GetDestinationOk() (*interface{}, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TrainingsCreateRequest) SetDestination(v interface{})`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *TrainingsCreateRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *TrainingsCreateRequest) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *TrainingsCreateRequest) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetInput

`func (o *TrainingsCreateRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TrainingsCreateRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TrainingsCreateRequest) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *TrainingsCreateRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *TrainingsCreateRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *TrainingsCreateRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetWebhook

`func (o *TrainingsCreateRequest) GetWebhook() interface{}`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *TrainingsCreateRequest) GetWebhookOk() (*interface{}, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *TrainingsCreateRequest) SetWebhook(v interface{})`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *TrainingsCreateRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *TrainingsCreateRequest) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *TrainingsCreateRequest) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


