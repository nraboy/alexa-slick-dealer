package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nraboy/slick-dealer/alexa"
)

type FeedResponse struct {
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func RequestFeed(mode string) (FeedResponse, error) {
	endpoint, _ := url.Parse("https://slickdeals.net/newsearch.php")
	queryParams := endpoint.Query()
	queryParams.Set("mode", mode)
	queryParams.Set("searcharea", "deals")
	queryParams.Set("searchin", "first")
	queryParams.Set("rss", "1")
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		return FeedResponse{}, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var feedResponse FeedResponse
		xml.Unmarshal(data, &feedResponse)
		return feedResponse, nil
	}
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "FrontpageDealIntent":
		response = HandleFrontpageDealIntent(request)
	case "PopularDealIntent":
		response = HandlePopularDealIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	default:
		response = HandleAboutIntent(request)
	}
	return response
}

func HandleFrontpageDealIntent(request alexa.Request) alexa.Response {
	feedResponse, _ := RequestFeed("frontpage")
	var builder alexa.SSMLBuilder
	builder.Say("Here are the current frontpage deals:")
	builder.Pause("1000")
	for _, item := range feedResponse.Channel.Item {
		builder.Say(item.Title)
		builder.Pause("1000")
	}
	return alexa.NewSSMLResponse("Frontpage Deals", builder.Build())
}

func HandlePopularDealIntent(request alexa.Request) alexa.Response {
	feedResponse, _ := RequestFeed("popdeals")
	var builder alexa.SSMLBuilder
	builder.Say("Here are the current popular deals:")
	builder.Pause("1000")
	for _, item := range feedResponse.Channel.Item {
		builder.Say(item.Title)
		builder.Pause("1000")
	}
	return alexa.NewSSMLResponse("Popular Deals", builder.Build())
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("Give me the frontpage deals.")
	builder.Pause("1000")
	builder.Say("Give me the popular deals.")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
