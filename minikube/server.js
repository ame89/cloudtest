let http = require('http');

let handleRequest = function(request, response) {
  console.log('Received request for URL: ' + request.url);
  response.writeHead(200);
  response.write('hello ')
  response.end(request.url)
};
let server = http.createServer(handleRequest);
server.listen(8080);
