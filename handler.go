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

// Sub implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) Sub(ctx context.Context, req *api.SubRequest) (resp *api.SubResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateTopic implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) CreateTopic(ctx context.Context, req *api.CreateTopicRequest) (resp *api.CreateTopicResponse, err error) {
	// TODO: Your code here...
	return
}

// CreatePart implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) CreatePart(ctx context.Context, req *api.CreatePartRequest) (resp *api.CreatePartResponse, err error) {
	// TODO: Your code here...
	return
}

// SetPartitionState implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) SetPartitionState(ctx context.Context, req *api.SetPartitionStateRequest) (resp *api.SetPartitionStateResponse, err error) {
	// TODO: Your code here...
	return
}

// BroInfo implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) BroInfo(ctx context.Context, req *api.BroInfoRequest) (resp *api.BroInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ProGetBroker implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) ProGetBroker(ctx context.Context, req *api.ProGetBrokRequest) (resp *api.ProGetBrokResponse, err error) {
	// TODO: Your code here...
	return
}

// ConStartGetBroker implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) ConStartGetBroker(ctx context.Context, req *api.ConStartGetBrokRequest) (resp *api.ConStartGetBrokResponse, err error) {
	// TODO: Your code here...
	return
}

// BroGetConfig implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) BroGetConfig(ctx context.Context, req *api.BroGetConfigRequest) (resp *api.BroGetConfigResponse, err error) {
	// TODO: Your code here...
	return
}

// BecomeLeader implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) BecomeLeader(ctx context.Context, req *api.BecomeLeaderRequest) (resp *api.BecomeLeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// GetNewLeader implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) GetNewLeader(ctx context.Context, req *api.GetNewLeaderRequest) (resp *api.GetNewLeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePTPOffset implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) UpdatePTPOffset(ctx context.Context, req *api.UpdatePTPOffsetRequest) (resp *api.UpdatePTPOffsetResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateDup implements the ZkServer_OperationsImpl interface.
func (s *ZkServer_OperationsImpl) UpdateDup(ctx context.Context, req *api.UpdateDupRequest) (resp *api.UpdateDupResponse, err error) {
	// TODO: Your code here...
	return
}
