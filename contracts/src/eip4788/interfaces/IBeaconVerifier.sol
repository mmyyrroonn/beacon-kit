// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import { SSZ } from "../SSZ.sol";

interface IBeaconVerifier {
    /// @notice Emitted when the zero validator pubkey Generalized Index is
    /// changed.
    event ZeroValidatorPubkeyGIndexChanged(
        uint256 newZeroValidatorPubkeyGIndex
    );
    /// @notice Emitted when the execution number Generalized Index is changed.
    event ExecutionNumberGIndexChanged(uint256 newExecutionNumberGIndex);

    /// @dev Generalized Index of the pubkey of the first validator (validator
    /// index of 0) in the registry of the beacon state in the beacon block.
    /// @dev In the Deneb beacon chain fork, this should be 3254554418216960.
    function zeroValidatorPubkeyGIndex() external view returns (uint256);

    /// @dev Generalized Index of the execution number in the beacon state in
    /// the beacon block.
    /// @dev In the Deneb beacon chain fork, this should be 5894.
    function executionNumberGIndex() external view returns (uint256);

    /// @notice Get the parent beacon block root from a block's timestamp.
    /// @param timestamp `uint64` timestamp of the block.
    function getParentBeaconBlockRootAt(uint64 timestamp)
        external
        view
        returns (bytes32);

    /// @notice Get the parent beacon block root at `block.timestamp`.
    function getParentBeaconBlockRoot() external view returns (bytes32);

    /// @notice Verifies the proposer within the beacon block at the given
    /// timestamp. Reverts if proof invalid.
    /// @param timestamp `uint64` timestamp of the block.
    /// @param validatorPubkeyProof `bytes32[]` proof of the validator pubkey.
    /// @param validatorPubkey `ValidatorPubkey` to verify.
    /// @param proposerIndex `uint64` validator index of the proposer.
    function verifyBeaconBlockProposer(
        uint64 timestamp,
        bytes32[] calldata validatorPubkeyProof,
        bytes calldata validatorPubkey,
        uint64 proposerIndex
    )
        external
        view;

    /// @notice Verifies the execution number in the beacon block at the given
    /// timestamp. Reverts if proof invalid.
    /// @param timestamp `uint64` timestamp of the block.
    /// @param executionNumberProof `bytes32[]` proof of the execution number.
    /// @param blockNumber `uint64` execution number of the block.
    function verifyExecutionNumber(
        uint64 timestamp,
        bytes32[] calldata executionNumberProof,
        uint64 blockNumber
    )
        external
        view;
}
