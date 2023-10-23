package main
import(
	"log"
	//"fmt"
	"os"
	"strings"
	"bufio"
	//"strconv"
	"math/rand"
	"time"
	"context"

	"google.golang.org/grpc"
	pb "github.com/BeerJob/tdist2/proto"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	content, err := os.Open("DATA.txt")
	if err != nil{
		log.Fatal("Fatal error 100: Unable to read file")
	}
	var lines []string
	scanner := bufio.NewScanner(content)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	randomPoint := rand.Intn(len(lines))
	virtualPoint := 0
	status := "MUERTO"
	for i:=0;i<len(lines);i++{
		virtualPoint = (randomPoint+i)%len(lines)
		if rand.Intn(99)<=45{
			status = "MUERTO"
		}else{
			status = "INFECTADO"
		}
		log.Printf("%s %s", lines[virtualPoint], status)
		conn, err := grpc.Dial("10.6.46.108:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con la OMS")
		}
		defer conn.Close()
		cliente := pb.NewNameNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := cliente.Recepcion_Info(ctx, &pb.Datos{Nombre: strings.SplitN(lines[virtualPoint]," ",2)[0],Apellido: strings.SplitN(lines[virtualPoint]," ",2)[1],Estado: status})
		if err != nil{
			log.Print("No hay respuesta de la OMS")
		}else{
			log.Printf("%s", r.Ok)
		}
		if i>=4{
			time.Sleep(3 * time.Second)
		}
	}
}
