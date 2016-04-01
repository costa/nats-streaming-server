// Copyright 2016 Apcera Inc. All rights reserved.

package stores

import (
	"testing"
)

func createDefaultMemStore(t *testing.T) *MemoryStore {
	ms, err := NewMemoryStore()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	ms.SetChannelLimits(testDefaultChannelLimits)
	return ms
}

func TestMSBasicCreate(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testBasicCreate(t, ms, "MEMORY")
}

func TestMSNothingRecoveredOnFreshStart(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testNothingRecoveredOnFreshStart(t, ms)
}

func TestMSNewChannel(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testNewChannel(t, ms)
}

func TestMSCloseIdempotent(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testCloseIdempotent(t, ms)
}

func TestMSBasicMsgStore(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testBasicMsgStore(t, ms)
}

func TestMSMsgsState(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testMsgsState(t, ms)
}

func TestMSMaxMsgs(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	limitCount := 100

	limits := testDefaultChannelLimits
	limits.MaxNumMsgs = limitCount

	ms.SetChannelLimits(limits)

	testMaxMsgs(t, ms, limitCount)
}

func TestMSMaxChannels(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	limitCount := 2

	limits := testDefaultChannelLimits
	limits.MaxChannels = limitCount

	ms.SetChannelLimits(limits)

	testMaxChannels(t, ms, limitCount)
}

func TestMSMaxSubs(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	limitCount := 2

	limits := testDefaultChannelLimits
	limits.MaxSubs = limitCount

	ms.SetChannelLimits(limits)

	testMaxSubs(t, ms, limitCount)
}

func TestMSBasicSubStore(t *testing.T) {
	ms := createDefaultMemStore(t)
	defer ms.Close()

	testBasicSubStore(t, ms)
}
