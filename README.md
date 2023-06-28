# Smart-Contract-Function-Selectors-
The featured Go codebase outputs in the cli the hex value of the function selectors present in a smart contract from the .abi file of the compiled .sol file <br>

After you have compiled your .sol file using  `solc`  in the cli, you might want to have the hexadecimal values of the different function selectors present 
in the smart contratct.<br>
Remember , for each non-anonymous functions declared in your smart contract ,  you have a function selector assigned to it. <br>

In an Ethereum based network, the solidity compiler  (in our case `solc` )  automatically generates the function selectors based on the function name and parameters types defined in the contract . <br>

When a function is called on a smart contract, the function selector is used to identify which function should be executed. 

To run the program :   `go run functionSelector.go`   
