# Smart-Contract-Function-Selectors-

The featured Go codebase outputs in the cli the hex value of the function selectors present in a smart contract from the .abi file of the compiled .sol file <br>

After you have compiled your .sol file using `solc` in the cli, you might want to have the hexadecimal values of the different function selectors present
in the smart contract.<br>
Remember , for each declared function in your smart contract , there is one function selector. <br>
The value of the function selector is determined during compilation time. <br>
The function selector is a 4-byte hash (derived from a sha3 Algorithm) value that represents the function signature, and it is calculated based on the function name and its parameter types.

The function selectors are then included in the compiled bytecode, and they are used to identify and differentiate between different functions when you interact with the contract on the Ethereum blockchain. The function selector acts as a unique identifier for each function in the contract, allowing the EVM (Ethereum Virtual Machine) to correctly route and execute function calls based on their selectors.

To run the program : `go run functionSelector.go`

-Import Statements: Import necessary packages such as "encoding/json" for JSON parsing, "encoding/hex" for hexadecimal encoding, "fmt" for printing, "os" for file operations, and "log" for logging errors.
-Struct Definitions: Define Go structs to represent the structure of the ABI file and its contents. The Function struct represents a function in the smart contract, with fields for name, type, inputs, outputs, payable status, and state mutability. The Input and Output structs represent the input and output parameters of a function, respectively.
-Main Function: The main function is the entry point of the program. It reads the ABI file (Counter.abi in this case), unmarshals its contents into a slice of Function structs, iterates over each function, calculates the function selector using the calculateSelector function, and prints the resulting selectors.
-Calculate Selector Function: The calculateSelector function takes the name and input parameters of a function, concatenates them into a function signature, calculates the Keccak-256 hash of the signature, and extracts the first four bytes as the function selector. This selector is then encoded as a hexadecimal string and returned.
