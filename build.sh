rm handler.zip
rm slick-dealer
GOOS=linux go build
zip handler.zip ./slick-dealer
