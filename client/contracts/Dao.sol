//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.9;

contract Dao {
    //a dao consist of memebres and proposals that are created by members
    //members can vote on proposals and the proposal can be executed if it has enough votes
    struct Tx {
        address to;
        uint amount;
        bool complete;
    }

    struct Proposal {
        uint id;
        Member proposer;
        string description;
        uint voteCount;
        uint txCount;
        mapping(address => bool) voters;
        mapping(uint => Tx) transactions;
        bool executed;
    }

    struct Member {
        bytes32 id;
        address wallet;
        string name;
        bool exists;
        uint contribution;
    }

    struct DaoStruct {
        bytes32 id;
        address creator;
        uint balance;
        string name;
        uint proposalCount;
        mapping(uint => Proposal) proposals;
        mapping(address => Member) members;
    }

    struct ViewDao {
        bytes32 id;
        address creator;
        uint balance;
        string name;
    }

    mapping(bytes32 => DaoStruct) daos;

    //functions
    //get dao by id
    function getDao(bytes32 id) private view returns (DaoStruct storage) {
        return daos[id];
    }

    //create a new dao
    function createDao(
        string memory dao_name,
        string memory creator_name
    ) public {
        bytes32 id = keccak256(abi.encodePacked(dao_name, msg.sender));
        DaoStruct storage dao = daos[id];
        dao.id = id;
        dao.creator = msg.sender;
        dao.name = dao_name;

        //add the creator as a member
        Member storage member = dao.members[msg.sender];
        member.id = keccak256(abi.encodePacked(creator_name, msg.sender));
        member.wallet = msg.sender;
        member.name = creator_name;
        member.exists = true;
    }

    //get dao by name
    function getDaoByName(
        string memory name
    ) public view returns (ViewDao memory) {
        bytes32 id = keccak256(abi.encodePacked(name, msg.sender));
        DaoStruct storage dao = daos[id];
        ViewDao memory viewDao;
        viewDao.id = dao.id;
        viewDao.creator = dao.creator;
        viewDao.balance = dao.balance;
        viewDao.name = dao.name;
        return viewDao;
    }

    //add a new member to the dao
    function addMember(
        bytes32 daoId,
        address wallet,
        string memory name
    ) public {
        DaoStruct storage dao = getDao(daoId);
        require(dao.creator == msg.sender, "only the creator can add members");
        Member storage member = dao.members[wallet];
        member.id = keccak256(abi.encodePacked(name, wallet));
        member.wallet = wallet;
        member.name = name;
        member.exists = true;
    }

    //get member by wallet
    function getMember(
        bytes32 daoId,
        address wallet
    ) public view returns (Member memory) {
        DaoStruct storage dao = getDao(daoId);
        Member storage member = dao.members[wallet];
        return member;
    }

    //create a new proposal
    //emit event proposal created with the proposal id, description and proposer

    event ProposalCreated(
        uint proposalId,
        string description,
        address proposer
    );

    function createProposal(bytes32 daoId, string memory description) public {
        DaoStruct storage dao = getDao(daoId);
        require(
            dao.members[msg.sender].exists,
            "only members can create proposals"
        );
        require(
            dao.members[msg.sender].contribution > 0,
            "only members with a contribution can create proposals"
        );
        Proposal storage proposal = dao.proposals[dao.proposalCount];
        proposal.id = dao.proposalCount;
        proposal.proposer = dao.members[msg.sender];
        proposal.description = description;
        proposal.txCount = 0;
        proposal.voteCount = 0;
        proposal.executed = false;
        dao.proposalCount++;

        emit ProposalCreated(proposal.id, proposal.description, msg.sender);
    }

    //get proposal by id
    function getProposal(
        bytes32 daoId,
        uint proposalId
    )
        public
        view
        returns (
            uint id,
            string memory description,
            uint voteCount,
            uint txCount,
            bool executed,
            address proposer
        )
    {
        DaoStruct storage dao = getDao(daoId);
        Proposal storage proposal = dao.proposals[proposalId];
        return (
            proposal.id,
            proposal.description,
            proposal.voteCount,
            proposal.txCount,
            proposal.executed,
            proposal.proposer.wallet
        );
    }

    //add transactions to a proposal
    function addTransaction(
        bytes32 daoId,
        uint proposalId,
        address[] memory to,
        uint[] memory amount
    ) public {
        DaoStruct storage dao = getDao(daoId);
        require(dao.members[msg.sender].exists);
        Proposal storage proposal = dao.proposals[proposalId];
        require(
            proposal.proposer.wallet == msg.sender,
            "only the proposer can add transactions"
        );
        for (uint i = 0; i < to.length; i++) {
            Tx storage txi = proposal.transactions[i];
            txi.to = to[i];
            txi.amount = amount[i];
        }
    }

    //vote on a proposal
    function vote(bytes32 daoId, uint proposalId) public {
        DaoStruct storage dao = getDao(daoId);
        require(
            dao.members[msg.sender].exists,
            "only members can vote on proposals"
        );
        Proposal storage proposal = dao.proposals[proposalId];
        require(
            !proposal.voters[msg.sender],
            "only one vote per member is allowed"
        );
        proposal.voters[msg.sender] = true;
        proposal.voteCount++;
    }

    //execute a proposal
    event ProposalExecuted(
        uint proposalId,
        string description,
        address proposer
    );

    function executeProposal(bytes32 daoId, uint proposalId) public {
        DaoStruct storage dao = getDao(daoId);
        require(
            dao.members[msg.sender].exists,
            "only members can execute proposals"
        );
        Proposal storage proposal = dao.proposals[proposalId];
        require(
            proposal.voteCount > dao.proposalCount / 2,
            "the proposal needs more than 50% of the votes to be executed"
        );
        require(!proposal.executed, "the proposal has already been executed");

        for (uint i = 0; i < proposal.txCount; i++) {
            Tx storage txi = proposal.transactions[i];
            require(
                txi.amount <= dao.balance,
                "the dao does not have enough funds to execute the proposal"
            );
            payable(txi.to).transfer(txi.amount);
            dao.balance -= txi.amount;
            txi.complete = true;
        }

        proposal.executed = true;
        emit ProposalExecuted(
            proposal.id,
            proposal.description,
            proposal.proposer.wallet
        );
    }
}
