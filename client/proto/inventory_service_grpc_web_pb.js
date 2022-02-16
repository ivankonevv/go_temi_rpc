/**
 * @fileoverview gRPC-Web generated client stub for inventory
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.inventory = require('./inventory_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.inventory.InventoryServiceClient =
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
proto.inventory.InventoryServicePromiseClient =
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
 *   !proto.inventory.NewInventoryRequest,
 *   !proto.inventory.NewInventoryResponse>}
 */
const methodDescriptor_InventoryService_NewInventory = new grpc.web.MethodDescriptor(
  '/inventory.InventoryService/NewInventory',
  grpc.web.MethodType.UNARY,
  proto.inventory.NewInventoryRequest,
  proto.inventory.NewInventoryResponse,
  /**
   * @param {!proto.inventory.NewInventoryRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.inventory.NewInventoryResponse.deserializeBinary
);


/**
 * @param {!proto.inventory.NewInventoryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.inventory.NewInventoryResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.inventory.NewInventoryResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.inventory.InventoryServiceClient.prototype.newInventory =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/inventory.InventoryService/NewInventory',
      request,
      metadata || {},
      methodDescriptor_InventoryService_NewInventory,
      callback);
};


/**
 * @param {!proto.inventory.NewInventoryRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.inventory.NewInventoryResponse>}
 *     Promise that resolves to the response
 */
proto.inventory.InventoryServicePromiseClient.prototype.newInventory =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/inventory.InventoryService/NewInventory',
      request,
      metadata || {},
      methodDescriptor_InventoryService_NewInventory);
};


module.exports = proto.inventory;

