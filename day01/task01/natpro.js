// JavaScript source code

var numberstring = "";
var numberarray;
var result = 0;

function calculate() {
    readNumbers();
    convertNumbers();
    handleNumbers();
}

function readNumbers() {
    numberstring = document.getElementById("numbers").innerText;
    //console.log(numberstring);
}


function convertNumbers() {
    numberarray = numberstring.split("\n");
    //console.log(numberarray);
}

function handleNumbers() {
    for (var line = 0; line < numberarray.length; line++) {
        //console.log(numberarray[line]);
        if (numberarray[line].charAt(0) == '+') {
            addNumber(splitNumberFromSign(numberarray[line]));
        } else if (numberarray[line].charAt(0) == '-') {
            minusNumber(splitNumberFromSign(numberarray[line]));
        }
    }
    console.log("Ergebnis:" + result);
}


function splitNumberFromSign(stringnumber) {
    var number = parseInt(stringnumber.substr(1));
    return number;
}

function addNumber(number) {
    result += number;
}

function minusNumber(number) {
    result -= number;
}
