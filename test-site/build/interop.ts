const esbuild = require("../../scripts/esbuild").installForTests();

const createResolver = (options) => {
  const { create } = require("enhanced-resolve");
  const resolveCache = new Map();
  const resolveSync = create.sync({
    aliasFields: ["browser"],
    conditionNames: ["import", "require", "node"],
    mainFields: options.mainFields,
    mainFiles: options.mainFiles,
    extensions: options.extensions,
    preferRelative: options.preferRelative,
    addMatchAll: false,
    alias: options.alias,
    plugins: [],
  });
  const node_resolve = (id, dir) => {
    if (resolveCache.get(id + dir)) {
      return resolveCache.get(id + dir);
    }
    let result;
    try {
      result = resolveSync({}, dir, id);
    } catch (err) {
      // hack, scss、less 需要支持~xxx/yyy从npm加载
      if (id.endsWith("scss") || id.endsWith("less")) {
        result = resolveSync({}, dir, id.replace(/^~/, ""));
      } else {
        throw err;
      }
    }
    resolveCache.set(id + dir, result);
    return result;
  };
  return node_resolve;
};

const resolver = createResolver({
  mainFields: ["lynx", "module", "main"],
  mainFiles: ["index.lynx", "index"],
  alias: {},
  extensions: [
    ".jsx",
    ".tsx",
    ".js",
    ".ts",
    ".css",
    ".less",
    ".sass",
    ".scss",
    ".json",
  ],
  preferRelative: false,
});

const AdapterPlugin = () => {
  return {
    name: "esbuild:adapter",
    setup(build) {
      build.onResolve({ filter: /.*/ }, async (args) => {
        console.log("onResolve: ", args);
        console.log("moduleType: ", args.path.includes("tslib") ? 6 : 0);
        return {
          // moduleType: args.path.includes("tslib") ? 6 : 0,
          moduleType: args.path === "tslib" ? 6 : 0,
          // external: args.path === "tslib",
          path: resolver(args.path, args.resolveDir),
          cacheDisable: true,
        };
      });
    },
  };
};

async function exec() {
  // interop test
  await esbuild.build({
    entryPoints: ["src/interop.ts"],
    outdir: "dist",
    platform: "neutral",
    bundle: true,
    watch: false,
    plugins: [AdapterPlugin()],
    format: "cjs",
    target: "es2019",
    tsconfig: require.resolve("../tsconfig.json"),
    external: ["tslib"],
  });
}

exec();
