/**
 * @fileoverview gRPC-Web generated client stub for metal
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.metal = require('./metal-doors_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.metal.MetalDoorsApiClient =
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
proto.metal.MetalDoorsApiPromiseClient =
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
 *   !proto.metal.PostsRequest,
 *   !proto.metal.PostsResponse>}
 */
const methodDescriptor_MetalDoorsApi_GetPosts = new grpc.web.MethodDescriptor(
  '/metal.MetalDoorsApi/GetPosts',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.metal.PostsRequest,
  proto.metal.PostsResponse,
  /**
   * @param {!proto.metal.PostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metal.PostsResponse.deserializeBinary
);


/**
 * @param {!proto.metal.PostsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.metal.PostsResponse>}
 *     The XHR Node Readable Stream
 */
proto.metal.MetalDoorsApiClient.prototype.getPosts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/metal.MetalDoorsApi/GetPosts',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_GetPosts);
};


/**
 * @param {!proto.metal.PostsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.metal.PostsResponse>}
 *     The XHR Node Readable Stream
 */
proto.metal.MetalDoorsApiPromiseClient.prototype.getPosts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/metal.MetalDoorsApi/GetPosts',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_GetPosts);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.metal.CreatePostRequest,
 *   !proto.metal.CreatePostResponse>}
 */
const methodDescriptor_MetalDoorsApi_CreatePost = new grpc.web.MethodDescriptor(
  '/metal.MetalDoorsApi/CreatePost',
  grpc.web.MethodType.UNARY,
  proto.metal.CreatePostRequest,
  proto.metal.CreatePostResponse,
  /**
   * @param {!proto.metal.CreatePostRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metal.CreatePostResponse.deserializeBinary
);


/**
 * @param {!proto.metal.CreatePostRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.metal.CreatePostResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.metal.CreatePostResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.metal.MetalDoorsApiClient.prototype.createPost =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/metal.MetalDoorsApi/CreatePost',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_CreatePost,
      callback);
};


/**
 * @param {!proto.metal.CreatePostRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.metal.CreatePostResponse>}
 *     Promise that resolves to the response
 */
proto.metal.MetalDoorsApiPromiseClient.prototype.createPost =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/metal.MetalDoorsApi/CreatePost',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_CreatePost);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.metal.SingleMetalDoorRequest,
 *   !proto.metal.SingleMetalDoorResponse>}
 */
const methodDescriptor_MetalDoorsApi_GetSingleMetalDoor = new grpc.web.MethodDescriptor(
  '/metal.MetalDoorsApi/GetSingleMetalDoor',
  grpc.web.MethodType.UNARY,
  proto.metal.SingleMetalDoorRequest,
  proto.metal.SingleMetalDoorResponse,
  /**
   * @param {!proto.metal.SingleMetalDoorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metal.SingleMetalDoorResponse.deserializeBinary
);


/**
 * @param {!proto.metal.SingleMetalDoorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.metal.SingleMetalDoorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.metal.SingleMetalDoorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.metal.MetalDoorsApiClient.prototype.getSingleMetalDoor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/metal.MetalDoorsApi/GetSingleMetalDoor',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_GetSingleMetalDoor,
      callback);
};


/**
 * @param {!proto.metal.SingleMetalDoorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.metal.SingleMetalDoorResponse>}
 *     Promise that resolves to the response
 */
proto.metal.MetalDoorsApiPromiseClient.prototype.getSingleMetalDoor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/metal.MetalDoorsApi/GetSingleMetalDoor',
      request,
      metadata || {},
      methodDescriptor_MetalDoorsApi_GetSingleMetalDoor);
};


module.exports = proto.metal;

