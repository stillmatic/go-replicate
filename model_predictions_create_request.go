/*
Replicate HTTP API

A web service for running Replicate models

API version: 1.0.0-a1
Contact: team@replicate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package replicate

import (
	"encoding/json"
)

// checks if the PredictionsCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictionsCreateRequest{}

// PredictionsCreateRequest struct for PredictionsCreateRequest
type PredictionsCreateRequest struct {
	// The model's input as a JSON object. The input depends on what model you are running. To see the available inputs, click the \"Run with API\" tab on the model you are running. For example, stability-ai/stable-diffusion takes prompt as an input.  Files should be passed as data URLs or HTTPS URLs. 
	Input interface{} `json:"input,omitempty"`
	// The ID of the model version that you want to run.
	Version interface{} `json:"version,omitempty"`
	// An HTTPS URL for receiving a webhook when the prediction has new output. The webhook will be a POST request where the request body is the same as the response body of the [get prediction](#predictions.get) operation. If there are network problems, we will retry the webhook a few times, so make sure it can be safely called more than once.
	Webhook interface{} `json:"webhook,omitempty"`
	// By default, we will send requests to your webhook URL whenever there are new logs, new outputs, or the prediction has finished. You can change which events trigger webhook requests by specifying `webhook_events_filter` in the prediction request.  * `start`: immediately on prediction start * `output`: each time a prediction generates an output (note that predictions can generate multiple outputs) * `logs`: each time log output is generated by a prediction * `completed`: when the prediction reaches a terminal state (succeeded/canceled/failed)  For example, if you only wanted requests to be sent at the start and end of the prediction, you would provide:  ```json {   \"version\": \"5c7d5dc6dd8bf75c1acaa8565735e7986bc5b66206b55cca93cb72c9bf15ccaa\",   \"input\": {     \"text\": \"Alice\"   },   \"webhook\": \"https://example.com/my-webhook\",   \"webhook_events_filter\": [\"start\", \"completed\"] } ```  Requests for event types `output` and `logs` will be sent at most once every 500ms. If you request `start` and `completed` webhooks, then they'll always be sent regardless of throttling. 
	WebhookEventsFilter interface{} `json:"webhook_events_filter,omitempty"`
}

// NewPredictionsCreateRequest instantiates a new PredictionsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionsCreateRequest() *PredictionsCreateRequest {
	this := PredictionsCreateRequest{}
	return &this
}

// NewPredictionsCreateRequestWithDefaults instantiates a new PredictionsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionsCreateRequestWithDefaults() *PredictionsCreateRequest {
	this := PredictionsCreateRequest{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredictionsCreateRequest) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredictionsCreateRequest) GetInputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *PredictionsCreateRequest) HasInput() bool {
	if o != nil && IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *PredictionsCreateRequest) SetInput(v interface{}) {
	o.Input = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredictionsCreateRequest) GetVersion() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredictionsCreateRequest) GetVersionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return &o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PredictionsCreateRequest) HasVersion() bool {
	if o != nil && IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given interface{} and assigns it to the Version field.
func (o *PredictionsCreateRequest) SetVersion(v interface{}) {
	o.Version = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredictionsCreateRequest) GetWebhook() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredictionsCreateRequest) GetWebhookOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return &o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *PredictionsCreateRequest) HasWebhook() bool {
	if o != nil && IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given interface{} and assigns it to the Webhook field.
func (o *PredictionsCreateRequest) SetWebhook(v interface{}) {
	o.Webhook = v
}

// GetWebhookEventsFilter returns the WebhookEventsFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredictionsCreateRequest) GetWebhookEventsFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEventsFilter
}

// GetWebhookEventsFilterOk returns a tuple with the WebhookEventsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredictionsCreateRequest) GetWebhookEventsFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEventsFilter) {
		return nil, false
	}
	return &o.WebhookEventsFilter, true
}

// HasWebhookEventsFilter returns a boolean if a field has been set.
func (o *PredictionsCreateRequest) HasWebhookEventsFilter() bool {
	if o != nil && IsNil(o.WebhookEventsFilter) {
		return true
	}

	return false
}

// SetWebhookEventsFilter gets a reference to the given interface{} and assigns it to the WebhookEventsFilter field.
func (o *PredictionsCreateRequest) SetWebhookEventsFilter(v interface{}) {
	o.WebhookEventsFilter = v
}

func (o PredictionsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionsCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.WebhookEventsFilter != nil {
		toSerialize["webhook_events_filter"] = o.WebhookEventsFilter
	}
	return toSerialize, nil
}

type NullablePredictionsCreateRequest struct {
	value *PredictionsCreateRequest
	isSet bool
}

func (v NullablePredictionsCreateRequest) Get() *PredictionsCreateRequest {
	return v.value
}

func (v *NullablePredictionsCreateRequest) Set(val *PredictionsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionsCreateRequest(val *PredictionsCreateRequest) *NullablePredictionsCreateRequest {
	return &NullablePredictionsCreateRequest{value: val, isSet: true}
}

func (v NullablePredictionsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


