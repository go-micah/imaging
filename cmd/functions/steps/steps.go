package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-micah/imaging"
	"github.com/go-micah/imaging/utils"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	width := request.QueryStringParameters["width"]
	if len(width) == 0 {
		width = "512"
	}

	height := request.QueryStringParameters["height"]
	if len(height) == 0 {
		height = "512"
	}

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "{\"message\":\"invalid width\"}",
		}, nil
	}

	widthInt = utils.Abs(widthInt)

	if widthInt > 1024 {
		widthInt = 1024
	}

	heightInt, err := strconv.Atoi(height)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "{\"message\":\"invalid height\"}",
		}, nil
	}

	heightInt = utils.Abs(heightInt)

	if heightInt > 1024 {
		heightInt = 1024
	}

	img := imaging.Steps(widthInt, heightInt)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "{\"message\":\"error encoding the image\"}",
		}, nil
	}

	b64img := base64.StdEncoding.EncodeToString(buf.Bytes())

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
		Body:            b64img,
		IsBase64Encoded: true,
		StatusCode:      http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
