// This file was automatically generated by "compat-table.js"

package compat

type Engine uint8

const (
	Chrome Engine = iota
	Edge
	ES
	Firefox
	IE
	IOS
	Node
	Opera
	Safari
)

func (e Engine) String() string {
	switch e {
	case Chrome:
		return "chrome"
	case Edge:
		return "edge"
	case ES:
		return "es"
	case Firefox:
		return "firefox"
	case IE:
		return "ie"
	case IOS:
		return "ios"
	case Node:
		return "node"
	case Opera:
		return "opera"
	case Safari:
		return "safari"
	}
	return ""
}

type JSFeature uint64

const (
	ArbitraryModuleNamespaceNames JSFeature = 1 << iota
	ArraySpread
	Arrow
	AsyncAwait
	AsyncGenerator
	BigInt
	Class
	ClassField
	ClassPrivateAccessor
	ClassPrivateBrandCheck
	ClassPrivateField
	ClassPrivateMethod
	ClassPrivateStaticAccessor
	ClassPrivateStaticField
	ClassPrivateStaticMethod
	ClassStaticBlocks
	ClassStaticField
	Const
	DefaultArgument
	Destructuring
	DynamicImport
	ExponentOperator
	ExportStarAs
	ForAwait
	ForOf
	Generator
	Hashbang
	ImportAssertions
	ImportMeta
	Let
	LogicalAssignment
	NestedRestBinding
	NewTarget
	NodeColonPrefixImport
	NodeColonPrefixRequire
	NullishCoalescing
	ObjectAccessors
	ObjectExtensions
	ObjectRestSpread
	OptionalCatchBinding
	OptionalChain
	RestArgument
	TemplateLiteral
	TopLevelAwait
	TypeofExoticObjectIsObject
	UnicodeEscapes
)

func (features JSFeature) Has(feature JSFeature) bool {
	return (features & feature) != 0
}

var jsTable = map[JSFeature]map[Engine][]versionRange{
	ArbitraryModuleNamespaceNames: {
		Chrome:  {{start: v{90, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{87, 0, 0}}},
		Node:    {{start: v{16, 0, 0}}},
	},
	ArraySpread: {
		Chrome:  {{start: v{46, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{36, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Arrow: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	AsyncAwait: {
		Chrome:  {{start: v{55, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2017, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{7, 6, 0}}},
		Opera:   {{start: v{42, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	AsyncGenerator: {
		Chrome:  {{start: v{63, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	BigInt: {
		Chrome:  {{start: v{67, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{68, 0, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Opera:   {{start: v{54, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	Class: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ClassField: {
		Chrome:  {{start: v{73, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{69, 0, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	ClassPrivateAccessor: {
		Chrome:  {{start: v{84, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateBrandCheck: {
		Chrome:  {{start: v{91, 0, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{16, 9, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateField: {
		Chrome:  {{start: v{84, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateMethod: {
		Chrome:  {{start: v{84, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticAccessor: {
		Chrome:  {{start: v{84, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticField: {
		Chrome:  {{start: v{74, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateStaticMethod: {
		Chrome:  {{start: v{84, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassStaticBlocks: {
		Chrome: {{start: v{91, 0, 0}}},
		ES:     {{start: v{2022, 0, 0}}},
		Node:   {{start: v{16, 11, 0}}},
	},
	ClassStaticField: {
		Chrome:  {{start: v{73, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{75, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	Const: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{51, 0, 0}}},
		IE:      {{start: v{11, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	DefaultArgument: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Destructuring: {
		Chrome:  {{start: v{51, 0, 0}}},
		Edge:    {{start: v{18, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	DynamicImport: {
		Chrome:  {{start: v{63, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{5, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{13, 2, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	ExponentOperator: {
		Chrome:  {{start: v{52, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{7, 0, 0}}},
		Opera:   {{start: v{39, 0, 0}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	ExportStarAs: {
		Chrome:  {{start: v{72, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{80, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
	},
	ForAwait: {
		Chrome:  {{start: v{63, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	ForOf: {
		Chrome:  {{start: v{51, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Generator: {
		Chrome:  {{start: v{50, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{37, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Hashbang: {
		Chrome:  {{start: v{74, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ImportAssertions: {
		Chrome: {{start: v{91, 0, 0}}},
		Node:   {{start: v{16, 14, 0}}},
	},
	ImportMeta: {
		Chrome:  {{start: v{64, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{62, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	Let: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{51, 0, 0}}},
		IE:      {{start: v{11, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	LogicalAssignment: {
		Chrome:  {{start: v{85, 0, 0}}},
		Edge:    {{start: v{85, 0, 0}}},
		ES:      {{start: v{2021, 0, 0}}},
		Firefox: {{start: v{79, 0, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{15, 0, 0}}},
		Opera:   {{start: v{71, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	NestedRestBinding: {
		Chrome:  {{start: v{49, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{47, 0, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	NewTarget: {
		Chrome:  {{start: v{46, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{41, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	NodeColonPrefixImport: {
		Node: {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{14, 13, 1}}},
	},
	NodeColonPrefixRequire: {
		Node: {{start: v{14, 18, 0}, end: v{15, 0, 0}}, {start: v{16, 0, 0}}},
	},
	NullishCoalescing: {
		Chrome:  {{start: v{80, 0, 0}}},
		Edge:    {{start: v{80, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{72, 0, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{14, 0, 0}}},
		Opera:   {{start: v{67, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ObjectAccessors: {
		Chrome:  {{start: v{5, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{5, 0, 0}}},
		Firefox: {{start: v{2, 0, 0}}},
		IE:      {{start: v{9, 0, 0}}},
		IOS:     {{start: v{6, 0, 0}}},
		Node:    {{start: v{0, 10, 0}}},
		Opera:   {{start: v{10, 10, 0}}},
		Safari:  {{start: v{3, 1, 0}}},
	},
	ObjectExtensions: {
		Chrome:  {{start: v{44, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ObjectRestSpread: {
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{55, 0, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Opera:   {{start: v{47, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalCatchBinding: {
		Chrome:  {{start: v{66, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2019, 0, 0}}},
		Firefox: {{start: v{58, 0, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{53, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalChain: {
		Chrome:  {{start: v{91, 0, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{74, 0, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{16, 9, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	RestArgument: {
		Chrome:  {{start: v{47, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{43, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{34, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	TemplateLiteral: {
		Chrome:  {{start: v{41, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		IOS:     {{start: v{9, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{28, 0, 0}}},
		Safari:  {{start: v{9, 0, 0}}},
	},
	TopLevelAwait: {
		Chrome:  {{start: v{89, 0, 0}}},
		Edge:    {{start: v{89, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{89, 0, 0}}},
		Node:    {{start: v{14, 8, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	TypeofExoticObjectIsObject: {
		Chrome:  {{start: v{0, 0, 0}}},
		Edge:    {{start: v{0, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{0, 0, 0}}},
		IOS:     {{start: v{0, 0, 0}}},
		Node:    {{start: v{0, 0, 0}}},
		Opera:   {{start: v{0, 0, 0}}},
		Safari:  {{start: v{0, 0, 0}}},
	},
	UnicodeEscapes: {
		Chrome:  {{start: v{44, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{9, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{9, 0, 0}}},
	},
}

// Return all features that are not available in at least one environment
func UnsupportedJSFeatures(constraints map[Engine][]int) (unsupported JSFeature) {
	for feature, engines := range jsTable {
		for engine, version := range constraints {
			if versionRanges, ok := engines[engine]; !ok || !isVersionSupported(versionRanges, version) {
				unsupported |= feature
			}
		}
	}
	return
}
