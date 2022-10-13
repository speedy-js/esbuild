package bundler

import (
	"testing"

	"github.com/evanw/esbuild/internal/config"
)

var importstar_ts_suite = suite{
	name: "importstar_ts",
}

func TestExcludeExportsForEntryPoint(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.tsx": `
				export default class A extends Component<Props, State> {
					state = {
						count: 1000,
					};
				
					renderXXX = () => <view><text>asdasdasdasdasdasd</text></view>
				
					render() {
						return (
							<view style={{ display: "flex", flexDirection: "column" }}>
								<view
									style={{ width: "100rpx", height: "100rpx", backgroundColor: "red" }}
									bindtap={() => {
										this.setState(preState => ({ count: preState.count + 1 }));
									}}
								>
									<text>Plus 1</text>
								</view>
								<text>{this.state.count}</text>
								{this.renderXXX()}
							</view>
						);
					}
				}
			`,
		},
		entryPaths: []string{"/entry.tsx"},
		options: config.Options{
			Mode:                  config.ModeBundle,
			AbsOutputFile:         "/out.js",
			ExcludeExportForEntry: true,
			OutputFormat:          config.FormatCommonJS,
		},
	})
}

func TestExcludeExportsForEntryPointWithTwoClasses(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.tsx": `
				class Test extends Component {
					render(){
						return <block>
						<text>test1</text>
						<text>test2</text>
					</block>
					}
				}
				
				export default class Fake extends Component {
					state = {
						color: 'red',
						flag: false
					}
					render() {
						return (
							<block>
								<text>test</text>
								<Test />
							</block>
						);
					}
				}
			`,
		},
		entryPaths: []string{"/entry.tsx"},
		options: config.Options{
			Mode:                  config.ModeBundle,
			AbsOutputFile:         "/out.js",
			ExcludeExportForEntry: true,
			OutputFormat:          config.FormatCommonJS,
		},
	})
}

func TestCommonClassExport(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.tsx": `
				export default class A extends Component<Props, State> {
					state = {
						count: 1000,
					};
				
					renderXXX = () => <view><text>asdasdasdasdasdasd</text></view>
				
					render() {
						return (
							<view style={{ display: "flex", flexDirection: "column" }}>
								<view
									style={{ width: "100rpx", height: "100rpx", backgroundColor: "red" }}
									bindtap={() => {
										this.setState(preState => ({ count: preState.count + 1 }));
									}}
								>
									<text>Plus 1</text>
								</view>
								<text>{this.state.count}</text>
								{this.renderXXX()}
							</view>
						);
					}
				}
			`,
		},
		entryPaths: []string{"/entry.tsx"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
			OutputFormat:  config.FormatCommonJS,
		},
	})
}

func TestTSImportStarUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportImportStarUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportImportStarNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportImportStarCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarAsUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarAsNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarAsCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarExportStarCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
			"/bar.ts": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarCommonJSUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.ts": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarCommonJSCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.ts": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarCommonJSNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.ts": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarAndCommonJS(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				const ns2 = require('./foo')
				console.log(ns.foo, ns2.foo)
			`,
			"/foo.ts": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarNoBundleUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarNoBundleCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarNoBundleNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarMangleNoBundleUnused(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarMangleNoBundleCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSImportStarMangleNoBundleNoCapture(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestTSReExportTypeOnlyFileES6(t *testing.T) {
	importstar_ts_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.ts": `
				import * as ns from './re-export'
				console.log(ns.foo)
			`,
			"/re-export.ts": `
				export * from './types1'
				export * from './types2'
				export * from './types3'
				export * from './values'
			`,
			"/types1.ts": `
				export interface Foo {}
				export type Bar = number
				console.log('some code')
			`,
			"/types2.ts": `
				import {Foo} from "./type"
				export {Foo}
				console.log('some code')
			`,
			"/types3.ts": `
				export {Foo} from "./type"
				console.log('some code')
			`,
			"/values.ts": `
				export let foo = 123
			`,
			"/type.ts": `
				export type Foo = number
			`,
		},
		entryPaths: []string{"/entry.ts"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}
