package response 

func Response(s string) (string, string, string){
	var s1 string = "" // msg 
	var s2 string = "" // path image
	var s3 string = "" // caption image
	// Menu
	if s == "init" {
		s1 = 	"*Banco de Galicia* 			\n" + 
				"------------------- 		\n" +
			  	"1. Torre Central		\n" +
				"2. Plaza Galicia		\n" +
				"3. Casa Matriz			\n" +
				"4. Corrientes 411		\n" +
				"5. Peron 456			\n" +
				"6. Corrientes 409		\n" 
		s2 = ""
		s3 = ""
	}	
	if s == "1" {
		s1 = "*Torre Central* 				\n" + 
			  "------------------- 			\n" + 
			  "02 Puestos gerenciales		\n" + 
			  "14 Puestos de trabajo		\n" +
			  "02 Trabajadores de licencia	\n" +
			  "08 Puestos disponibles		\n" 
		s2 = "img/image.jpeg"
		s3 = "Piso 1 | Torre Central"
	}
       if s == "2" {
                s1 = "*Plaza Galicia*                           \n" +
                          "-------------------                  \n" +
                          "07 Puestos gerenciales               \n" +
                          "10 Puestos de trabajo                \n" +
                          "04 Trabajadores de licencia  \n" +
                          "04 Puestos disponibles               \n"
                s2 = "img/image.jpeg"
                s3 = "Piso 1 | Plaza Galicia"
        }
       if s == "3" {
                s1 = "*Casa Matriz*                           \n" +
                          "-------------------                  \n" +
                          "05 Puestos gerenciales               \n" +
                          "27 Puestos de trabajo                \n" +
                          "07 Trabajadores de licencia  \n" +
                          "02 Puestos disponibles               \n"
                s2 = "img/image.jpeg"
                s3 = "Piso 1 | Casa Matriz"
        }
       if s == "4" {
                s1 = "*Corrientes 411*                           \n" +
                          "-------------------                  \n" +
                          "03 Puestos gerenciales               \n" +
                          "14 Puestos de trabajo                \n" +
                          "00 Trabajadores de licencia  \n" +
                          "06 Puestos disponibles               \n"
                s2 = "img/image.jpeg"
                s3 = "Piso 1 | Corrientes 411"
        }
       if s == "5" {
                s1 = "*Peron 456*                           \n" +
                          "-------------------                  \n" +
                          "06 Puestos gerenciales               \n" +
                          "27 Puestos de trabajo                \n" +
                          "09 Trabajadores de licencia  \n" +
                          "10 Puestos disponibles               \n"
                s2 = "img/image.jpeg"
                s3 = "Piso 1 | Peron 456"
        }
       if s == "6" {
                s1 = "*Corrientes 409*                           \n" +
                          "-------------------                  \n" +
                          "02 Puestos gerenciales               \n" +
                          "14 Puestos de trabajo                \n" +
                          "02 Trabajadores de licencia  \n" +
                          "08 Puestos disponibles               \n"
                s2 = "img/image.jpeg"
                s3 = "Piso 1 | Corrientes 409"
        }
	return s1, s2, s3
}
