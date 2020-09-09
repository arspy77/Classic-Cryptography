# Classic-Cryptography

A Desktop App to simulate various classic cryptography using Go with Fyne.

## Libraries Used
Below are the libraries used in this project:
1. Fyne
2. Gonum

## The Keys of each Algorithm

There are 8 Algorithm that can be tried in the application wiht various rules for the keys:

### 1. Vigenere Cipher

The key must be made up of only latin alphabets, any other symbol inputted will be ignored.

### 2. Auto Key Vigenere

The key must be made up of only latin alphabets, any other symbol inputted will be ignored.

### 3. Full Vigenere

The key must be made up of only latin alphabets, any other symbol inputted will be ignored.

The Alphabet field must be filled with all 26 latin alphabets with all letters appearing exactly once and without any other symbol in between.

Alphabet field example: qwertyuiopasdfghjklzxcvbnm 

### 4. Extended Vigenere

The key can be made up of any of the 256 ASCII symbol.

### 5. Playfair Cipher

The key must be made up only latin characters except for 'J'. Any other symbol inputted will be ignored.

### 6. Super Encryption

The key must be made up only latin characters. Any other symbol inputted will be ignored.

The Number field must be filled with a positive number.

### 7. Hill Cipher

The key must represent a m * m matrix with the determinant of the matrix not zero and not divisible by 2 or 13. Below is an example on how to input the matrix:

1|2,3|5

Each '|' represents a border between columns and ',' between rows. So the above notation represents a 2 * 2 matrix with 1 and 2 in the first row and 3 and 5 in the second row.

### 8. Affine Cipher

The M key must be relative prime with 26
The B key can be any integer