pragma solidity 0.5.1;

contract Nexty {
    mapping(address => address) signers;
    address[] sealers;
    
    constructor(address[] memory _sealers) public {
        for (uint i = 0; i < _sealers.length; i++) {
            signers[_sealers[i]] = _sealers[i];
            sealers.push(_sealers[i]);
        }
    }
    
    function join() public {
        sealers.push(msg.sender);
    }
    
    function leave() public {
        for (uint i = 0; i < sealers.length; i++) {
            if (msg.sender == sealers[i]) {
                sealers[i] = sealers[sealers.length - 1];
                delete sealers[sealers.length - 1];
                sealers.length--;
                return;
            }
        }
    }
}
