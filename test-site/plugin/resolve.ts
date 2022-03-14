import type {Plugin} from '../../lib/shared/types';
import path from 'path';

let resolvePlugin: Plugin = {
  name: 'test_resolve',
  setup(build) {
    build.onResolve({filter: /.*/}, args => {
      if (args.path.indexOf(".tsx") > -1) {
        console.log("test_resolve", args.path);
        return {
          path: path.resolve(__dirname, "../src/index.tsx"),
          namespace: "file",
          cacheEnable: true
        }
      }
      return undefined;
    })
  },
}

export {resolvePlugin};