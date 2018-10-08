# Slick Dealer Skill for Amazon Alexa

[Slick Dealer](https://www.amazon.com/gp/product/B07J43J36F?ie=UTF8&ref-suffix=ss_rw) is an Amazon Alexa Skill built with the Go programming language. It is an unofficial virtual assistant to the [Slickdeals](https://www.slickdeals.net) website, with no endorsements from Slickdeals.

## How it Works

While Slickdeals does not have an official API for consuming its data, it does offer a public RSS feed that contains deal titles along with descriptions and links. The RSS feed is in XML format which is easily readable with Golang.

The Slick Dealer Skill, while written in Golang, is designed to be ran as an AWS Lambda function. The responses to the function are formatted for Amazon Alexa.

## Author Information

This Skill was written by [Nic Raboy](https://www.nraboy.com), who is an advocate of modern web and mobile development technologies. He has experience in Java, JavaScript, Golang and a variety of frameworks such as Angular, NativeScript, and Apache Cordova. Nic writes about his development experiences related to making web and mobile development easier to understand.

## Resources

The Polyglot Developer - [https://www.thepolyglotdeveloper.com](https://www.thepolyglotdeveloper.com)