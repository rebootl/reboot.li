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
  if (!entry.deleted) {
    entry.deleted = false;
  }
  if (!entry.version) {
    entry.version = 1;
  }
  if (!entry.last) {
    entry.last = true;
  }
}

writeData(outFile);
