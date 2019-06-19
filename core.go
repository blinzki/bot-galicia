package main

// Bot-galicia

import (
	"fmt"
	"log"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"time"
	"strings"
	"github.com/bot-galicia"
)

type waHandler struct {
	c *whatsapp.Conn
}

//HandleError needs to be implemented to be a valid WhatsApp handler
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

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (wh *waHandler) HandleTextMessage(message whatsapp.TextMessage) {

	if message.Text == "12345" {
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Text: saludo(1),
		}

	msgId, err := wh.c.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("Message Sent -> ID : "+msgId)
	}
		//sendImg(*waHandler)
		//fmt.Println("Pasa")
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


func sendImg(wac *whatsapp.Conn) {

	img, err := os.Open("image.jpg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}

		msg := whatsapp.ImageMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: "number@s.whatsapp.net",
			},
			Type:    "image/jpeg",
			Caption: "Hello Gopher!",
			Content: img,
		}

		msgId,err := wac.Send(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v", err)
			os.Exit(1)		
		} else {
			fmt.Println("Message Sent -> ID : "+msgId)
	}
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
	/*
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: "584166208443@s.whatsapp.net",
			},
			Text: "Plomo",
		}

	msgId, err := wac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("Message Sent -> ID : "+msgId)
	}
*/

}

