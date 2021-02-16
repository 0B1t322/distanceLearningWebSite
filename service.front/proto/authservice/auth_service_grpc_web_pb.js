/**
 * @fileoverview gRPC-Web generated client stub for authservice
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.authservice = require('./auth_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authservice.AuthServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authservice.AuthServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.authservice.AuthRequest,
 *   !proto.authservice.AuthResponse>}
 */
const methodDescriptor_AuthService_SignIn = new grpc.web.MethodDescriptor(
  '/authservice.AuthService/SignIn',
  grpc.web.MethodType.UNARY,
  proto.authservice.AuthRequest,
  proto.authservice.AuthResponse,
  /**
   * @param {!proto.authservice.AuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.AuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authservice.AuthRequest,
 *   !proto.authservice.AuthResponse>}
 */
const methodInfo_AuthService_SignIn = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authservice.AuthResponse,
  /**
   * @param {!proto.authservice.AuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.authservice.AuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authservice.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authservice.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authservice.AuthServiceClient.prototype.signIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authservice.AuthService/SignIn',
      request,
      metadata || {},
      methodDescriptor_AuthService_SignIn,
      callback);
};


/**
 * @param {!proto.authservice.AuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authservice.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.authservice.AuthServicePromiseClient.prototype.signIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authservice.AuthService/SignIn',
      request,
      metadata || {},
      methodDescriptor_AuthService_SignIn);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.authservice.AuthRequest,
 *   !proto.authservice.AuthResponse>}
 */
const methodDescriptor_AuthService_SignUp = new grpc.web.MethodDescriptor(
  '/authservice.AuthService/SignUp',
  grpc.web.MethodType.UNARY,
  proto.authservice.AuthRequest,
  proto.authservice.AuthResponse,
  /**
   * @param {!proto.authservice.AuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.AuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authservice.AuthRequest,
 *   !proto.authservice.AuthResponse>}
 */
const methodInfo_AuthService_SignUp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authservice.AuthResponse,
  /**
   * @param {!proto.authservice.AuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.authservice.AuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authservice.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authservice.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authservice.AuthServiceClient.prototype.signUp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authservice.AuthService/SignUp',
      request,
      metadata || {},
      methodDescriptor_AuthService_SignUp,
      callback);
};


/**
 * @param {!proto.authservice.AuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authservice.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.authservice.AuthServicePromiseClient.prototype.signUp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authservice.AuthService/SignUp',
      request,
      metadata || {},
      methodDescriptor_AuthService_SignUp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.authservice.Token,
 *   !proto.authservice.TokenInfo>}
 */
const methodDescriptor_AuthService_Check = new grpc.web.MethodDescriptor(
  '/authservice.AuthService/Check',
  grpc.web.MethodType.UNARY,
  proto.authservice.Token,
  proto.authservice.TokenInfo,
  /**
   * @param {!proto.authservice.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.TokenInfo.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authservice.Token,
 *   !proto.authservice.TokenInfo>}
 */
const methodInfo_AuthService_Check = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authservice.TokenInfo,
  /**
   * @param {!proto.authservice.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.TokenInfo.deserializeBinary
);


/**
 * @param {!proto.authservice.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authservice.TokenInfo)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authservice.TokenInfo>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authservice.AuthServiceClient.prototype.check =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authservice.AuthService/Check',
      request,
      metadata || {},
      methodDescriptor_AuthService_Check,
      callback);
};


/**
 * @param {!proto.authservice.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authservice.TokenInfo>}
 *     Promise that resolves to the response
 */
proto.authservice.AuthServicePromiseClient.prototype.check =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authservice.AuthService/Check',
      request,
      metadata || {},
      methodDescriptor_AuthService_Check);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.authservice.DeleteUserReq,
 *   !proto.authservice.DeleteUserResp>}
 */
const methodDescriptor_AuthService_DeleteUser = new grpc.web.MethodDescriptor(
  '/authservice.AuthService/DeleteUser',
  grpc.web.MethodType.UNARY,
  proto.authservice.DeleteUserReq,
  proto.authservice.DeleteUserResp,
  /**
   * @param {!proto.authservice.DeleteUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.DeleteUserResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authservice.DeleteUserReq,
 *   !proto.authservice.DeleteUserResp>}
 */
const methodInfo_AuthService_DeleteUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authservice.DeleteUserResp,
  /**
   * @param {!proto.authservice.DeleteUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authservice.DeleteUserResp.deserializeBinary
);


/**
 * @param {!proto.authservice.DeleteUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authservice.DeleteUserResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authservice.DeleteUserResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authservice.AuthServiceClient.prototype.deleteUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authservice.AuthService/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_DeleteUser,
      callback);
};


/**
 * @param {!proto.authservice.DeleteUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authservice.DeleteUserResp>}
 *     Promise that resolves to the response
 */
proto.authservice.AuthServicePromiseClient.prototype.deleteUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authservice.AuthService/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_DeleteUser);
};


module.exports = proto.authservice;

