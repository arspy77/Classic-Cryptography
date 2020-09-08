package cipher

import (
	"fmt"
	
	"math"
	
)

func Transpose(plainText string,number int)string{
	text := processString(plainText)
	var arrayText [][]string
	var arrayChar []string
	cipherText := ""
	for i:=0;i<len(text);i++{
		if(i>0 && i % number == 0){
			arrayText= append(arrayText,arrayChar)
			arrayChar = nil
		}
		arrayChar = append(arrayChar,string(text[i]))
		
		if(i+1 ==len(text)){
			arrayText = append(arrayText,arrayChar)
		}
	}
	for j:= 0;j<number;j++{
		for i:= 0;i<len(arrayText);i++{
			if(len(arrayText[i])>j){
				cipherText+=arrayText[i][j]
			}
		}
	}
	
	
	return cipherText
}
func DecipherTranspose(cipherText string,number int)string{
	var arrayText [][]string
	var arrayChar []string
	no := int(math.Ceil( float64(len(cipherText)) / float64(number)))
	
	plainText := ""
	baris:= 0
	kolom:=0
	for i:=0;i<len(cipherText);i++{
		if(kolom>=no){
			arrayText= append(arrayText,arrayChar)
			arrayChar = nil
			baris++
			kolom=0
		}else if(baris+(kolom*number)>=len(cipherText)){
			arrayText= append(arrayText,arrayChar)
			arrayChar = nil
			baris++
			kolom=0
			
		}
		arrayChar = append(arrayChar,string(cipherText[i]))
		kolom++
		if(i+1 ==len(cipherText)){
			arrayText = append(arrayText,arrayChar)
		}
	}
	fmt.Println(arrayChar)
	for j:= 0;j<no;j++{
		for i:= 0;i<len(arrayText);i++{
			if(len(arrayText[i])>j){
				plainText+=arrayText[i][j]
			}
		}
	}
	
	return plainText
}

func SuperEncryption(plainText string,key string,number int)string{
	vignereEnc:= Vigenere(plainText,key)
	return Transpose(vignereEnc,number)
}
func DecipherSuperEncryption(cipherText string,key string,number int)string{
	transposePlaintext := DecipherTranspose(cipherText,number)
	return DecipherVigenere(transposePlaintext,key)
}