const { log } = require('console');
const fs = require('fs');
const path = require('path');
const numbers = {
	'one': 1,
	'two': 2,
	'three': 3,
	'four': 4,
	'five': 5,
	'six': 6,
	'seven': 7,
	'eight': 8,
	'nine': 9
}

function crackCalibrationDocument(documentContent) {
	return documentContent.split('\n')
	.filter(line => line.length >= 1)
	.map(line => extractNumberFromLine(line))
	.reduce((total, current) => total + current, 0);
}

function extractNumberFromLine(line) {
	let i = 0,
	    j = line.length - 1;
	let firstNumber;
	let lastNumber;

	while (i <= j && ( !firstNumber || !lastNumber )) {
		firstNumber = getNumberIfAny(line, i);
		lastNumber = getNumberIfAny(line, j);
		if (!firstNumber) i++;
		if (!lastNumber) j--;
	}
	
	return parseInt(`${firstNumber}${lastNumber}`);
}

function getNumberIfAny(line, startIndex) {
	return isNumber(line.charAt(startIndex))
	? line.charAt(startIndex)
	: getNumberFromLetteralIfAny(line, startIndex);
}

function isNumber(c) {
  return c >= '0' && c <= '9';
}

function getNumberFromLetteralIfAny(line, startIndex) {
	return Object.keys(numbers)
	.filter(key => key.length + startIndex <= line.length)
	.map(key => numbers[line.substring(startIndex, key.length + startIndex)])
	.find(num => num)
}

const filePath = path.join(__dirname, "./input.txt");
const fileContent = fs.readFileSync(filePath, 'utf8');
console.log(crackCalibrationDocument(fileContent));
