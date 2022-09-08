# Introduction

This platform is a copy of Adam's Freeman project, described in book "Pro Go The Complete Guide to Programming Reliable and Efficient Software Using Golang (2022) Adam Freeman" with some improvements and adaptations for use in Brucheion project.

## Config

* changed logic config loading: all parameters load from config.json, but platform check existing variables in os.env with prefix defined in json and redefine config values to os.env values in the presence of; new env can be added to leaf level only (without "_" symbol)
* environment variables take precedence over config.json.
* for redefine a variable, you must specify all path. 
for example:
config.json variable sql:driverName = postgres
os.env should be: [prefix]_sql_driverName = postgres
where prefix is the config.json parameter system:prefix

## Handling

* add possibility to render URL-path with dot and colon symbols (http/handling/routes.go/generateRegularExpression())
* add render Content-Type = "multipart/form-data" for load files (http/handling/params/processor,go/GetParametersFromRequest()). In this time rendering one file only, it have an opportunity to load multiple...
 

## Pipeline

To handle HTTP requests from browsers, we create a simple pipeline that will contain middleware components that can inspect and modify requests. When an HTTP request arrives, it will be passed to each registered middleware component in the pipeline, giving each component the chance to process the request and contribute to the response. Components will also be able to terminate request processing, preventing the request from being forwarded to the remaining components in the pipeline. Once the request has reached the end of the pipeline, it works its way back along the pipeline so that components have a chance to make further changes or do further work
 
