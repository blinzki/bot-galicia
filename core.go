package main

// PoC Galicia

import (
	// Local packages
        "github.com/blinzki/botgalicia/srv"

	// External packages
	"github.com/Baozisoftware/qrcode-terminal-go"
        "github.com/Rhymen/go-whatsapp"	 

	// Golanf packages
	"fmt"
	"log"
	"os"
	"time"
	"strings"
)

type waHandler struct {
	c *whatsapp.Conn
}

func (h *waHandler) HandleError(err error) {

	if e, ok := err.(*whatsapp.ErrConnectionFailed); ok {
		log.Printf("Connection failed, underlying error: %v", e.Err)
		log.Println("Waiting 30sec...")
		<-time.After(30 * time.Second)
		log.Println("Reconnecting...")
		err := h.c.Restore()
		if err != nil {
			log.Fatalf("Restore failed: %v", err)
		}
	} else {
		log.Printf("error occoured: %v\n", err)
	}
}

func (wh *waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	
	txt, path, cpt := response.Response(message.Text)
	timesend := int64(message.Info.Timestamp)
	timenow	:= time.Now().Unix() - 10
	
	if txt != "" && timesend > timenow {
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Text: txt,
		}
		msgId, err := wh.c.Send(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v", err)
			os.Exit(1)
		} else {
			fmt.Println("Message Sent -> ID : "+msgId)
		}
	}
	if path != "" && timesend > timenow{
		
		img, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cpt)
		msg := whatsapp.ImageMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Type:    "image/jpeg",
			Caption: cpt,
			Content: img,
		}

		msgId,err := wh.c.Send(msg)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v", err)
			os.Exit(1)		
		} else {
			fmt.Println("Message Sent -> ID : " + msgId)
			fmt.Println("Enviado a	  -> ID : " + message.Info.RemoteJid)
			fmt.Print("Hora 		  -> envio : ")
			fmt.Println(message.Info.Timestamp)
			fmt.Print("Hora 		  -> actual : ")
			fmt.Println(time.Now().Unix())			
			
		}
	}	
	fmt.Printf("%v %v %v %v\n\t%v\n", message.Info.Timestamp, message.Info.Id, message.Info.RemoteJid, message.Info.QuotedMessageID, message.Text)
}

func (*waHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	data, err := message.Download()
	if err != nil {
		return
	}
	filename := fmt.Sprintf("%v/%v.%v", os.TempDir(), message.Info.Id, strings.Split(message.Type, "/")[1])
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return
	}
	_, err = file.Write(data)
	if err != nil {
		return
	}
	log.Printf("%v %v\n\timage reveived, saved at:%v\n", message.Info.Timestamp, message.Info.RemoteJid, filename)
}

func main() {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		panic(err)
	}

	//Add handler
	wac.AddHandler(&waHandler{wac})
	
	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
	}
	fmt.Printf("login successful, session: %v\n", session)
	
	<-time.After(1000 * time.Second)

}

