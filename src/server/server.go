package main
import(
	"net"
	"fmt"
)

var(
	
)

func process(conn net.Conn){
	
	defer conn.Close()
	//print out the address of connected client
		fmt.Println("I was connected by client,and waiting for your message-->",conn.RemoteAddr().String())
	
	//receive data from client
	for{
		buf := make([]byte,512)
		//Reading data from conn
		n,err := conn.Read(buf)
		if err!=nil{
			fmt.Println("error at read..", err)
			return
		}
		
		fmt.Print("From client: " + string(buf[:n]))
	}
}

func main(){
	
	fmt.Println("Server listen on...")
	//open port NO.8888 listening 
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("error at listen...",err)
		return
	}
	//listen will be closed when exist
	defer listen.Close()
	//Waiting for the client to connect me ...
	for{
		fmt.Println("Waiting for the client to connect me ..")
		conn, err := listen.Accept()
		
		if err != nil{
		fmt.Println("error at Accept()..",err)
		}
		
		go process(conn)
		
	}
	
	
}
