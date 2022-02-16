/**
 * @fileoverview gRPC-Web generated client stub for wood.doors
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var wood$colors_pb = require('./wood-colors_pb.js')
const proto = {};
proto.wood = {};
proto.wood.doors = require('./wood-doors_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.wood.doors.WoodDoorsApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.wood.doors.WoodDoorsApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.wood.doors.WoodDoorsRequest,
 *   !proto.wood.doors.WoodDoorsResponse>}
 */
const methodDescriptor_WoodDoorsApi_GetWoodDoors = new grpc.web.MethodDescriptor(
  '/wood.doors.WoodDoorsApi/GetWoodDoors',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.wood.doors.WoodDoorsRequest,
  proto.wood.doors.WoodDoorsResponse,
  /**
   * @param {!proto.wood.doors.WoodDoorsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.wood.doors.WoodDoorsResponse.deserializeBinary
);


/**
 * @param {!proto.wood.doors.WoodDoorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.doors.WoodDoorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.doors.WoodDoorsApiClient.prototype.getWoodDoors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.doors.WoodDoorsApi/GetWoodDoors',
      request,
      metadata || {},
      methodDescriptor_WoodDoorsApi_GetWoodDoors);
};


/**
 * @param {!proto.wood.doors.WoodDoorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.doors.WoodDoorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.doors.WoodDoorsApiPromiseClient.prototype.getWoodDoors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.doors.WoodDoorsApi/GetWoodDoors',
      request,
      metadata || {},
      methodDescriptor_WoodDoorsApi_GetWoodDoors);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.wood.doors.CreateWoodDoorRequest,
 *   !proto.wood.doors.CreateWoodDoorResponse>}
 */
const methodDescriptor_WoodDoorsApi_CreateWoodDoor = new grpc.web.MethodDescriptor(
  '/wood.doors.WoodDoorsApi/CreateWoodDoor',
  grpc.web.MethodType.UNARY,
  proto.wood.doors.CreateWoodDoorRequest,
  proto.wood.doors.CreateWoodDoorResponse,
  /**
   * @param {!proto.wood.doors.CreateWoodDoorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.wood.doors.CreateWoodDoorResponse.deserializeBinary
);


/**
 * @param {!proto.wood.doors.CreateWoodDoorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.wood.doors.CreateWoodDoorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.wood.doors.CreateWoodDoorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.wood.doors.WoodDoorsApiClient.prototype.createWoodDoor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/wood.doors.WoodDoorsApi/CreateWoodDoor',
      request,
      metadata || {},
      methodDescriptor_WoodDoorsApi_CreateWoodDoor,
      callback);
};


/**
 * @param {!proto.wood.doors.CreateWoodDoorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.wood.doors.CreateWoodDoorResponse>}
 *     Promise that resolves to the response
 */
proto.wood.doors.WoodDoorsApiPromiseClient.prototype.createWoodDoor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/wood.doors.WoodDoorsApi/CreateWoodDoor',
      request,
      metadata || {},
      methodDescriptor_WoodDoorsApi_CreateWoodDoor);
};


module.exports = proto.wood.doors;

