package response 

func Response(s string) (string, string, string){
	var s1 string = "" // msg 
	var s2 string = "" // path image
	var s3 string = "" // Caption image
	// Menu
	if s == "init" {
		s1 = 	"*Banco de Galicia* 	\n" + 
				"------------------- \n" +
			  	"1. Torre Central		\n" +
				"2. Plaza Galicia		\n" +
				"3. Casa Matriz		\n" +
				"4. Corrientes 411	\n" +
				"5. Peron 456			\n" +
				"6. Corrientes 409	\n" 
		s2 = ""
		s3 = ""
	}	
	if s == "1" {
		s1 = "*Torre Central* 			\n" + 
			  "------------------- 			\n" + 
			  "02 Puestos gerenciales		\n" + 
			  "14 Puestos de trabajo		\n" +
			  "02	Trabajadores de licencia\n" +
			  "08	Puestos disponibles		\n" 
		s2 = "img/image.jpeg"
		s3 = "Piso 1 | Torre Central"
	}
	return s1, s2, s3
}