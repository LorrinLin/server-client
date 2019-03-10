package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	//Dial to the server and get the connection
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("Error at Dial...", err)
		return
	}
	
	fmt.Println("client connect successfully!!",conn)
	reader := bufio.NewReader(os.Stdin)
	
	for{
		fmt.Println("Please input what you want to send, input 'exit' to exit...")
		// Reading by standard input
		
		line, err := reader.ReadString('\n')
	
		if err!= nil{
			fmt.Println("Error at read..",err)
		}
		
		//If "exit" inputed, exit
		line = strings.Trim(line, " \r\n")
		
		if  line == "exit"{
			fmt.Println("You are exited...")
			break
		}
	
		//send it to server
		n, err := conn.Write([]byte(line + "\n"))
		if err!= nil{
			fmt.Println("Error conn.write..", err)
		}
		fmt.Printf("Client sends data successful! ->%d\n", n)
		
	}
	
	
}