package response

import "strings"

func Response(s string) (string, string, string){
	s = strings.ToLower(s)
	var s1 string = "" // msg 
	var s2 string = "" // path image
	var s3 string = "" // caption image
	// Menu
	if s == "menu" {
		s1 = 	"*Banco de Galicia* 			\n" + 
				"------------------- 		\n" +
			  	"1. Torre Central		\n" +
				"2. Plaza Galicia		\n" +
				"3. Casa Matriz			\n" +
				"4. Corrientes 411		\n" +
				"5. Peron 456			\n" +
				"6. Corrientes 409		\n" +
				"------------------- 		\n" +
				"Escoja el numero de opcion	\n" 
		s2 = ""
		s3 = ""
	}	
	if s == "1" {
		s1 = ""
		s2 = "img/image1.jpeg"
		s3 = "*Piso 1 | Torre Central* 			\n" + 
			  "------------------- 			\n" + 
			  "02 Puestos gerenciales		\n" + 
			  "14 Puestos de trabajo		\n" +
			  "02 Trabajadores de licencia		\n" +
			  "08 Puestos disponibles		\n" 
	}
       if s == "2" {
		s1 = ""
		s2 = "img/image2.jpeg"
                s3 = "*Piso 4 | Plaza Galicia*                  \n" +
                          "-------------------                  \n" +
                          "07 Puestos gerenciales               \n" +
                          "10 Puestos de trabajo                \n" +
                          "04 Trabajadores de licencia  	\n" +
                          "04 Puestos disponibles               \n"
        }
       if s == "3" {
		s1 = ""
                s2 = "img/image3.jpeg"
                s3 = "*Piso 8 | Casa Matriz*                    \n" +
                          "-------------------                  \n" +
                          "05 Puestos gerenciales               \n" +
                          "27 Puestos de trabajo                \n" +
                          "07 Trabajadores de licencia		\n" +
                          "02 Puestos disponibles               \n"
        }
       if s == "4" {
		s1 = ""
                s2 = "img/image4.jpeg"
                s3 = "*Piso 3 | Corrientes 411*                 \n" +
                          "-------------------                  \n" +
                          "03 Puestos gerenciales               \n" +
                          "14 Puestos de trabajo                \n" +
                          "00 Trabajadores de licencia  	\n" +
                          "06 Puestos disponibles               \n"
        }
       if s == "5" {
		s1 = ""
                s2 = "img/image5.jpeg"
                s3 = "*Piso 6 | Peron 456*			\n" +
                          "-------------------                  \n" +
                          "06 Puestos gerenciales               \n" +
                          "27 Puestos de trabajo                \n" +
                          "09 Trabajadores de licencia 		\n" +
                          "10 Puestos disponibles               \n"
        }
       if s == "6" {
		s1 = ""
                s2 = "img/image6.jpeg"
                s3 = "*Piso 4 | Corrientes 409*                 \n" +
                          "-------------------                  \n" +
                          "02 Puestos gerenciales               \n" +
                          "14 Puestos de trabajo                \n" +
                          "02 Trabajadores de licencia  	\n" +
                          "08 Puestos disponibles               \n"
        }
	return s1, s2, s3
}
