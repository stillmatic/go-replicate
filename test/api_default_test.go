/*
Replicate HTTP API

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package replicate

import (
	"context"
	"testing"

	openapiclient "github.com/stillmatic/go-replicate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_replicate_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService CollectionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionSlug interface{}

		httpRes, err := apiClient.DefaultAPI.CollectionsGet(context.Background(), collectionSlug).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CollectionsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.CollectionsList(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ModelsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelOwner interface{}
		var modelName interface{}

		httpRes, err := apiClient.DefaultAPI.ModelsGet(context.Background(), modelOwner, modelName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ModelsVersionsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelOwner interface{}
		var modelName interface{}
		var versionId interface{}

		httpRes, err := apiClient.DefaultAPI.ModelsVersionsDelete(context.Background(), modelOwner, modelName, versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ModelsVersionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelOwner interface{}
		var modelName interface{}
		var versionId interface{}

		httpRes, err := apiClient.DefaultAPI.ModelsVersionsGet(context.Background(), modelOwner, modelName, versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ModelsVersionsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelOwner interface{}
		var modelName interface{}

		httpRes, err := apiClient.DefaultAPI.ModelsVersionsList(context.Background(), modelOwner, modelName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PredictionsCancel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var predictionId interface{}

		httpRes, err := apiClient.DefaultAPI.PredictionsCancel(context.Background(), predictionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PredictionsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.PredictionsCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PredictionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var predictionId interface{}

		httpRes, err := apiClient.DefaultAPI.PredictionsGet(context.Background(), predictionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PredictionsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.PredictionsList(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService TrainingsCancel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trainingId interface{}

		httpRes, err := apiClient.DefaultAPI.TrainingsCancel(context.Background(), trainingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService TrainingsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelOwner interface{}
		var modelName interface{}
		var versionId interface{}

		httpRes, err := apiClient.DefaultAPI.TrainingsCreate(context.Background(), modelOwner, modelName, versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService TrainingsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trainingId interface{}

		httpRes, err := apiClient.DefaultAPI.TrainingsGet(context.Background(), trainingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService TrainingsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.TrainingsList(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
