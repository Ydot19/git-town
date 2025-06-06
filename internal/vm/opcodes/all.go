// Generated by generate_opcodes_all.sh. Run `make fix` to update. DO NOT EDIT.

package opcodes

import "github.com/git-town/git-town/v20/internal/vm/shared"

// All provides all existing opcodes.
// This is used to iterate all opcode types.
func All() []shared.Opcode {
	return []shared.Opcode{
		&BranchCreateAndCheckoutExistingParent{},
		&BranchCreate{},
		&BranchCurrentResetToParent{},
		&BranchCurrentResetToSHAIfNeeded{},
		&BranchCurrentResetToSHA{},
		&BranchCurrentReset{},
		&BranchDeleteIfEmptyAtRuntime{},
		&BranchEnsureShippableChanges{},
		&BranchLocalDeleteContent{},
		&BranchLocalDelete{},
		&BranchLocalRename{},
		&BranchRemoteCreate{},
		&BranchRemoteSetToSHAIfNeeded{},
		&BranchRemoteSetToSHA{},
		&BranchReset{},
		&BranchTrackingCreate{},
		&BranchTrackingDelete{},
		&BranchTypeOverrideRemove{},
		&BranchTypeOverrideSet{},
		&BranchWithRemoteGoneDeleteIfEmptyAtRuntime{},
		&BrowserOpen{},
		&ChangesDiscard{},
		&ChangesStage{},
		&ChangesUnstageAll{},
		&CheckoutFirstExisting{},
		&CheckoutHistoryPreserve{},
		&CheckoutIfExists{},
		&CheckoutIfNeeded{},
		&CheckoutParentOrMain{},
		&CheckoutUncached{},
		&Checkout{},
		&CherryPick{},
		&CommitAutoUndo{},
		&CommitMessageCommentOut{},
		&CommitRemove{},
		&CommitRevertIfNeeded{},
		&CommitRevert{},
		&CommitWithMessage{},
		&Commit{},
		&ConfigRemove{},
		&ConfigSet{},
		&ConflictPhantomDetect{},
		&ConflictPhantomFinalize{},
		&ConflictPhantomResolve{},
		&ConnectorProposalMerge{},
		&FetchUpstream{},
		&LineageBranchRemove{},
		&LineageParentRemove{},
		&LineageParentSetFirstExisting{},
		&LineageParentSetIfExists{},
		&LineageParentSetToGrandParent{},
		&LineageParentSet{},
		&MergeAbort{},
		&MergeAlwaysProgram{},
		&MergeContinue{},
		&MergeFastForward{},
		&MergeParentIfNeeded{},
		&MergeParentResolvePhantomConflicts{},
		&MergeParent{},
		&MergeSquashAutoUndo{},
		&MergeSquashProgram{},
		&Merge{},
		&MessageQueue{},
		&ProgramEndOfBranch{},
		&ProposalCreate{},
		&ProposalUpdateSource{},
		&ProposalUpdateTargetToGrandParent{},
		&ProposalUpdateTarget{},
		&PullCurrentBranch{},
		&PushCurrentBranchForceIfNeeded{},
		&PushCurrentBranchForceIgnoreError{},
		&PushCurrentBranchForce{},
		&PushCurrentBranchIfLocal{},
		&PushCurrentBranchIfNeeded{},
		&PushCurrentBranch{},
		&PushTags{},
		&RebaseAbort{},
		&RebaseBranch{},
		&RebaseContinueIfNeeded{},
		&RebaseContinue{},
		&RebaseOntoKeepDeleted{},
		&RebaseOntoRemoveDeleted{},
		&RebaseParentIfNeeded{},
		&RebaseTrackingBranch{},
		&RegisterUndoablePerennialCommit{},
		&SnapshotInitialUpdateLocalSHAIfNeeded{},
		&SnapshotInitialUpdateLocalSHA{},
		&StashDrop{},
		&StashOpenChanges{},
		&StashPopIfNeeded{},
		&StashPop{},
		&UndoLastCommit{},
	} //exhaustruct:ignore
}
