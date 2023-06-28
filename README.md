# Smart-Contract-Function-Selectors-
The featured Go codebase outputs in the cli the hex value of the function selectors present in a smart contract from the .abi file of the compiled .sol file <br>

After you have compiled your .sol file using  `solc`  in the cli, you might want to have the hexadecimal values of the different function selectors present 
in the smart contract.<br>
Remember , for each declared function in your smart contract , there is one function selector. <br>
The value of the function selector is determined during compilation time. <br>
The function selector is a 4-byte hash (derived from a sha3 Algorithm)  value that represents the function signature, and it is calculated based on the function name and its parameter types.


The function selectors are then included in the compiled bytecode, and they are used to identify and differentiate between different functions when you interact with the contract on the Ethereum blockchain. The function selector acts as a unique identifier for each function in the contract, allowing the EVM (Ethereum Virtual Machine) to correctly route and execute function calls based on their selectors.

To run the program :   `go run functionSelector.go`   
