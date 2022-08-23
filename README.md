### Pipeline

To handle HTTP requests from browsers, we create a simple pipeline that will contain middleware components that can inspect and modify requests. When an HTTP request arrives, it will be passed to each registered middleware component in the pipeline, giving each component the chance to process the request and contribute to the response. Components will also be able to terminate request processing, preventing the request from being forwarded to the remaining components in the pipeline. Once the request has reached the end of the pipeline, it works its way back along the pipeline so that components have a chance to make further changes or do further work
 
