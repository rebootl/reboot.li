import fs from 'fs';
import path from 'path';

// files/paths
const inFile = 'entries-in.json';
const outFile = 'entries.json';

// load data
let data = JSON.parse(fs.readFileSync(inFile, 'utf8'));

// write data
const writeData = (f) => {
  fs.writeFileSync(f, JSON.stringify(data));
};

for (const entry of data) {
  entry.deleted = false;
}

writeData(outFile);
