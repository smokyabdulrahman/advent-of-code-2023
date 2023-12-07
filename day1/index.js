const fs = require('fs');
const path = require('path');

function crackCalibrationDocument(documentContent) {
	return documentContent.split(
}



console.log(crackCalibrationDocument(fs.readSync(path.join(__dirname, "./input.txt"))));
