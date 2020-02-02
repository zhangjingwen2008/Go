pragma solidity ^0.4.24;

contract Test{

    uint256 ui256=100;
    int8 i10 =10;

    function add() returns(uint256){            //整型：返回110
        return ui256 + uint256(i10);
    }

    function isEqual() returns(bool) {          //布尔：返回false
        return ui256 == uint256(i10);
    }
}