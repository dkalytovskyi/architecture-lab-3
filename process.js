const path = require('path');
const fs = require('fs');

const consoleArgs = process.argv.slice(2);

const examplesPath = path.join(__dirname, consoleArgs[0].substring(2) || 'examples');
const outputPath = path.join(__dirname, consoleArgs[1].substring(2) || 'output-js');

const arrayOfPromises = [];

function countFileLines(filePath) {
    return new Promise((resolve, reject) => {
        let lineCount = 0;
        fs.createReadStream(filePath)
        .on('data', (buffer) => {
            let idx = -1;
            do {
                idx = buffer.indexOf(10, idx+1);
                lineCount++;
            } while (idx !== -1);
        }).on('end', () => {
            resolve(lineCount);
        }).on('error', reject);
    });
}

function writeFiles(filePath, data) {
    return new Promise((resolve, reject) => {
        fs.writeFile(filePath, data.toString(), function (err) {
            if (err) reject;
            resolve();
        });
    });
}

function filesIterator(files, outputPath) {
    files.forEach(function (file) {
        arrayOfPromises.push(
            countFileLines(examplesPath + '/' + file).then(data => {
                writeFiles(outputPath + '/' + file.split('.')[0] + '.res', data.toString());
            })
        );
    });
}

function getDirectoryFiles(examplesPath) {
    return new Promise((resolve, reject) => {
        fs.readdir(examplesPath, function (err, files) {
            if (err) {
                reject;
            };
            try {
                fs.statSync(outputPath);
            }
            catch (err) {
                if (err.code === 'ENOENT') {
                    fs.mkdirSync(consoleArgs[1]);
                };
            };
            filesIterator(files, outputPath);
            resolve();
        });
    });
}

getDirectoryFiles(examplesPath).then(() => {
    Promise.all(arrayOfPromises).then(data => {
        console.log('Total number of processed files: ' + data.length);
    });
});
