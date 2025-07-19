namespace go api

struct PushRequest {
    1: string producer
    2: string topic
    3: string key
    4: binary message
}

struct PushResponse {
    1: bool     ret
}

struct PullRequest {
    1: string consumer
    2: string topic
    3: string key
}

struct PullResponse {
    1: binary   message
}

//consmer发送自己的host和ip,使broker连接上自己
struct InfoRequest {
    1: string ip_port
}

struct InfoResponse {
    1: bool ret
}

service Server_Operations {
    PushResponse            push(       1: PushRequest      req)               //producer used
    PullResponse            pull(       1: PullRequest      req)               //
    InfoResponse            info(    1: InfoRequest      req)            //consumer used
}

struct PubRequest{
    1: string meg
}

struct PubResponse{
    1: bool ret
}

struct PingPongRequest {
    1: bool ping
}

struct PingPongResponse{
    1: bool pong
}

service Client_Operations {
    PubResponse         pub(        1: PubRequest       req)
    PingPongResponse    pingpong(   1: PingPongRequest  req)
}