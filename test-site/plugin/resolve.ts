import type {Plugin} from '../../lib/shared/types';
import path from 'path';

let resolvePlugin: Plugin = {
  name: 'test_resolve',
  setup(build) {
    build.onResolve({filter: /.*/}, args => {
      // const file_paths = [
      //   path.resolve(__dirname, "../src/index.tsx"),
      //   path.resolve(__dirname, "../src/component/Child/index.tsx")
      // ];

      // file_paths.forEach(f => {
      //   fs.watch(f, async () => {
      //     await build.resolve(args.path, {
      //       importer: args.importer,
      //       resolveDir: args.resolveDir,
      //       kind: args.kind,
      //       namespace: args.namespace,
      //     });
      //   })
      // })

      if (args.path.indexOf(".tsx") > -1) {
        return {
          path: path.resolve(__dirname, "../src/index.tsx"),
          namespace: "file",
          cacheDisable: true
        }
      }
      return undefined;
    })
  },
}

export {resolvePlugin};
