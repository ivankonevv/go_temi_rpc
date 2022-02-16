/**
 * @fileoverview gRPC-Web generated client stub for wood.color
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.wood = {};
proto.wood.color = require('./wood-colors_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.wood.color.ColorsApiClient =
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
proto.wood.color.ColorsApiPromiseClient =
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
 *   !proto.wood.color.ColorsRequest,
 *   !proto.wood.color.ColorsResponse>}
 */
const methodDescriptor_ColorsApi_GetColors = new grpc.web.MethodDescriptor(
  '/wood.color.ColorsApi/GetColors',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.wood.color.ColorsRequest,
  proto.wood.color.ColorsResponse,
  /**
   * @param {!proto.wood.color.ColorsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.wood.color.ColorsResponse.deserializeBinary
);


/**
 * @param {!proto.wood.color.ColorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.color.ColorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.color.ColorsApiClient.prototype.getColors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.color.ColorsApi/GetColors',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_GetColors);
};


/**
 * @param {!proto.wood.color.ColorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.color.ColorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.color.ColorsApiPromiseClient.prototype.getColors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.color.ColorsApi/GetColors',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_GetColors);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.wood.color.NewColorRequest,
 *   !proto.wood.color.NewColorResponse>}
 */
const methodDescriptor_ColorsApi_CreateColor = new grpc.web.MethodDescriptor(
  '/wood.color.ColorsApi/CreateColor',
  grpc.web.MethodType.UNARY,
  proto.wood.color.NewColorRequest,
  proto.wood.color.NewColorResponse,
  /**
   * @param {!proto.wood.color.NewColorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.wood.color.NewColorResponse.deserializeBinary
);


/**
 * @param {!proto.wood.color.NewColorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.wood.color.NewColorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.wood.color.NewColorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.wood.color.ColorsApiClient.prototype.createColor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/wood.color.ColorsApi/CreateColor',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_CreateColor,
      callback);
};


/**
 * @param {!proto.wood.color.NewColorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.wood.color.NewColorResponse>}
 *     Promise that resolves to the response
 */
proto.wood.color.ColorsApiPromiseClient.prototype.createColor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/wood.color.ColorsApi/CreateColor',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_CreateColor);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.wood.color.ManufacturersWColorsRequest,
 *   !proto.wood.color.ManufacturersWColorsResponse>}
 */
const methodDescriptor_ColorsApi_GetMwColors = new grpc.web.MethodDescriptor(
  '/wood.color.ColorsApi/GetMwColors',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.wood.color.ManufacturersWColorsRequest,
  proto.wood.color.ManufacturersWColorsResponse,
  /**
   * @param {!proto.wood.color.ManufacturersWColorsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.wood.color.ManufacturersWColorsResponse.deserializeBinary
);


/**
 * @param {!proto.wood.color.ManufacturersWColorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.color.ManufacturersWColorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.color.ColorsApiClient.prototype.getMwColors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.color.ColorsApi/GetMwColors',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_GetMwColors);
};


/**
 * @param {!proto.wood.color.ManufacturersWColorsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.wood.color.ManufacturersWColorsResponse>}
 *     The XHR Node Readable Stream
 */
proto.wood.color.ColorsApiPromiseClient.prototype.getMwColors =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/wood.color.ColorsApi/GetMwColors',
      request,
      metadata || {},
      methodDescriptor_ColorsApi_GetMwColors);
};


module.exports = proto.wood.color;

