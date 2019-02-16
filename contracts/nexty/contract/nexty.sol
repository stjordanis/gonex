pragma solidity ^0.5.2;

/**
 * @title SafeMath
 * @dev Unsigned math operations with safety checks that revert on error
 */
library SafeMath {
    /**
     * @dev Multiplies two unsigned integers, reverts on overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b);

        return c;
    }

    /**
     * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    /**
     * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;

        return c;
    }

    /**
     * @dev Adds two unsigned integers, reverts on overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);

        return c;
    }

    /**
     * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
     * reverts when dividing by zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
interface IERC20 {
    function transfer(address to, uint256 value) external returns (bool);

    function approve(address spender, uint256 value) external returns (bool);

    function transferFrom(address from, address to, uint256 value) external returns (bool);

    function totalSupply() external view returns (uint256);

    function balanceOf(address who) external view returns (uint256);

    function allowance(address owner, address spender) external view returns (uint256);

    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approval(address indexed owner, address indexed spender, uint256 value);
}

/**
 * @title Nexty sealers management smart contract
 */
contract NextyGovernance {
    using SafeMath for uint256;

    // minimum of deposited NTF to join
    uint256 public constant MIN_NTF_AMOUNT = 100;
    // minimum blocks number distance from last leaved to current chain blocknumber to withdrawable
    uint256 public constant LOCK_DURATION = 5 * 60;

    enum SealerStatus {
        PENDING_ACTIVE,     // Sealer deposited enough NTFs into registration contract successfully.

        ACTIVE,             // Sealer send request to become a sealer 
                            // and added into activation sealer set successfully

        PENDING_WITHDRAW,   // Sealer send request to exit from activation sealer set successfully. 
                            // Sealer casted out of activation sealer set

        WITHDRAWN,          // Sealer already withdrawn their deposit NTFs successfully. 
                            // They can only make withdrawal after withdrawal period.

        PENALIZED           //Sealer marked as penalized node (update by consensus or voting result via dapp) 
                            //and cannot become active sealer and cannot withdraw balance neither.
    }

    struct Record {
        SealerStatus status;
        //ntf amount deposited
        uint256 balance;
        //sealer used address to seal blocks
        address coinbase;
        //withdrawable time after leaving
        uint256 unlockTime;
    }

    // NTF token contract, unit used to join Nexty sealers
    IERC20 public token;

    // get invester address of a coinbase
    mapping(address => Record) public sealers;

    address[] public coinbases;
    // check if address of a coinbase
    mapping(address => bool) public isCoinbase;
    // get sealer's address of a coinbase address
    mapping(address => address) public getSealer;

    event Deposited(address _sealer, uint _amount);
    event Joined(address _sealer, address _coinbase);
    event Left(address _sealer, address _coinbase);
    event Withdrawn(address _sealer, uint256 _amount);
    event Banned(address _sealer);
    event Unbanned(address _sealer);

    /**
    * Check if address is a valid destination to transfer tokens to
    * - must not be zero address
    * - must not be the token address
    * - must not be the sender's address
    */
    modifier validCoinbase(address _coinbase) {
        require(!isCoinbase[_coinbase], "coinbase already used");
        require(_coinbase != address(0x0), "coinbase zero");
        require(_coinbase != address(this), "same contract's address");
        require(_coinbase != msg.sender, "same sender's address");
        _;
    }

    modifier notBanned() {
        require(sealers[msg.sender].status != SealerStatus.PENALIZED, "banned ");
        _;
    }

    modifier joinable() {
        require(sealers[msg.sender].status != SealerStatus.ACTIVE, "already joined ");
        require(sealers[msg.sender].balance >= MIN_NTF_AMOUNT, "not enough ntf");
        _;
    }

    modifier leaveable() {
        require(sealers[msg.sender].status == SealerStatus.ACTIVE, "not joined ");
        _;
    }

    modifier withdrawable() {
        require(isWithdrawable(msg.sender), "unable to withdraw at the moment");
        _;
    }

    /**
    * contract initialize
    */
    constructor(address _token, address[] memory _sealers) public {
        token = IERC20(_token);
        for (uint i = 0; i < _sealers.length; i++) {
            coinbases.push(_sealers[i]);
            isCoinbase[_sealers[i]] = true;
            getSealer[_sealers[i]] = _sealers[i];
        }        
    }

    // Get ban status of a sealer's address
    function isBanned(address _address) public view returns(bool) {
        return (sealers[_address].status == SealerStatus.PENALIZED);
    }

    ////////////////////////////////

    function addCoinbase(address _coinbase) internal {
        isCoinbase[_coinbase] = true;
        coinbases.push(_coinbase);
    }

    function removeCoinbase(address _coinbase) internal {
        isCoinbase[_coinbase] = false;
        for (uint i = 0; i < coinbases.length; i++) {
            if (_coinbase == coinbases[i]) {
                coinbases[i] = coinbases[coinbases.length - 1];
                coinbases.length--;
                return;
            }
        }
    }

    /**
    * Transfer the NTF from token holder to registration contract. 
    * Sealer might have to approve contract to transfer an amount of NTF before calling this function.
    * @param _amount NTF Tokens to deposit
    */
    function deposit(uint256 _amount) public returns (bool) {
        token.transferFrom(msg.sender, address(this), _amount);
        sealers[msg.sender].balance = (sealers[msg.sender].balance).add(_amount);
        emit Deposited(msg.sender, _amount);
        return true;
    }
    
    /**
    * To allow deposited NTF participate joining in as sealer. 
    * Participate already must deposit enough NTF via Deposit function. 
    * It takes coinbase as parameter.
    * @param _coinbase Destination address
    */
    function join(address _coinbase) public notBanned joinable validCoinbase(_coinbase) returns (bool) {
        sealers[msg.sender].coinbase = _coinbase;
        sealers[msg.sender].status = SealerStatus.ACTIVE;
        getSealer[_coinbase] = msg.sender;
        addCoinbase(_coinbase);
        emit Joined(msg.sender, _coinbase);
        return true;
    }

    /**
    * Request to exit out of activation sealer set
    */
    function leave() public notBanned leaveable returns (bool) {
        address _coinbase = sealers[msg.sender].coinbase;

        sealers[msg.sender].coinbase = 0x0000000000000000000000000000000000000000;
        sealers[msg.sender].status = SealerStatus.PENDING_WITHDRAW;
        sealers[msg.sender].unlockTime = LOCK_DURATION.add(block.timestamp);
        delete getSealer[_coinbase];
        removeCoinbase(_coinbase);
        emit Left(msg.sender, _coinbase);
        return true;
    }

    /**
    * To withdraw sealer’s NTF balance when they already exited and after withdrawal period.
    */
    function withdraw() public notBanned withdrawable returns (bool) {
        uint256 amount = sealers[msg.sender].balance;
        sealers[msg.sender].balance = 0;
        sealers[msg.sender].status = SealerStatus.WITHDRAWN;        
        token.transfer(msg.sender, amount);
        emit Withdrawn(msg.sender, amount);
        return true;
    }

    function getStatusCode(SealerStatus _status) private pure returns(uint256){
        if (_status == SealerStatus.PENDING_ACTIVE) return 0;
        if (_status == SealerStatus.ACTIVE) return 1;
        if (_status == SealerStatus.PENDING_WITHDRAW) return 2;
        if (_status == SealerStatus.WITHDRAWN) return 3;
        return 127;
    }

    function getStatus(address _address) public view returns(uint256) {
        return getStatusCode(sealers[_address].status);
    }

    function getBalance(address _address) public view returns(uint256) {
        return sealers[_address].balance;
    }  

    function getCoinbase(address _address) public view returns(address) {
        return sealers[_address].coinbase;
    }  

    function getUnlockTime(address _address) public view returns(uint256) {
        return sealers[_address].unlockTime;
    }

    function isWithdrawable(address _address) public view returns(bool) {
        return  
        (sealers[_address].status != SealerStatus.ACTIVE) &&
        (sealers[_address].status != SealerStatus.PENALIZED) &&
        (sealers[_address].unlockTime < block.timestamp);
    }
}