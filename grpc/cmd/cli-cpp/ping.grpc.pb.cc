// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: ping.proto

#include "ping.pb.h"
#include "ping.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace pb {

static const char* Ping_method_names[] = {
  "/pb.Ping/Ping",
};

std::unique_ptr< Ping::Stub> Ping::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< Ping::Stub> stub(new Ping::Stub(channel));
  return stub;
}

Ping::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_Ping_(Ping_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Ping::Stub::Ping(::grpc::ClientContext* context, const ::pb::PingReq& request, ::pb::PingRes* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Ping_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::pb::PingRes>* Ping::Stub::AsyncPingRaw(::grpc::ClientContext* context, const ::pb::PingReq& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::pb::PingRes>::Create(channel_.get(), cq, rpcmethod_Ping_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::pb::PingRes>* Ping::Stub::PrepareAsyncPingRaw(::grpc::ClientContext* context, const ::pb::PingReq& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::pb::PingRes>::Create(channel_.get(), cq, rpcmethod_Ping_, context, request, false);
}

Ping::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Ping_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Ping::Service, ::pb::PingReq, ::pb::PingRes>(
          std::mem_fn(&Ping::Service::Ping), this)));
}

Ping::Service::~Service() {
}

::grpc::Status Ping::Service::Ping(::grpc::ServerContext* context, const ::pb::PingReq* request, ::pb::PingRes* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace pb

