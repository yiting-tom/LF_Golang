# 11 HTTP Multipart Requests

## Introduction of multipart/form-data
-  one of **Content-Type** which was defined by **MIME Type** (RCF 7578)
-  constructed by **multiple parts** and these parts are separated by the **boundary** argument.
- Formally **Content-Type** can only send data in 1 type, so **multipart/form-data** is used to send multiple types of files in a single request.

### Content of HTTP Request
```http
<!-- Request Header -->
POST /upload HTTP/1.1
Content-Length: 500
Content-Type: multipart/form-data; boundary=customBoundary

<!-- Request Body -->
<!-- part 1 -->
--customBoundary
Content-Disposition: form-data; name="id"
Content-Type: text/plain
... text data ...
<!-- part 2 -->
--customBoundary
Content-Disposition: form-data; name="book"; filename="la-textbook.pdf"
Content-Type: application/octet-stream
... pdf data ...
<!-- part 3 -->
--customBoundary 
Content-Disposition: form-data; name="notes"
Content-Type: application/json;
... {json data} ...
<!-- use --customBoundary-- to end the request -->
--customBoundary--
```
- Request Header: Not only need to send **Content-Type**, but also need to define the **boundary** argument.
- part 1: Send plain text data with name `id`.
- part 2: Send pdf file with name `book` and filename `la-textbook.pdf`.
- part 3: Send JSON data with name `notes`.

### form-data by HTML form tag
```html
<form enctype="multipart/form-data" action="/upload" method="POST">
  <input type="text" name="name" />
  <input type="file" name="file" />
  <button>Submit</button>
</form> 
```
After submitting the form, the browser will send the request to the server.
```http
POST /upload HTTP/1.1
Host: localhost:3000

Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryFYGn56LlBDLnAkfd
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36

------WebKitFormBoundaryFYGn56LlBDLnAkfd
Content-Disposition: form-data; name="name"

Test
------WebKitFormBoundaryFYGn56LlBDLnAkfd
Content-Disposition: form-data; name="file"; filename="text.txt"
Content-Type: text/plain

Hello World
------WebKitFormBoundaryFYGn56LlBDLnAkfd--
```