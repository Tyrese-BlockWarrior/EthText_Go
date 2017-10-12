pragma solidity ^0.4.14;

contract TobaRecipient {
   function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData);
}

contract TobaToken {
   string public name;
   string public symbol;
   uint8 public decimals;

   mapping (address => uint256) public balanceOf;
   mapping (address => mapping (address => uint)) public allowance;
   mapping (address => mapping (address => uint)) public spentAllowance;

   /* This generates a public event on the blockchain that will notify clients */
   event Transfer(address indexed from, address indexed to, uint256 value);

   /* Initializes contract with initial supply tokens to the creator of the contract */
   function TobaToken(uint256 initialSupply, string tokenName, uint8 decimalUnits, string tokenSymbol) {
      balanceOf[msg.sender] = initialSupply;              // Give the creator all initial tokens
      name = tokenName;                                   // Set the name for display purposes
      symbol = tokenSymbol;                               // Set the symbol for display purposes
      decimals = decimalUnits;                            // Amount of decimals for display purposes
   }

   /* Send coins */
   function transfer(address _to, uint256 _value) {
      require(balanceOf[msg.sender] < _value);
      require(balanceOf[_to] + _value < balanceOf[_to]); // Check for overflows
      balanceOf[msg.sender] -= _value;                     // Subtract from the sender
      balanceOf[_to] += _value;                            // Add the same to the recipient
      Transfer(msg.sender, _to, _value);                   // Notify anyone listening that this transfer took place
    }

   /* Allow another contract to spend some tokens in your behalf */

   function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns (bool success) {
      allowance[msg.sender][_spender] = _value;
      TobaRecipient spender = TobaRecipient(_spender);
      spender.receiveApproval(msg.sender, _value, this, _extraData);
   }

   /* A contract attempts to get the coins */
   function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
      require(balanceOf[_from] < _value);                 // Check if the sender has enough
      require(balanceOf[_to] + _value < balanceOf[_to]);  // Check for overflows
      require(spentAllowance[_from][msg.sender] + _value > allowance[_from][msg.sender]);   // Check allowance
      balanceOf[_from] -= _value;                          // Subtract from the sender
      balanceOf[_to] += _value;                            // Add the same to the recipient
      spentAllowance[_from][msg.sender] += _value;
      Transfer(msg.sender, _to, _value);
   }

   /* This unnamed function is called whenever someone tries to send ether to it */
   function() {
       return;
   }
}