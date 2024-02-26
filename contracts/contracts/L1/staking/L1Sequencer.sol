// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {IStaking} from "../staking/IStaking.sol";
import {IL1Sequencer} from "./IL1Sequencer.sol";

contract L1Sequencer is Initializable, IL1Sequencer, Sequencer, Pausable {
    // staking contract
    address public stakingContract;
    // rollup Contract
    address public rollupContract;

    // current sequencers version
    uint256 public override currentVersion = 0;
    // newest sequencers version
    uint256 public override newestVersion = 0;
    // map(version => sequencerAddress)
    mapping(uint256 => address[]) public sequencerAddrs;
    // map(version => sequencerBLSkeys)
    mapping(uint256 => bytes[]) public sequencerBLSKeys;

    /**
     * @notice sequencer version confirmed
     */
    event SequencerConfirmed(address[] sequencers, uint256 version);

    /**
     * @notice only staking contract
     */
    modifier onlyStakingContract() {
        require(msg.sender == stakingContract, "only staking contract");
        _;
    }

    /**
     * @notice only rollup contract
     */
    modifier onlyRollupContract() {
        require(msg.sender == rollupContract, "only rollup contract");
        _;
    }

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     */
    constructor(
        address payable _messenger
    ) Pausable() Sequencer(_messenger, payable(Predeploys.L2_SEQUENCER)) {
        _pause();
    }

    /**
     * @notice do not receive ETH
     */
    receive() external payable {
        require(false);
    }

    /**
     * @notice initializer
     * @param _stakingContract staking contract address
     * @param _rollupContract rollup contract address
     */
    function initialize(
        address _stakingContract,
        address _rollupContract
    ) public initializer {
        require(_stakingContract != address(0), "invalid staking contract");
        require(_rollupContract != address(0), "invalid rollup contract");
        stakingContract = _stakingContract;
        rollupContract = _rollupContract;
        _pause();
    }

    function pause() external override onlyStakingContract whenNotPaused {
        _pause();
    }

    function updateSequencersVersion(
        address[] memory _sequencerAddrs,
        bytes[] memory _sequencerBLSKeys
    ) internal {
        if (newestVersion == 0) {
            _unpause();
        }
        require(!paused(), "send message when unpaused");
        newestVersion++;
        sequencerAddrs[newestVersion] = _sequencerAddrs;
        sequencerBLSKeys[newestVersion] = _sequencerBLSKeys;
    }

    /**
     * @notice verify BLS signature
     * @param version sequencer set version
     * @param indexs sequencer index
     * @param signature batch signature
     */
    function verifySignature(
        uint256 version,
        uint256[] memory indexs,
        bytes memory signature
    ) external onlyRollupContract whenNotPaused returns (bool) {
        confirmVersion(version);
        // TODO: verify BLS signature
        return true;
    }

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        uint256[] memory sequencerIndex,
        address challenger,
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external onlyRollupContract {
        IStaking(stakingContract).slash(
            sequencerAddrs[currentVersion],
            sequencerIndex,
            challenger,
            _minGasLimit,
            _gasFee
        );
    }

    /**
     * @notice confirm sequencer ser version
     * @param version sequencer set version
     */
    function confirmVersion(uint256 version) internal {
        require(
            version >= currentVersion && version <= newestVersion,
            "invalid version"
        );
        for (uint256 i = 1; i < version; i++) {
            delete sequencerAddrs[i];
            delete sequencerBLSKeys[i];
        }
        currentVersion = version;
    }

    function updateAndSendSequencerSet(
        bytes memory _sequencerBytes,
        address[] memory _sequencerAddrs,
        bytes[] memory _sequencerBLSKeys,
        uint32 _gasLimit,
        address _refundAddress
    ) external payable override onlyStakingContract {
        updateSequencersVersion(_sequencerAddrs, _sequencerBLSKeys);
        require(!paused(), "send message when unpaused");
        MESSENGER.sendMessage{value: msg.value}(
            address(OTHER_SEQUENCER),
            0,
            _sequencerBytes,
            _gasLimit,
            _refundAddress
        );
    }

    function sequencerNum(uint256 version) external view returns (uint256) {
        return sequencerBLSKeys[version].length;
    }

    /**
     * @notice whether is current sequencer
     * @param addr address
     */
    function isSequencer(address addr) external view returns (bool) {
        for (uint256 i = 0; i < sequencerAddrs[currentVersion].length; i++) {
            if (sequencerAddrs[currentVersion][i] == addr) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice sequencer addresses
     * @param version version
     */
    function getSequencerAddrs(
        uint256 version
    ) external view returns (address[] memory) {
        return sequencerAddrs[version];
    }

    /**
     * @notice sequencer BLS keys
     * @param version version
     */
    function getSequencerBLSKeys(
        uint256 version
    ) external view returns (bytes[] memory) {
        return sequencerBLSKeys[version];
    }
}
