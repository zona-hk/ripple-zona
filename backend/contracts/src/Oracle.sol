// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract ZonaOracle {
    mapping(bytes32 => uint256) private indexes;
    address owner;

    constructor() {
        owner = msg.sender;
    }

    // categoryId is 0 or 1 based on what kind of index it is (AQI or real estate)
    function setValue(
        uint256 categoryId,
        string calldata cityName,
        uint256 value
    ) public {
        require(msg.sender == owner);
        bytes32 key = keccak256(abi.encodePacked(categoryId, cityName));
        indexes[key] = value;
    }

    function getValue(
        uint256 categoryId,
        string calldata cityName
    ) public view returns (uint256) {
        bytes32 key = keccak256(abi.encodePacked(categoryId, cityName));
        return indexes[key];
    }
}
