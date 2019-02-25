const http = require("http")

http.createServer((req,res)=>{
  console.log(req,res);
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('mac-peng');
}).listen("3000",()=>{
  console.log('success listen 3000');
})