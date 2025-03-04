import {loadPlugin, resolvePlugin} from '../plugin'
import fs from 'fs';
import path from 'path';

const esbuild = require('../../scripts/esbuild').installForTests()
// const esbuild = require('../../npm/esbuild/lib/main')

async function exec() {

  let res = await esbuild.build({
    entryPoints: ["src/index.tsx"],
    outfile: "dist/main.js",
    minify: false,
    bundle: true,
    // watch: true,
    incremental: true,
    sourcemap: true,
    plugins: [loadPlugin, resolvePlugin]
  });

  const file_paths = [
    path.resolve(__dirname, "../src/index.tsx"),
    path.resolve(__dirname, "../src/component/Child/index.tsx")
  ];

  file_paths.forEach(f => {
    fs.watch(f, () => {
      res.rebuild(file_paths);
    })
  })

}

exec();

