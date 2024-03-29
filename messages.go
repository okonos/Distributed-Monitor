package main

import "fmt"

type requestVoteMsg struct {
	Term         int
	LastLogIndex int
	LastLogTerm  int
}

type requestVoteResponse struct {
	Term        int
	VoteGranted bool
}

type appendEntriesMsg struct {
	Term         int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []entry // log entries to store (empty for heartbeat)
	LeaderCommit int     // leader's commitIndex
}

type appendEntriesResponse struct {
	Term       int
	Success    bool // true if follower had entry matching prevLog{Index, Term}
	MatchIndex int
}

type clientRequest struct {
	Command  string
	Argument int
}

func (req clientRequest) String() string {
	if req.Command == "PUT" {
		return fmt.Sprintf("%s %2d", req.Command, req.Argument)
	}
	return req.Command + "   " // 3 spaces for alignment
}

type clientResponse struct {
	Text  string
	Value int
}
