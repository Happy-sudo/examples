namespace go xxx.v1

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service XXX {
    Response XXX(1: Request req)
}
