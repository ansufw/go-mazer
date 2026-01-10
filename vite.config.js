import { defineConfig, normalizePath } from 'vite'
import path, { resolve } from 'path'
import { viteStaticCopy } from 'vite-plugin-static-copy';

// Modules and extensions
// If the value is true, then it will copy the files inside the `dist` folders
// But if the value is false, it will copy the entire module files and folders
const modulesToCopy = {
    "@icon/dripicons": false, // With dist folder = false
    "@fortawesome/fontawesome-free": false,
    "apexcharts": true,
    "perfect-scrollbar": true,
    "sweetalert2": true,
    "feather-icons": true,
    "flatpickr": true,
    "filepond": true,
    "filepond-plugin-file-validate-size": true,
    "filepond-plugin-file-validate-type": true,
    "filepond-plugin-image-crop": true,
    "filepond-plugin-image-exif-orientation": true,
    "filepond-plugin-image-filter": true,
    "filepond-plugin-image-preview": true,
    "filepond-plugin-image-resize": true,
    "dragula": true,
    "chart.js": true,
    "parsleyjs": true,
    "summernote": true,
    "jquery": true,
    "quill": true,
    "simple-datatables": true,
    "datatables.net": false,
    "datatables.net-bs5": false,
    "jsvectormap": true,
    "toastify-js": false,
    "rater-js": false,
    "choices.js": false,
    "tinymce": false,
};

const copyModules = Object.keys(modulesToCopy).map(moduleName => {
    const withDist = modulesToCopy[moduleName]
    return {
        src: normalizePath(resolve(__dirname, `./node_modules/${moduleName}${withDist ? '/dist' : ''}`)),
        dest: 'extensions',
        rename: moduleName
    }
})

export default defineConfig({
    publicDir: false,
    base: '/assets/compiled/',
    build: {
        emptyOutDir: true,
        outDir: resolve(__dirname, 'views/assets/compiled'),
        rollupOptions: {
            input: {
                app: resolve(__dirname, 'views/assets/js/app.js'),
                style: resolve(__dirname, 'views/assets/scss/app.scss'),
                dark: resolve(__dirname, 'views/assets/scss/themes/dark/app-dark.scss'),
                flag: resolve(__dirname, 'views/assets/scss/pages/flag.scss'),
                sweetalert2: resolve(__dirname, 'views/assets/scss/pages/sweetalert2.scss'),
                todo: resolve(__dirname, 'views/assets/scss/widgets/todo.scss'),
                chat: resolve(__dirname, 'views/assets/scss/widgets/chat.scss'),
                dripicons: resolve(__dirname, 'views/assets/scss/pages/dripicons.scss'),
            },
            output: {
                entryFileNames: `js/[name].js`,
                chunkFileNames: `js/[name].js`,
                assetFileNames: (assetInfo) => {
                    const info = assetInfo.name.split('.');
                    const extType = info[info.length - 1];
                    if (['woff', 'woff2', 'ttf', 'otf', 'eot'].includes(extType)) {
                        return `css/fonts/[name][extname]`;
                    }
                    return `css/[name][extname]`;
                },
            }
        }
    },
    plugins: [
        viteStaticCopy({
            targets: [
                ...copyModules,
                {
                    src: normalizePath(resolve(__dirname, 'node_modules/bootstrap-icons/bootstrap-icons.svg')),
                    dest: 'static/images'
                }
            ]
        })
    ]
})
