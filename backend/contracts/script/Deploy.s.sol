// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {Master} from "../src/Master.sol";
import {ZonaOracle} from "../src/Oracle.sol";

contract CounterScript is Script {
    ZonaOracle oracle;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        oracle = new ZonaOracle();
        new Master(address(oracle));

        vm.stopBroadcast();
    }
}
