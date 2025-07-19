package main

import (
	"context"
	api "hanyuMQ/kitex_gen/api"
)

// Client_OperationsImpl implements the last service interface defined in the IDL.
type Client_OperationsImpl struct{}

// Pub implements the Client_OperationsImpl interface.
func (s *Client_OperationsImpl) Pub(ctx context.Context, req *api.PubRequest) (resp *api.PubResponse, err error) {
	// TODO: Your code here...
	return
}

// Pingpong implements the Client_OperationsImpl interface.
func (s *Client_OperationsImpl) Pingpong(ctx context.Context, req *api.PingPongRequest) (resp *api.PingPongResponse, err error) {
	// TODO: Your code here...
	return
}

// RequestVote implements the Raft_OperationsImpl interface.
func (s *Raft_OperationsImpl) RequestVote(ctx context.Context, rep *api.RequestVoteArgs_) (resp *api.RequestVoteReply, err error) {
	// TODO: Your code here...
	return
}

// AppendEntries implements the Raft_OperationsImpl interface.
func (s *Raft_OperationsImpl) AppendEntries(ctx context.Context, rep *api.AppendEntriesArgs_) (resp *api.AppendEntriesReply, err error) {
	// TODO: Your code here...
	return
}

// SnapShot implements the Raft_OperationsImpl interface.
func (s *Raft_OperationsImpl) SnapShot(ctx context.Context, rep *api.SnapShotArgs_) (resp *api.SnapShotReply, err error) {
	// TODO: Your code here...
	return
}

// Pingpongtest implements the Raft_OperationsImpl interface.
func (s *Raft_OperationsImpl) Pingpongtest(ctx context.Context, req *api.PingPongArgs_) (resp *api.PingPongReply, err error) {
	// TODO: Your code here...
	return
}
