/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "kempy007.atomicchocolate.atomicchocolate";

export interface Review {
  creator: string;
  id: number;
  title: string;
  review: string;
  projectID: number;
  created: string;
}

const baseReview: object = {
  creator: "",
  id: 0,
  title: "",
  review: "",
  projectID: 0,
  created: "",
};

export const Review = {
  encode(message: Review, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.review !== "") {
      writer.uint32(34).string(message.review);
    }
    if (message.projectID !== 0) {
      writer.uint32(40).uint64(message.projectID);
    }
    if (message.created !== "") {
      writer.uint32(50).string(message.created);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Review {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseReview } as Review;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.review = reader.string();
          break;
        case 5:
          message.projectID = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.created = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Review {
    const message = { ...baseReview } as Review;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.review !== undefined && object.review !== null) {
      message.review = String(object.review);
    } else {
      message.review = "";
    }
    if (object.projectID !== undefined && object.projectID !== null) {
      message.projectID = Number(object.projectID);
    } else {
      message.projectID = 0;
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = String(object.created);
    } else {
      message.created = "";
    }
    return message;
  },

  toJSON(message: Review): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.review !== undefined && (obj.review = message.review);
    message.projectID !== undefined && (obj.projectID = message.projectID);
    message.created !== undefined && (obj.created = message.created);
    return obj;
  },

  fromPartial(object: DeepPartial<Review>): Review {
    const message = { ...baseReview } as Review;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.review !== undefined && object.review !== null) {
      message.review = object.review;
    } else {
      message.review = "";
    }
    if (object.projectID !== undefined && object.projectID !== null) {
      message.projectID = object.projectID;
    } else {
      message.projectID = 0;
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = object.created;
    } else {
      message.created = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
