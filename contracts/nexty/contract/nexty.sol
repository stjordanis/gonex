pragma solidity 0.5.1;

contract Nexty {
    mapping(address => address) signers;
    address[] sealers;
    
    constructor() public {
        address earth = 0xD88D2AEB036D162db43E1B18EBD820cbD99d91b2;
        address sun = 0x1DeE87fc224a41c8699a3d8a8557918BE7e4795a;
        address moon = 0x8e2a6a6690156004185457c0ee70675Bc4C1B1A5;
        
        signers[sun] = 0x391730a41d2c27279b2B37e5E9209d5682C6d09a;
        signers[moon] = 0x766Ea022b264cd64Fb346b2B88409137790354b5;
        signers[earth] = 0x4DF4f0232bD4aC6808C98871959850EaC389013d;
        
        sealers.push(sun);
        sealers.push(moon);
        sealers.push(earth);
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
