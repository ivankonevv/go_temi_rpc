/**
 * @fileoverview gRPC-Web generated client stub for user
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.user = require('./user_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.user.UserServiceClient =
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
proto.user.UserServicePromiseClient =
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
 *   !proto.user.GetUserRequest,
 *   !proto.user.GetUserResponse>}
 */
const methodDescriptor_UserService_GetUser = new grpc.web.MethodDescriptor(
  '/user.UserService/GetUser',
  grpc.web.MethodType.UNARY,
  proto.user.GetUserRequest,
  proto.user.GetUserResponse,
  /**
   * @param {!proto.user.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.GetUserResponse.deserializeBinary
);


/**
 * @param {!proto.user.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.GetUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.GetUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.UserServiceClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser,
      callback);
};


/**
 * @param {!proto.user.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.GetUserResponse>}
 *     Promise that resolves to the response
 */
proto.user.UserServicePromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.GetAllUsersRequest,
 *   !proto.user.GetUserResponse>}
 */
const methodDescriptor_UserService_GetAllUsers = new grpc.web.MethodDescriptor(
  '/user.UserService/GetAllUsers',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.user.GetAllUsersRequest,
  proto.user.GetUserResponse,
  /**
   * @param {!proto.user.GetAllUsersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.GetUserResponse.deserializeBinary
);


/**
 * @param {!proto.user.GetAllUsersRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.user.GetUserResponse>}
 *     The XHR Node Readable Stream
 */
proto.user.UserServiceClient.prototype.getAllUsers =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/user.UserService/GetAllUsers',
      request,
      metadata || {},
      methodDescriptor_UserService_GetAllUsers);
};


/**
 * @param {!proto.user.GetAllUsersRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.user.GetUserResponse>}
 *     The XHR Node Readable Stream
 */
proto.user.UserServicePromiseClient.prototype.getAllUsers =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/user.UserService/GetAllUsers',
      request,
      metadata || {},
      methodDescriptor_UserService_GetAllUsers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.NewUserRequest,
 *   !proto.user.NewUserResponse>}
 */
const methodDescriptor_UserService_CreateUser = new grpc.web.MethodDescriptor(
  '/user.UserService/CreateUser',
  grpc.web.MethodType.UNARY,
  proto.user.NewUserRequest,
  proto.user.NewUserResponse,
  /**
   * @param {!proto.user.NewUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.NewUserResponse.deserializeBinary
);


/**
 * @param {!proto.user.NewUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.NewUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.NewUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.UserServiceClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.UserService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_CreateUser,
      callback);
};


/**
 * @param {!proto.user.NewUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.NewUserResponse>}
 *     Promise that resolves to the response
 */
proto.user.UserServicePromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.UserService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_CreateUser);
};


module.exports = proto.user;

