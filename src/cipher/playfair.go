package cipher

import (

	"strings"
)


func processKey( key string)string{
	var vKey string
	pKey := processString(key)
	vKey = ""
	for i:= 0;i<len(pKey);i++{
		if(string(pKey[i]) != "J"){
			if(!strings.Contains(vKey,string(pKey[i]))){
				vKey += string(pKey[i])
			}
		}
		
	}
	for i:= 65 ; i<=90;i++{
		if(i != 74){
			if(!strings.Contains(vKey,string(rune(i)))){
				vKey += string(rune(i))
			}
		}
		
	}
	return vKey
}

func keyToSquare(key string) [5][5]string{
	var square [5][5]string
	n:= 0
	for i:= 0;i<5;i++{
		for j:=0;j<5;j++{
			square[i][j]=string(key[n])
			n++
		}
	}
	return square
}

func processPlayfairPT(plainText string) string{
	vPlainText := ""
	pPlainText := processString(plainText)
	for i:= 0;i<len(pPlainText);i++{
		if(string(pPlainText[i]) == "J"){
			vPlainText += "I"
		}else{
			vPlainText+=string(pPlainText[i])
		}
	}
	
	for i:= 0;i<len(vPlainText);i++{
		if(string(vPlainText[i])!= "X"){
			if(i < len(vPlainText)-1){
				if(vPlainText[i]==vPlainText[i+1]){
					vPlainText = vPlainText[:i+1]+"X"+vPlainText[i+1:]
				}
			}
			
		}
		i++
	}
	if(len(vPlainText)%2 != 0){
		vPlainText+="X"
	}
	return vPlainText
}

func findLocation(square [5][5]string,char string) [2]int{
	x := [2]int{}
	for i:= 0;i<5;i++{
		for j:= 0;j<5;j++{
			if(char == square[i][j]){
				x = [2]int{i,j}
				return x
			}
		}
	}
	return x

}

func Playfair(plainText string,key string) string{
	squareKey := keyToSquare(processKey(key))
	pPlainText := processPlayfairPT(plainText)
	cipherText := ""

	for i:= 0;i<len(pPlainText);i++{
		char1 := findLocation(squareKey,string(pPlainText[i]))
		char2 := findLocation(squareKey,string(pPlainText[i+1]))
		if(char1[0]==char2[0]){
			char1[1] = (char1[1]+1)%5 
			char2[1] = (char2[1]+1)%5
			
			cipherText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		} else if(char1[1]==char2[1]){
			char1[0] = (char1[0]+1)%5 
			char2[0] = (char2[0]+1)%5
			
			cipherText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		} else{
			kol1 := char1[1]
			kol2 := char2[1]
			char1[1] = kol2
			char2[1] = kol1
			
			cipherText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		}
		i++
	}
	return cipherText
}
func DecipherPlayfair(cipherText string,key string) string{
	squareKey := keyToSquare(processKey(key))
	plainText:= ""
	for i:= 0;i<len(cipherText);i++{
		char1 := findLocation(squareKey,string(cipherText[i]))
		char2 := findLocation(squareKey,string(cipherText[i+1]))
		if(char1[0]==char2[0]){
			if(char1[1]-1 < 0){
				char1[1] = (char1[1]-1)+5 
			}else{
				char1[1] = (char1[1]-1)%5 
			}
			if(char2[1]-1 < 0){
				char2[1] = (char2[1]-1)+5 
			}else{
				char2[1] = (char2[1]-1)%5
			}
			plainText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		} else if(char1[1]==char2[1]){
			
			if(char1[0]-1 < 0){
				char1[0] = (char1[0]-1)+5 
			}else{
				char1[0] = (char1[0]-1)%5 
			}
			if(char2[0]-1 < 0){
				char2[0] = (char2[0]-1)+5 
			}else{
				char2[0] = (char2[0]-1)%5
			}
			plainText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		} else{
			kol1 := char1[1]
			kol2 := char2[1]
			char1[1] = kol2
			char2[1] = kol1
			
			plainText += squareKey[char1[0]][char1[1]] + squareKey[char2[0]][char2[1]]
		}
		i++
	}
	return plainText
}