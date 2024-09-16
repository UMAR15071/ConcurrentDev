class Roman {
    public static void main(String[] args){
        int number = romanNumber("III");
    }
    
}
public static int romanNumber(String roman){
    int i = 0;
    char pre = '\u0000';
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