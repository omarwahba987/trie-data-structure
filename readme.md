This is the trie program
for run the program on command line
	1- change directory to program file // cmd/cli
	2- open terminal
	3- type command "go mod tidy" 
	// for handle packges and import them
	4- type command "go mod vendor"
	// for run the program
	5- type command "go run main.go"
	6- follow the instructions to construct the trie and enter the number of 			persons you want to enter
		 
	7- enter the data (name, address)
	8- now you can by follwing the instruction in cmd to see the trie paths 
	        and **search** about address for specific person by name
	        
for run the program on server Apis
	1- change directory to program file // cmd/server
	2- open terminal
	3- type command "go mod tidy" 
	// for handle packges and import them
	4- type command "go mod vendor"
	// for run the program
	5- type command "go run main.go"
	6- when "server runs on port :  8081" then server is ready 
	7- open post man and for insert person to the tree 
		make a POST request with url localhost:8081/insert
		and its body 
		{
    			"name":"person_name",
    			"address":"person_address"
		} 
	for visualize the trie make a GET request with url localhost:8081/visualize
