namespace go greet.v1

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service Greet {
    Response Greet(1: Request req)
}
