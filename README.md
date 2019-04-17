# Mandrill as a microservice
An OMG service for Mandrill, is a powerful delivery service that can be used for personalized, one-to-one e-commerce emails, and automated transactional emails like password resets, order confirmations, and welcome messages.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-mandrill.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-mandrill)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-mandrill/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-mandrill)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Send Email
```sh
$ omg run send -a from=<SENDER_MAIL_ADDRESS> -a to=<RECEIVER_EMAIL_ADDRESS> -a message=<EMAIL_MESSAGE_BODY> -a template_name=<TEMPLATE_NAME> -a subject=<EMAIL_SUBJECT> -e API_KEY=<API_KEY>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-mandrill .
```
### RUN
```
docker run -p 3000:3000 microservice-mandrill
```
