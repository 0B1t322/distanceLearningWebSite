/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./coursesservice_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.CoursesServiceClient =
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
proto.CoursesServicePromiseClient =
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
 *   !proto.AddCourseReq,
 *   !proto.AddCourseResp>}
 */
const methodDescriptor_CoursesService_AddCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/AddCourse',
  grpc.web.MethodType.UNARY,
  proto.AddCourseReq,
  proto.AddCourseResp,
  /**
   * @param {!proto.AddCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.AddCourseReq,
 *   !proto.AddCourseResp>}
 */
const methodInfo_CoursesService_AddCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.AddCourseResp,
  /**
   * @param {!proto.AddCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddCourseResp.deserializeBinary
);


/**
 * @param {!proto.AddCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.AddCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.AddCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.addCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/AddCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddCourse,
      callback);
};


/**
 * @param {!proto.AddCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.AddCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.addCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/AddCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddCourse);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UpdateCourseReq,
 *   !proto.UpdateCourseResp>}
 */
const methodDescriptor_CoursesService_UpdateCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/UpdateCourse',
  grpc.web.MethodType.UNARY,
  proto.UpdateCourseReq,
  proto.UpdateCourseResp,
  /**
   * @param {!proto.UpdateCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.UpdateCourseReq,
 *   !proto.UpdateCourseResp>}
 */
const methodInfo_CoursesService_UpdateCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.UpdateCourseResp,
  /**
   * @param {!proto.UpdateCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateCourseResp.deserializeBinary
);


/**
 * @param {!proto.UpdateCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.UpdateCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.UpdateCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.updateCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/UpdateCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateCourse,
      callback);
};


/**
 * @param {!proto.UpdateCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.UpdateCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.updateCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/UpdateCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateCourse);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DeleteCourseReq,
 *   !proto.DeleteCourseResp>}
 */
const methodDescriptor_CoursesService_DeleteCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/DeleteCourse',
  grpc.web.MethodType.UNARY,
  proto.DeleteCourseReq,
  proto.DeleteCourseResp,
  /**
   * @param {!proto.DeleteCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DeleteCourseReq,
 *   !proto.DeleteCourseResp>}
 */
const methodInfo_CoursesService_DeleteCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DeleteCourseResp,
  /**
   * @param {!proto.DeleteCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteCourseResp.deserializeBinary
);


/**
 * @param {!proto.DeleteCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DeleteCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DeleteCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.deleteCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/DeleteCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteCourse,
      callback);
};


/**
 * @param {!proto.DeleteCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DeleteCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.deleteCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/DeleteCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteCourse);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.GetCourseReq,
 *   !proto.GetCourseResp>}
 */
const methodDescriptor_CoursesService_GetCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/GetCourse',
  grpc.web.MethodType.UNARY,
  proto.GetCourseReq,
  proto.GetCourseResp,
  /**
   * @param {!proto.GetCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GetCourseReq,
 *   !proto.GetCourseResp>}
 */
const methodInfo_CoursesService_GetCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GetCourseResp,
  /**
   * @param {!proto.GetCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetCourseResp.deserializeBinary
);


/**
 * @param {!proto.GetCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GetCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GetCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.getCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/GetCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_GetCourse,
      callback);
};


/**
 * @param {!proto.GetCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GetCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.getCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/GetCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_GetCourse);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.GetAllCoursesReq,
 *   !proto.GetAllCoursesResp>}
 */
const methodDescriptor_CoursesService_GetAllCourses = new grpc.web.MethodDescriptor(
  '/CoursesService/GetAllCourses',
  grpc.web.MethodType.UNARY,
  proto.GetAllCoursesReq,
  proto.GetAllCoursesResp,
  /**
   * @param {!proto.GetAllCoursesReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetAllCoursesResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GetAllCoursesReq,
 *   !proto.GetAllCoursesResp>}
 */
const methodInfo_CoursesService_GetAllCourses = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GetAllCoursesResp,
  /**
   * @param {!proto.GetAllCoursesReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetAllCoursesResp.deserializeBinary
);


/**
 * @param {!proto.GetAllCoursesReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GetAllCoursesResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GetAllCoursesResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.getAllCourses =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/GetAllCourses',
      request,
      metadata || {},
      methodDescriptor_CoursesService_GetAllCourses,
      callback);
};


/**
 * @param {!proto.GetAllCoursesReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GetAllCoursesResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.getAllCourses =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/GetAllCourses',
      request,
      metadata || {},
      methodDescriptor_CoursesService_GetAllCourses);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.AddTaskHeaderReq,
 *   !proto.AddTaskHeaderResp>}
 */
const methodDescriptor_CoursesService_AddTaskHeader = new grpc.web.MethodDescriptor(
  '/CoursesService/AddTaskHeader',
  grpc.web.MethodType.UNARY,
  proto.AddTaskHeaderReq,
  proto.AddTaskHeaderResp,
  /**
   * @param {!proto.AddTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddTaskHeaderResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.AddTaskHeaderReq,
 *   !proto.AddTaskHeaderResp>}
 */
const methodInfo_CoursesService_AddTaskHeader = new grpc.web.AbstractClientBase.MethodInfo(
  proto.AddTaskHeaderResp,
  /**
   * @param {!proto.AddTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddTaskHeaderResp.deserializeBinary
);


/**
 * @param {!proto.AddTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.AddTaskHeaderResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.AddTaskHeaderResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.addTaskHeader =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/AddTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddTaskHeader,
      callback);
};


/**
 * @param {!proto.AddTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.AddTaskHeaderResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.addTaskHeader =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/AddTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddTaskHeader);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UpdateTaskHeaderReq,
 *   !proto.UpdateTaskHeaderResp>}
 */
const methodDescriptor_CoursesService_UpdateTaskHeader = new grpc.web.MethodDescriptor(
  '/CoursesService/UpdateTaskHeader',
  grpc.web.MethodType.UNARY,
  proto.UpdateTaskHeaderReq,
  proto.UpdateTaskHeaderResp,
  /**
   * @param {!proto.UpdateTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateTaskHeaderResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.UpdateTaskHeaderReq,
 *   !proto.UpdateTaskHeaderResp>}
 */
const methodInfo_CoursesService_UpdateTaskHeader = new grpc.web.AbstractClientBase.MethodInfo(
  proto.UpdateTaskHeaderResp,
  /**
   * @param {!proto.UpdateTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateTaskHeaderResp.deserializeBinary
);


/**
 * @param {!proto.UpdateTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.UpdateTaskHeaderResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.UpdateTaskHeaderResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.updateTaskHeader =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/UpdateTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateTaskHeader,
      callback);
};


/**
 * @param {!proto.UpdateTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.UpdateTaskHeaderResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.updateTaskHeader =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/UpdateTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateTaskHeader);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DeleteTaskHeaderReq,
 *   !proto.DeleteTaskHeaderResp>}
 */
const methodDescriptor_CoursesService_DeleteTaskHeader = new grpc.web.MethodDescriptor(
  '/CoursesService/DeleteTaskHeader',
  grpc.web.MethodType.UNARY,
  proto.DeleteTaskHeaderReq,
  proto.DeleteTaskHeaderResp,
  /**
   * @param {!proto.DeleteTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteTaskHeaderResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DeleteTaskHeaderReq,
 *   !proto.DeleteTaskHeaderResp>}
 */
const methodInfo_CoursesService_DeleteTaskHeader = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DeleteTaskHeaderResp,
  /**
   * @param {!proto.DeleteTaskHeaderReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteTaskHeaderResp.deserializeBinary
);


/**
 * @param {!proto.DeleteTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DeleteTaskHeaderResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DeleteTaskHeaderResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.deleteTaskHeader =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/DeleteTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteTaskHeader,
      callback);
};


/**
 * @param {!proto.DeleteTaskHeaderReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DeleteTaskHeaderResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.deleteTaskHeader =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/DeleteTaskHeader',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteTaskHeader);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.AddTaskReq,
 *   !proto.AddTaskResp>}
 */
const methodDescriptor_CoursesService_AddTask = new grpc.web.MethodDescriptor(
  '/CoursesService/AddTask',
  grpc.web.MethodType.UNARY,
  proto.AddTaskReq,
  proto.AddTaskResp,
  /**
   * @param {!proto.AddTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddTaskResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.AddTaskReq,
 *   !proto.AddTaskResp>}
 */
const methodInfo_CoursesService_AddTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.AddTaskResp,
  /**
   * @param {!proto.AddTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddTaskResp.deserializeBinary
);


/**
 * @param {!proto.AddTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.AddTaskResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.AddTaskResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.addTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/AddTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddTask,
      callback);
};


/**
 * @param {!proto.AddTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.AddTaskResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.addTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/AddTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UpdateTaskReq,
 *   !proto.UpdateTaskResp>}
 */
const methodDescriptor_CoursesService_UpdateTask = new grpc.web.MethodDescriptor(
  '/CoursesService/UpdateTask',
  grpc.web.MethodType.UNARY,
  proto.UpdateTaskReq,
  proto.UpdateTaskResp,
  /**
   * @param {!proto.UpdateTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateTaskResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.UpdateTaskReq,
 *   !proto.UpdateTaskResp>}
 */
const methodInfo_CoursesService_UpdateTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.UpdateTaskResp,
  /**
   * @param {!proto.UpdateTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateTaskResp.deserializeBinary
);


/**
 * @param {!proto.UpdateTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.UpdateTaskResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.UpdateTaskResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.updateTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/UpdateTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateTask,
      callback);
};


/**
 * @param {!proto.UpdateTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.UpdateTaskResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.updateTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/UpdateTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_UpdateTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DeleteTaskReq,
 *   !proto.DeleteTaskResp>}
 */
const methodDescriptor_CoursesService_DeleteTask = new grpc.web.MethodDescriptor(
  '/CoursesService/DeleteTask',
  grpc.web.MethodType.UNARY,
  proto.DeleteTaskReq,
  proto.DeleteTaskResp,
  /**
   * @param {!proto.DeleteTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteTaskResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DeleteTaskReq,
 *   !proto.DeleteTaskResp>}
 */
const methodInfo_CoursesService_DeleteTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DeleteTaskResp,
  /**
   * @param {!proto.DeleteTaskReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteTaskResp.deserializeBinary
);


/**
 * @param {!proto.DeleteTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DeleteTaskResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DeleteTaskResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.deleteTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteTask,
      callback);
};


/**
 * @param {!proto.DeleteTaskReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DeleteTaskResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.deleteTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.AddUserInCourseReq,
 *   !proto.AddUserInCourseResp>}
 */
const methodDescriptor_CoursesService_AddUserInCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/AddUserInCourse',
  grpc.web.MethodType.UNARY,
  proto.AddUserInCourseReq,
  proto.AddUserInCourseResp,
  /**
   * @param {!proto.AddUserInCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddUserInCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.AddUserInCourseReq,
 *   !proto.AddUserInCourseResp>}
 */
const methodInfo_CoursesService_AddUserInCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.AddUserInCourseResp,
  /**
   * @param {!proto.AddUserInCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddUserInCourseResp.deserializeBinary
);


/**
 * @param {!proto.AddUserInCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.AddUserInCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.AddUserInCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.addUserInCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/AddUserInCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddUserInCourse,
      callback);
};


/**
 * @param {!proto.AddUserInCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.AddUserInCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.addUserInCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/AddUserInCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_AddUserInCourse);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DeleteUserInCourseReq,
 *   !proto.DeleteUserInCourseResp>}
 */
const methodDescriptor_CoursesService_DeleteUserInCourse = new grpc.web.MethodDescriptor(
  '/CoursesService/DeleteUserInCourse',
  grpc.web.MethodType.UNARY,
  proto.DeleteUserInCourseReq,
  proto.DeleteUserInCourseResp,
  /**
   * @param {!proto.DeleteUserInCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteUserInCourseResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DeleteUserInCourseReq,
 *   !proto.DeleteUserInCourseResp>}
 */
const methodInfo_CoursesService_DeleteUserInCourse = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DeleteUserInCourseResp,
  /**
   * @param {!proto.DeleteUserInCourseReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteUserInCourseResp.deserializeBinary
);


/**
 * @param {!proto.DeleteUserInCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DeleteUserInCourseResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DeleteUserInCourseResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.CoursesServiceClient.prototype.deleteUserInCourse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/CoursesService/DeleteUserInCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteUserInCourse,
      callback);
};


/**
 * @param {!proto.DeleteUserInCourseReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DeleteUserInCourseResp>}
 *     Promise that resolves to the response
 */
proto.CoursesServicePromiseClient.prototype.deleteUserInCourse =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/CoursesService/DeleteUserInCourse',
      request,
      metadata || {},
      methodDescriptor_CoursesService_DeleteUserInCourse);
};


module.exports = proto;

