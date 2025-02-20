// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ZonaOracle} from "./Oracle.sol";

struct Position {
    bytes32 oracleKey; // Unique identifier for fetching data
    uint256 guess;
    bool long;
    uint256 timeframe; // i.e. 1h, 2h, 4h, 6h, 8h, 24h
    uint256 timestamp;
    bool isOccupied;
    bool resolvable;
}

struct PositionInfo {
    bytes32 oracleKey;
    uint256 timeframe;
}

struct Player {
    uint256 wins;
    uint256 losses;
    bool initialized;
    mapping(bytes32 => Position) positions;
    mapping(bytes32 => uint256) actualValues;
    uint256 numPositions;
    mapping(string => uint256) cityCount;
    mapping(uint256 => uint256) categoriesCount;
}

struct City {
    string cityName;
    bool realEstate;
    bool airQuality;
}

contract Master {
    uint256 public constant HOUR = 3600; // hour in seconds
    uint256 public constant REALESTATE = 0;
    uint256 public constant AIRQUALITY = 1;

    event PositionResolved(
        address indexed player,
        bytes32 indexed oracleKey,
        uint256 guess,
        bool long,
        uint256 timeframe,
        uint256 timestamp,
        bool won,
        uint256 finalValue
    );

    //uint256[6] public timeframes;
    uint256[7] public timeframes; // DEBUG
    address private owner;
    ZonaOracle private oracle;
    uint256 public currentTime;
    address[] public players;

    mapping(address => Player) public records;

    mapping(string => bool) public validCitiesRealEstate;
    mapping(string => bool) public validCitiesAirQuality;

    function initTimeframes() private {
        timeframes[0] = 1;
        timeframes[1] = 2;
        timeframes[2] = 4;
        timeframes[3] = 6;
        timeframes[4] = 8;
        timeframes[5] = 24;
        //timeframes[6] = 0; // DEBUG
    }

    function initCities() private {
        // Real estate
        validCitiesRealEstate["hongkong"] = true;
        validCitiesRealEstate["singapore"] = true;
        validCitiesRealEstate["dubai"] = true;
        validCitiesRealEstate["sydney"] = true;
        validCitiesRealEstate["melbourne"] = true;
        validCitiesRealEstate["adelaide"] = true;
        validCitiesRealEstate["brisbane"] = true;
        validCitiesRealEstate["london"] = true;

        // Air Quality
        validCitiesAirQuality["delhi"] = true;
        validCitiesAirQuality["hongkong"] = true;
    }

    constructor(address oracle_addr) {
        owner = msg.sender;
        oracle = ZonaOracle(oracle_addr); //replace with Stork oracle later
        initTimeframes();
        initCities();
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not Owner");

        _;
    }

    function newPlayer(address addr) private {
        records[addr].initialized = true;
        players.push(addr);
    }

    function addCity(
        string calldata city,
        bool realEstate,
        bool airQuality
    ) public onlyOwner {
        require(
            validCitiesRealEstate[city] != false &&
                validCitiesAirQuality[city] != false,
            "City exists already"
        );
        require(
            realEstate == true || airQuality == true,
            "you must set realEstate or airQuality to true"
        );
        if (realEstate) {
            validCitiesRealEstate[city] = true;
        } else {
            validCitiesAirQuality[city] = true;
        }
    }

    function createPosition(
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex,
        bool long
    ) public {
        // NOTES:
        // timeframeIndex is 0, 1, 2, etc according to the timeframe (e.g. index 0 = 1h, 1 = 2h etc...)
        require(
            validCitiesRealEstate[cityName] == true ||
                validCitiesAirQuality[cityName] == true,
            "City is not valid"
        );
        require(timeframeIndex < timeframes.length);
        Player storage player = records[msg.sender];
        bytes32 positionKey = keccak256(
            abi.encodePacked(categoryId, cityName, timeframes[timeframeIndex])
        );
        // NOTE: resolve position if exists
        if (player.positions[positionKey].isOccupied) {
            // resolve position first if it exists
            resolvePosition(categoryId, cityName, timeframeIndex);
        }

        // value is the index value during the guessed time
        if (player.initialized == false) {
            newPlayer(msg.sender);
        }
        uint256 currentGuess = oracle.getValue(categoryId, cityName); // Get the current value for the index
        require(currentGuess != 0); // incase the oracle hasn't been updated for the city
        bytes32 oracleKey = keccak256(abi.encodePacked(categoryId, cityName));
        player.positions[positionKey] = Position(
            oracleKey,
            currentGuess,
            long,
            timeframes[timeframeIndex],
            currentTime,
            true,
            false
        );
        player.numPositions++;
        player.categoriesCount[categoryId]++;
        player.cityCount[cityName]++;
    }

    function getCategoriesCount(
        address player,
        uint256 categoryId
    ) public view returns (uint256) {
        return records[player].categoriesCount[categoryId];
    }

    function getCityCount(
        address player,
        string calldata cityName
    ) public view returns (uint256) {
        return records[player].cityCount[cityName];
    }

    function fillActualValues(
        address playerAddr,
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex,
        uint256 actualValue
    ) public onlyOwner {
        require(timeframeIndex < timeframes.length);
        Player storage player = records[playerAddr];
        uint256 timeframe = timeframes[timeframeIndex];
        bytes32 positionKey = keccak256(
            abi.encodePacked(categoryId, cityName, timeframes[timeframeIndex])
        );
        Position storage position = player.positions[positionKey];
        if (position.isOccupied && !position.resolvable) {
            if (currentTime >= position.timestamp + HOUR * timeframe) {
                player.actualValues[positionKey] = actualValue;
                position.resolvable = true;
            }
        }
    }

    function forceResolvePositionsAll(
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex
    ) public onlyOwner {
        require(timeframeIndex < timeframes.length);
        for (uint256 i; i < players.length; ) {
            Player storage player = records[players[i]];
            bytes32 positionKey = keccak256(
                abi.encodePacked(
                    categoryId,
                    cityName,
                    timeframes[timeframeIndex]
                )
            );
            Position storage position = player.positions[positionKey];
            position.resolvable = true;
            unchecked {
                i++;
            }
        }
    }

    function forceResolvePositionsPlayer(
        address playerAddr,
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex
    ) public onlyOwner {
        require(timeframeIndex < timeframes.length);
        Player storage player = records[playerAddr];
        bytes32 positionKey = keccak256(
            abi.encodePacked(categoryId, cityName, timeframes[timeframeIndex])
        );
        Position storage position = player.positions[positionKey];
        position.resolvable = true;
    }

    function fillActualValuesAll(
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex,
        uint256 actualValue
    ) public onlyOwner {
        for (uint256 i; i < players.length; ) {
            fillActualValues(
                players[i],
                categoryId,
                cityName,
                timeframeIndex,
                actualValue
            );
            unchecked {
                i++;
            }
        }
    }

    function fillActualValuesBatches(
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex,
        uint256 actualValue,
        uint256 startIndex,
        uint256 endIndex
    ) public onlyOwner {
        require(timeframeIndex < timeframes.length);
        require(startIndex < endIndex);
        require(endIndex <= players.length);
        for (uint256 i = startIndex; i < endIndex; ) {
            fillActualValues(
                players[i],
                categoryId,
                cityName,
                timeframeIndex,
                actualValue
            );
            unchecked {
                i++;
            }
        }
    }

    function getActualValue(
        address playerAddr,
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex
    ) public view returns (uint256) {
        require(timeframeIndex < timeframes.length);
        bytes32 positionKey = keccak256(
            abi.encodePacked(categoryId, cityName, timeframes[timeframeIndex])
        );
        return records[playerAddr].actualValues[positionKey];
    }

    function resolvePosition(
        uint256 categoryId,
        string calldata cityName,
        uint256 timeframeIndex
    ) public {
        // NOTES:
        // Position key = hash(categoryId + cityName + timeframe)
        // Oracle key = hash(categoryId + cityName)
        require(timeframeIndex < timeframes.length);
        Player storage player = records[msg.sender];
        require(
            validCitiesRealEstate[cityName] == true ||
                validCitiesAirQuality[cityName] == true,
            "City is not valid"
        );
        // value is the index value during the guessed time
        if (player.initialized == false) {
            newPlayer(msg.sender);
        }

        bytes32 positionKey = keccak256(
            abi.encodePacked(categoryId, cityName, timeframes[timeframeIndex])
        );
        Position storage pos = player.positions[positionKey]; // get the position that is being resolved
        //uint256 currentValue = oracle.getValue(categoryId, cityName); // Get the current value for the index
        uint256 currentValue = player.actualValues[positionKey];
        require(currentValue != 0, "Actual value has not been updated");
        require(
            pos.resolvable,
            "Position is not resolvable (actual value not yet updated)"
        );
        require(pos.isOccupied, "Position does not exist");
        require(
            currentTime >= (HOUR * pos.timeframe) + pos.timestamp,
            "Timeframe has not passed yet. Cannot resolve"
        );
        bool won;
        if (pos.long == true) {
            // player tried to long
            if (currentValue > pos.guess) {
                addWin(msg.sender);
                won = true;
            } else {
                addLoss(msg.sender);
                won = false;
            }
        } else {
            // player tried to short
            if (currentValue < pos.guess) {
                addWin(msg.sender);
                won = true;
            } else {
                addLoss(msg.sender);
                won = false;
            }
        }
        pos.isOccupied = false;
        pos.resolvable = false;

        // emit event
        emit PositionResolved(
            msg.sender,
            pos.oracleKey,
            pos.guess,
            pos.long,
            pos.timeframe,
            pos.timestamp,
            won,
            currentValue
        );
    }

    function updateTime(uint256 newTimestamp) public onlyOwner {
        require(newTimestamp >= currentTime, "No timetravel");
        currentTime = newTimestamp;
    }

    function addWin(address addr) private {
        Player storage player = records[addr];
        require(player.initialized == true);
        player.wins++;
    }

    function addLoss(address addr) private {
        Player storage player = records[addr];
        require(player.initialized == true);
        player.losses++;
    }

    function getActivePlayers() public view returns (address[] memory) {
        return players;
    }

    function getPlayerWinLoss(
        address addr
    ) public view returns (uint256 wins, uint256 losses) {
        Player storage player = records[addr];
        return (player.wins, player.losses);
    }

    function getPlayerPositions(
        address addr,
        string calldata cityName,
        uint256 categoryId
    ) public view returns (Position[] memory) {
        require(
            validCitiesRealEstate[cityName] == true ||
                validCitiesAirQuality[cityName] == true,
            "City is not valid"
        );
        require(
            categoryId == 1 || categoryId == 0,
            "invalid category (has to be 1 or 0)"
        );
        Player storage player = records[addr];
        // calculate position key for each timeframe and each city
        Position[] memory ret = new Position[](player.numPositions);

        uint256 index = 0;
        for (uint256 i = 0; i < timeframes.length; ) {
            bytes32 positionKey = keccak256(
                abi.encodePacked(categoryId, cityName, timeframes[i])
            );
            if (player.positions[positionKey].isOccupied) {
                ret[index] = player.positions[positionKey];
                unchecked {
                    index++;
                }
            }
            unchecked {
                i++;
            }
        }

        return ret;
    }
}
