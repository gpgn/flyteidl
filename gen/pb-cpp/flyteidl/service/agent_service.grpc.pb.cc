// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: flyteidl/service/agent_service.proto

#include "flyteidl/service/agent_service.pb.h"
#include "flyteidl/service/agent_service.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace flyteidl {
namespace service {

static const char* AgentService_method_names[] = {
  "/flyteidl.service.AgentService/CreateTask",
  "/flyteidl.service.AgentService/GetTask",
  "/flyteidl.service.AgentService/DeleteTask",
};

std::unique_ptr< AgentService::Stub> AgentService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< AgentService::Stub> stub(new AgentService::Stub(channel));
  return stub;
}

AgentService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_CreateTask_(AgentService_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetTask_(AgentService_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_DeleteTask_(AgentService_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status AgentService::Stub::CreateTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskCreateRequest& request, ::flyteidl::service::TaskCreateResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CreateTask_, context, request, response);
}

void AgentService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskCreateRequest* request, ::flyteidl::service::TaskCreateResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskCreateResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskCreateRequest* request, ::flyteidl::service::TaskCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, reactor);
}

void AgentService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskCreateResponse>* AgentService::Stub::AsyncCreateTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskCreateRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskCreateResponse>::Create(channel_.get(), cq, rpcmethod_CreateTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskCreateResponse>* AgentService::Stub::PrepareAsyncCreateTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskCreateRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskCreateResponse>::Create(channel_.get(), cq, rpcmethod_CreateTask_, context, request, false);
}

::grpc::Status AgentService::Stub::GetTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskGetRequest& request, ::flyteidl::service::TaskGetResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetTask_, context, request, response);
}

void AgentService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskGetRequest* request, ::flyteidl::service::TaskGetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskGetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskGetRequest* request, ::flyteidl::service::TaskGetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, reactor);
}

void AgentService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskGetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskGetResponse>* AgentService::Stub::AsyncGetTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskGetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskGetResponse>::Create(channel_.get(), cq, rpcmethod_GetTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskGetResponse>* AgentService::Stub::PrepareAsyncGetTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskGetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskGetResponse>::Create(channel_.get(), cq, rpcmethod_GetTask_, context, request, false);
}

::grpc::Status AgentService::Stub::DeleteTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskDeleteRequest& request, ::flyteidl::service::TaskDeleteResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_DeleteTask_, context, request, response);
}

void AgentService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskDeleteRequest* request, ::flyteidl::service::TaskDeleteResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskDeleteResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, std::move(f));
}

void AgentService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::flyteidl::service::TaskDeleteRequest* request, ::flyteidl::service::TaskDeleteResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, reactor);
}

void AgentService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::TaskDeleteResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskDeleteResponse>* AgentService::Stub::AsyncDeleteTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskDeleteRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskDeleteResponse>::Create(channel_.get(), cq, rpcmethod_DeleteTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::TaskDeleteResponse>* AgentService::Stub::PrepareAsyncDeleteTaskRaw(::grpc::ClientContext* context, const ::flyteidl::service::TaskDeleteRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::TaskDeleteResponse>::Create(channel_.get(), cq, rpcmethod_DeleteTask_, context, request, false);
}

AgentService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      AgentService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< AgentService::Service, ::flyteidl::service::TaskCreateRequest, ::flyteidl::service::TaskCreateResponse>(
          std::mem_fn(&AgentService::Service::CreateTask), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      AgentService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< AgentService::Service, ::flyteidl::service::TaskGetRequest, ::flyteidl::service::TaskGetResponse>(
          std::mem_fn(&AgentService::Service::GetTask), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      AgentService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< AgentService::Service, ::flyteidl::service::TaskDeleteRequest, ::flyteidl::service::TaskDeleteResponse>(
          std::mem_fn(&AgentService::Service::DeleteTask), this)));
}

AgentService::Service::~Service() {
}

::grpc::Status AgentService::Service::CreateTask(::grpc::ServerContext* context, const ::flyteidl::service::TaskCreateRequest* request, ::flyteidl::service::TaskCreateResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status AgentService::Service::GetTask(::grpc::ServerContext* context, const ::flyteidl::service::TaskGetRequest* request, ::flyteidl::service::TaskGetResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status AgentService::Service::DeleteTask(::grpc::ServerContext* context, const ::flyteidl::service::TaskDeleteRequest* request, ::flyteidl::service::TaskDeleteResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace flyteidl
}  // namespace service

