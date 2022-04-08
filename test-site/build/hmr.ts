const esbuild = require('../../scripts/esbuild').installForTests()

async function exec() {
  await esbuild.build({
    entryPoints: ["src/entry.ts"],
    outfile: "dist/main.js",
    minify: false,
    bundle: true,
    watch: true,
    sourcemap: true,
    platform: "node",
    treeShaking: false,
    plugins: []
  });
}

exec();