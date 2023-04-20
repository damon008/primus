namespace go hello


struct ReqBody {
    1: string name
    2: i32 type
    3: string email
}

struct Request {
	1: string data
	2: string message
	3: ReqBody reqBody
}

struct Msg {
	1: i64 status
	2: i64 code
	3: string msg
}

struct Response {
	1: Msg msg
	2: string data
}

service HelloService {
    Response echo(1: Request req)
    Response Get()
    Response GetH(1: i32 id)(api.get="/api/v1/getH/:id")
    Response Post(1: Request req)(api.post="/api/v1/create")
}