// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Hello{
    string Msg;
    function setMsg(string memory _msg)public{
        Msg=_msg;
    }
    function getMsg()view public returns(string memory){
        return Msg;
    }
}