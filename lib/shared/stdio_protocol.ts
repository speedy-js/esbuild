// The JavaScript API communicates with the Go child process over stdin/stdout
// using this protocol. It's a very simple binary protocol that uses primitives
// and nested arrays and maps. It's basically JSON with UTF-8 encoding and an
// additional byte array primitive. You must send a response after receiving a
// request because the other end is blocking on the response coming back.

import * as types from "./types";

export interface BuildRequest {
  command: "build";
  key: number;
  entries: [string, string][]; // Use an array instead of a map to preserve order
  flags: string[];
  write: boolean;
  stdinContents: string | null;
  stdinResolveDir: string | null;
  absWorkingDir: string;
  incremental: boolean;
  nodePaths: string[];
  plugins?: BuildPlugin[];
  serve?: ServeRequest;
  mangleCache?: Record<string, string | false>;
}

export interface ServeRequest {
  port?: number;
  host?: string;
  servedir?: string;
}

export interface ServeResponse {
  port: number;
  host: string;
}

export interface ServeStopRequest {
  command: "serve-stop";
  key: number;
}

export interface BuildPlugin {
  name: string;
  onResolve: { id: number; filter: string; namespace: string }[];
  onLoad: { id: number; filter: string; namespace: string }[];
}

export interface BuildResponse {
  errors: types.Message[];
  warnings: types.Message[];
  rebuild: boolean;
  watch: boolean;
  outputFiles?: BuildOutputFile[];
  metafile?: string;
  mangleCache?: Record<string, string | false>;
  writeToStdout?: Uint8Array;
}

export interface BuildOutputFile {
  path: string;
  contents: Uint8Array;
}

export interface PingRequest {
  command: "ping";
}

export interface RebuildRequest {
  command: "rebuild";
  key: number;
  changefile: string[];
}

export interface RebuildDisposeRequest {
  command: "rebuild-dispose";
  key: number;
}

export interface WatchStopRequest {
  command: "watch-stop";
  key: number;
}

export interface OnRequestRequest {
  command: "serve-request";
  key: number;
  args: types.ServeOnRequestArgs;
}

export interface OnWaitRequest {
  command: "serve-wait";
  key: number;
  error: string | null;
}

export interface OnWatchRebuildRequest {
  command: "watch-rebuild";
  key: number;
  args: types.BuildResult;
}

export interface TransformRequest {
  command: "transform";
  flags: string[];
  input: string;
  inputFS: boolean;
  mangleCache?: Record<string, string | false>;
}

export interface TransformResponse {
  errors: types.Message[];
  warnings: types.Message[];

  code: string;
  codeFS: boolean;

  map: string;
  mapFS: boolean;

  mangleCache?: Record<string, string | false>;
}

export interface FormatMsgsRequest {
  command: "format-msgs";
  messages: types.Message[];
  isWarning: boolean;
  color?: boolean;
  terminalWidth?: number;
}

export interface FormatMsgsResponse {
  messages: string[];
}

export interface AnalyzeMetafileRequest {
  command: "analyze-metafile";
  metafile: string;
  color?: boolean;
  verbose?: boolean;
}

export interface AnalyzeMetafileResponse {
  result: string;
}

export interface OnStartRequest {
  command: "on-start";
  key: number;
}

export interface OnStartResponse {
  errors?: types.PartialMessage[];
  warnings?: types.PartialMessage[];
}

export interface ResolveRequest {
  command: "resolve";
  key: number;
  path: string;
  pluginName: string;
  importer?: string;
  namespace?: string;
  resolveDir?: string;
  kind?: string;
  pluginData?: number;
  origin?: boolean;
}

export interface ResolveResponse {
  errors: types.Message[];
  warnings: types.Message[];

  path: string;
  external: boolean;
  sideEffects: boolean;
  namespace: string;
  suffix: string;
  pluginData: number;
}

export interface OnResolveRequest {
  command: "on-resolve";
  key: number;
  ids: number[];
  path: string;
  importer: string;
  namespace: string;
  resolveDir: string;
  kind: types.ImportKind;
  pluginData: number;
}

export interface OnResolveResponse {
  id?: number;
  pluginName?: string;

  errors?: types.PartialMessage[];
  warnings?: types.PartialMessage[];

  path?: string;
  external?: boolean;
  sideEffects?: boolean;
  namespace?: string;
  suffix?: string;
  pluginData?: number;

  cacheDisable?: boolean;

  watchFiles?: string[];
  watchDirs?: string[];
}

export interface OnLoadRequest {
  command: "on-load";
  key: number;
  ids: number[];
  path: string;
  namespace: string;
  suffix: string;
  pluginData: number;
}

export interface OnLoadResponse {
  id?: number;
  pluginName?: string;

  errors?: types.PartialMessage[];
  warnings?: types.PartialMessage[];

  contents?: Uint8Array;
  resolveDir?: string;
  loader?: string;
  pluginData?: number;

  cacheDisable?: boolean;

  watchFiles?: string[];
  watchDirs?: string[];
}

////////////////////////////////////////////////////////////////////////////////

export interface Packet {
  id: number;
  isRequest: boolean;
  value: Value;
}

export type Value =
  | null
  | boolean
  | number
  | string
  | Uint8Array
  | Value[]
  | { [key: string]: Value };

export function encodePacket(packet: Packet): Uint8Array {
  let visit = (value: Value) => {
    if (value === null) {
      bb.write8(0);
    } else if (typeof value === "boolean") {
      bb.write8(1);
      bb.write8(+value);
    } else if (typeof value === "number") {
      bb.write8(2);
      bb.write32(value | 0);
    } else if (typeof value === "string") {
      bb.write8(3);
      bb.write(encodeUTF8(value));
    } else if (value instanceof Uint8Array) {
      bb.write8(4);
      bb.write(value);
    } else if (value instanceof Array) {
      bb.write8(5);
      bb.write32(value.length);
      for (let item of value) {
        visit(item);
      }
    } else {
      let keys = Object.keys(value);
      bb.write8(6);
      bb.write32(keys.length);
      for (let key of keys) {
        bb.write(encodeUTF8(key));
        visit(value[key]);
      }
    }
  };

  let bb = new ByteBuffer();
  bb.write32(0); // Reserve space for the length
  bb.write32((packet.id << 1) | +!packet.isRequest);
  visit(packet.value);
  writeUInt32LE(bb.buf, bb.len - 4, 0); // Patch the length in
  return bb.buf.subarray(0, bb.len);
}

export function decodePacket(bytes: Uint8Array): Packet {
  let visit = (): Value => {
    switch (bb.read8()) {
      case 0: // null
        return null;
      case 1: // boolean
        return !!bb.read8();
      case 2: // number
        return bb.read32();
      case 3: // string
        return decodeUTF8(bb.read());
      case 4: // Uint8Array
        return bb.read();
      case 5: {
        // Value[]
        let count = bb.read32();
        let value: Value[] = [];
        for (let i = 0; i < count; i++) {
          value.push(visit());
        }
        return value;
      }
      case 6: {
        // { [key: string]: Value }
        let count = bb.read32();
        let value: { [key: string]: Value } = {};
        for (let i = 0; i < count; i++) {
          value[decodeUTF8(bb.read())] = visit();
        }
        return value;
      }
      default:
        throw new Error("Invalid packet");
    }
  };

  let bb = new ByteBuffer(bytes);
  let id = bb.read32();
  let isRequest = (id & 1) === 0;
  id >>>= 1;
  let value = visit();
  if (bb.ptr !== bytes.length) {
    throw new Error("Invalid packet");
  }
  return {id, isRequest, value};
}

class ByteBuffer {
  len = 0;
  ptr = 0;

  constructor(public buf = new Uint8Array(1024)) {
  }

  private _write(delta: number): number {
    if (this.len + delta > this.buf.length) {
      let clone = new Uint8Array((this.len + delta) * 2);
      clone.set(this.buf);
      this.buf = clone;
    }
    this.len += delta;
    return this.len - delta;
  }

  write8(value: number): void {
    let offset = this._write(1);
    this.buf[offset] = value;
  }

  write32(value: number): void {
    let offset = this._write(4);
    writeUInt32LE(this.buf, value, offset);
  }

  write(bytes: Uint8Array): void {
    let offset = this._write(4 + bytes.length);
    writeUInt32LE(this.buf, bytes.length, offset);
    this.buf.set(bytes, offset + 4);
  }

  private _read(delta: number): number {
    if (this.ptr + delta > this.buf.length) {
      throw new Error("Invalid packet");
    }
    this.ptr += delta;
    return this.ptr - delta;
  }

  read8(): number {
    return this.buf[this._read(1)];
  }

  read32(): number {
    return readUInt32LE(this.buf, this._read(4));
  }

  read(): Uint8Array {
    let length = this.read32();
    let bytes = new Uint8Array(length);
    let ptr = this._read(bytes.length);
    bytes.set(this.buf.subarray(ptr, ptr + length));
    return bytes;
  }
}

export let encodeUTF8: (text: string) => Uint8Array;
export let decodeUTF8: (bytes: Uint8Array) => string;

// For the browser and node 12.x
if (typeof TextEncoder !== "undefined" && typeof TextDecoder !== "undefined") {
  let encoder = new TextEncoder();
  let decoder = new TextDecoder();
  encodeUTF8 = (text) => encoder.encode(text);
  decodeUTF8 = (bytes) => decoder.decode(bytes);
}

// For node 10.x
else if (typeof Buffer !== "undefined") {
  encodeUTF8 = (text) => {
    let buffer: Uint8Array = Buffer.from(text);

    // The test framework called "Jest" breaks node's Buffer API. Normally
    // instances of Buffer are also instances of Uint8Array, but not when
    // esbuild is run inside of whatever weird environment Jest uses. More
    // info: https://github.com/facebook/jest/issues/4422.
    if (!(buffer instanceof Uint8Array)) {
      // Construct a new Uint8Array with the contents of the buffer to force
      // it to be a Uint8Array instance. This is wasteful since it's slower
      // than just using the Buffer, but this should only happen when esbuild
      // is run inside of Jest.
      buffer = new Uint8Array(buffer);
    }

    return buffer;
  };
  decodeUTF8 = (bytes) => {
    let {buffer, byteOffset, byteLength} = bytes;
    return Buffer.from(buffer, byteOffset, byteLength).toString();
  };
} else {
  throw new Error("No UTF-8 codec found");
}

export function readUInt32LE(buffer: Uint8Array, offset: number): number {
  return (
    buffer[offset++] |
    (buffer[offset++] << 8) |
    (buffer[offset++] << 16) |
    (buffer[offset++] << 24)
  );
}

function writeUInt32LE(
  buffer: Uint8Array,
  value: number,
  offset: number
): void {
  buffer[offset++] = value;
  buffer[offset++] = value >> 8;
  buffer[offset++] = value >> 16;
  buffer[offset++] = value >> 24;
}
