pragma solidity ^0.8.10;

contract Bank {
    struct UserAccount {
        uint256 staked;
        uint256 paidout;
        uint256 timestamp; //creation timestamp
        uint256 yieldRate; //annual
    }

    struct LoanAccount {
        uint256 loaned;
        uint256 paid;
        uint256 timestamp; //creation timestamp
        uint256 interestRate; //annual
    }

    address public creator;
    uint256 public accruementCycleLength = 1 * 60 * 60 * 24 * 30;
    uint256 public yieldRate = 1;
    uint256 public interestRate;

    mapping(address => UserAccount) UserAccounts;
    mapping(address => LoanAccount) LoanAccounts;

    //Logging Events
    event Stake(address indexed depositor, uint256 amount, uint256 yieldRate);
    event Payout(address indexed depositor, uint256 amount, uint256 pending);
    event Loan(address indexed loaner, uint256 amount, uint256 interestRate);
    event Repayment(address indexed loaner, uint256 amount, uint256 pending);

    constructor() {
        creator = msg.sender;
    }

    function adjustInterestRate(uint256 newrate) public {
        require(msg.sender == creator, "only creator can adjust interest rate");
        interestRate = newrate;
    }

    function adjustYieldRate(uint256 newrate) public {
        require(msg.sender == creator, "only creator can adjust yield rate");
        yieldRate = newrate;
    }

    function currentInterestRate() public view returns (uint256) {
        return interestRate;
    }

    function currentYieldRate() public view returns (uint256) {
        return yieldRate;
    }

    function currentAccruementLength() public view returns (uint256) {
        return accruementCycleLength;
    }

    function stake() public payable {
        if (UserAccounts[msg.sender].yieldRate == 0) {
            UserAccounts[msg.sender] = UserAccount(
                msg.value,
                0,
                block.timestamp,
                yieldRate
            );
        }
    }

    function payout() public {
        uint256 payout = payoutAmount();

        require(payout > 0, "nothing to payout");
        payable(msg.sender).transfer(payout);
    }

    // function loan(uint256 amount) public {
    //     if (!LoanAccountExists[msg.sender]) {
    //         LoanAccounts[msg.sender] = LoanAccount(
    //             amount,
    //             0,
    //             block.timestamp,
    //             interestRate,
    //             accruementCycleLength
    //         );
    //         LoanAccountExists[msg.sender] = true;
    //         payable(msg.sender).transfer(amount);
    //     }
    // }

    function payoutAmount() public view returns (uint256) {
        uint256 cycles = (block.timestamp -
            UserAccounts[msg.sender].timestamp) / accruementCycleLength;

        uint256 cyclesPerYear = (365 * 24 * 60 * 60) / accruementCycleLength;

        uint256 yield = (UserAccounts[msg.sender].yieldRate / cyclesPerYear) *
            cycles;

        uint256 totalPayout = (1 + yield / 100) *
            UserAccounts[msg.sender].staked;

        return totalPayout - UserAccounts[msg.sender].paidout;
    }
}
