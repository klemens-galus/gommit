package main

import (
	"testing"
	"gommit.go/commit"
)

var(
	goodCommit = "fix: this is a good commit"
	badCommit = "test fix: this is a good commit"
	goodSize = "fix: this is a good commit"
	badSize = "test fix"
)
func TestSize(t *testing.T){
	result := gommit.CheckCommitSize(badSize)
	if result != false {
		t.Errorf("Expected false for commit message %s", badSize)
	}
	result = gommit.CheckCommitSize(goodSize)
	if result != true {
		t.Errorf("Expected true for commit message %s", goodSize)
	}
}

func TestLint(t *testing.T){
	result := gommit.CheckCommitLint(goodCommit)
	if result != true {
		t.Errorf("Expected true for commit message %s", goodCommit)
	}
	result = gommit.CheckCommitLint(badCommit)
	if result != false {
		t.Errorf("Expected false for commit message %s", badCommit)
	}
}