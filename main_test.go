package main

import (
	"testing"

	"github.com/nraboy/slick-dealer/alexa"
	"github.com/stretchr/testify/assert"
)

func TestFeedRequest(t *testing.T) {
	feedResponse, err := RequestFeed("frontpage")
	assert.Nil(t, err, "The error should be nil")
	assert.NotEmpty(t, feedResponse, "The response should not be empty")
	assert.NotZero(t, feedResponse.Channel.Item, "The feed should have at least one entry")
}

func TestHandleFrontpageDealIntent(t *testing.T) {
	response := HandleFrontpageDealIntent(alexa.Request{})
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandlePopularDealIntent(t *testing.T) {
	response := HandlePopularDealIntent(alexa.Request{})
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleHelpIntent(t *testing.T) {
	response := HandleHelpIntent(alexa.Request{})
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleAboutIntent(t *testing.T) {
	response := HandleAboutIntent(alexa.Request{})
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}
