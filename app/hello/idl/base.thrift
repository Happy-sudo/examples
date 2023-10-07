namespace go base.v1

// 固定 （code：1 正常，0：错误，-1：异常）
// （msg：可作为 error处理）
struct EmptyReply {
    1:i8 code
    2:string msg
}

struct PageInfo {
    1:i64 page (api.vd="$>0")
    2:i64 page_size (api.vd="$>0")
}
