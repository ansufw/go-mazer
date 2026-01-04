import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const modulesToCopy = {
    "apexcharts": true,
    "perfect-scrollbar": true,
    "sweetalert2": true,
    "feather-icons": true,
};

const filesToCopy = [
    {
        src: 'node_modules/bootstrap-icons/bootstrap-icons.svg',
        dest: 'views/assets/static/images/bootstrap-icons.svg'
    }
];

// Define base paths
// Assumes script is run from project root or processed relative to it
const rootDir = process.cwd();
const nodeModulesDir = path.join(rootDir, 'node_modules');
const destBaseDir = path.join(rootDir, 'views/assets/extensions');

// Ensure destination base exists
if (!fs.existsSync(destBaseDir)) {
    fs.mkdirSync(destBaseDir, { recursive: true });
}

Object.keys(modulesToCopy).forEach(moduleName => {
    const withDist = modulesToCopy[moduleName];
    const sourceDir = path.join(nodeModulesDir, moduleName, withDist ? 'dist' : '');
    const destDir = path.join(destBaseDir, moduleName);

    console.log(`Copying ${moduleName}...`);

    try {
        // Remove existing to ensure clean copy
        if (fs.existsSync(destDir)) {
            fs.rmSync(destDir, { recursive: true, force: true });
        }

        // Create destination directory
        fs.mkdirSync(destDir, { recursive: true });

        // Copy recursively
        // fs.cpSync is available in Node.js 16.7.0+
        fs.cpSync(sourceDir, destDir, { recursive: true });

        console.log(`✅ Copied ${moduleName} to ${destDir}`);
    } catch (err) {
        console.error(`❌ Failed to copy ${moduleName}:`, err.message);
    }
});

filesToCopy.forEach(file => {
    const sourcePath = path.join(rootDir, file.src);
    const destPath = path.join(rootDir, file.dest);
    const destDir = path.dirname(destPath);

    console.log(`Copying ${file.src} to ${file.dest}...`);

    try {
        // Ensure destination directory exists
        if (!fs.existsSync(destDir)) {
            fs.mkdirSync(destDir, { recursive: true });
        }

        // Copy file
        fs.copyFileSync(sourcePath, destPath);

        console.log(`✅ Copied ${path.basename(sourcePath)}`);
    } catch (err) {
        console.error(`❌ Failed to copy ${file.src}:`, err.message);
    }
});
