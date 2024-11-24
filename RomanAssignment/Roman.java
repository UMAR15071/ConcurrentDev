/*
 * 
 * Name: Syed Umar
 * Student ID: C00278724
 * 
 * Logic: The solution in this file convert Roman numbers given in form of Strings converted into integers. The logic is that we will simply 
 * go through the String and check each character. The character we get represent a number eg: I = 1 and V = 5. We have to make sure
 * while going through each one of these characters we keep track of the previous character, so that incase we get IV which is equal to 4. When we 
 * get to I we add 1 but when we get to V instead of adding 5 we will add 3 because our previous character is I which makes IV as 4. This is the 
 * basic which cover all the possible test cases. 
 */

class Roman {
    public static void main(String[] args){
        int number = romanNumber("III");
    }
    
}
public static int romanNumber(String roman){
    int i = 0;
    char pre = '\u0000';//value for null character
    int number = 0;
    for(i = 0; i < roman.length(); i++){
        if(roman.charAt(i) == 'I'){
            number = number + 1;
            pre = 'I';
        }
        if(roman.charAt(i) == 'V'){
            if(pre == 'I'){
                number = number + 3;
            }
            else{
                number = number + 5;
            }
            pre = 'V';
        }
        if(roman.charAt(i) == 'X'){
            if(pre == 'I'){
                number = number + 8;
            }
            else{
                number = number + 10;
            }
            pre = 'X';
        }
        if(roman.charAt(i) == 'L'){
            if(pre == 'X'){
                number = number + 30;
            }
            else{
                number = number + 50;
            }
            pre = 'L';
        }
        if(roman.charAt(i) == 'C'){
            if(pre == 'X'){
                number = number + 80;
            }
            else{
                number = number + 100;
            }
            pre = 'C';
        }
        if(roman.charAt(i) == 'D'){
            if(pre == 'C'){
                number = number + 300;
            }
            else{
                number = number + 500;
            }
            pre = 'D';
        }
        if(roman.charAt(i) == 'M'){
            if(pre == 'C'){
                number = number + 800;
            }
            else{
                number = number + 1000;
            }
            pre = 'M';
        }
        return number;

    }
}