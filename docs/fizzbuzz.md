# Fizzbuzz

Fizzbuzz takes two strings and three integers

#### Query-String

Parameter | Description
----------|-------------------
Int1      | first multiple 
Int2      | second multiple
Limit     | limit of fizzbuzz
String1   | string replacing the first multiple
String2   | second replacing the second multiple


### GET /fizzbuzz

Endpoint to get the fizz buzz json result. The parameters must be present in the query.

#### Request

```http
GET /fizzbuzz?String1=toto&String2=tata&Int1=2&Int2=4&Limit=10 HTTP/1.1
```
#### Status codes

* 200 OK, when the fizzbuzz has been successfully created
* 400 Bad Request, when the query parameters are false
* 401 Unauthorized, when int parameter is < to 1

### GET /fizzbuzz/example

Get a fizzbuzz with default parameters.

String1 = "fizz", String2 = "buzz", Int1 = 3, Int2 = 5, Limit = 100