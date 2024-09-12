#!/usr/bin/env node

import { exec } from "child_process";
import { fileURLToPath } from "url";
import path from "path";

exec("go version", (error, _, __) => {
  if (error) {
    console.error(`Please install Go first!`);
    return;
  }
});

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const goFilePath = path.join(__dirname, "main.go");

exec(`go run ${goFilePath}`, (error, stdout, stderr) => {
  if (error) {
    console.error(`Error: ${error.message}`);
    return;
  }
  if (stderr) {
    console.error(`stderr: ${stderr}`);
    return;
  }
  console.log(`CSS reset has been added to your project! 🎉`);
});
