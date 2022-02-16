/**
 * @fileoverview gRPC-Web generated client stub for callback
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.callback = require('./callback_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.callback.CallbackServiceClient =
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
proto.callback.CallbackServicePromiseClient =
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
 *   !proto.callback.CallbackRequest,
 *   !proto.callback.CallbackResponse>}
 */
const methodDescriptor_CallbackService_Callback = new grpc.web.MethodDescriptor(
  '/callback.CallbackService/Callback',
  grpc.web.MethodType.UNARY,
  proto.callback.CallbackRequest,
  proto.callback.CallbackResponse,
  /**
   * @param {!proto.callback.CallbackRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.callback.CallbackResponse.deserializeBinary
);


/**
 * @param {!proto.callback.CallbackRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.callback.CallbackResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.callback.CallbackResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.callback.CallbackServiceClient.prototype.callback =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/callback.CallbackService/Callback',
      request,
      metadata || {},
      methodDescriptor_CallbackService_Callback,
      callback);
};


/**
 * @param {!proto.callback.CallbackRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.callback.CallbackResponse>}
 *     Promise that resolves to the response
 */
proto.callback.CallbackServicePromiseClient.prototype.callback =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/callback.CallbackService/Callback',
      request,
      metadata || {},
      methodDescriptor_CallbackService_Callback);
};


module.exports = proto.callback;

