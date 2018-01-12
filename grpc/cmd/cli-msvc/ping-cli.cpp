// ping-cli.cpp : Defines the entry point for the console application.
//

#include <SDKDDKVer.h>

#include <stdio.h>
#include <tchar.h>

#include <grpc++/grpc++.h>
#include "../cli-cpp/ping.grpc.pb.h"
#include <iostream>
#include <chrono>

using namespace std;

int main() {
	auto addr = "192.168.99.100:8889";
	//auto addr = "localhost:8888";

	auto stub = pb::Ping::NewStub(
		grpc::CreateChannel(addr, grpc::InsecureChannelCredentials())
	);


	const int N = 20;
	double tsum = 0;
	double tmin = 9999;
	double tmax = 0;
	int count = 0;

	for (int i = 0; i < N; i++) {
		grpc::ClientContext ctx;
		pb::PingReq req;
		req.set_msg("abd");

		pb::PingRes res;

		auto t1 = chrono::system_clock::now();

		auto status = stub->Ping(&ctx, req, &res);

		auto t2 = chrono::system_clock::now();

		if (!status.ok()) {
			cerr << "failed " << status.error_message() << endl;
			return -1;
		}

		double d = chrono::duration_cast<chrono::microseconds>(t2 - t1).count() / 1000.0;

		cout << "Reply: " << res.msg() << " (" << d << " ms)" << endl;

		if (i > 0) {
			count++;
			tsum += d;
			tmin = min(tmin, d);
			tmax = max(tmax, d);
		}
	}

	cout << "ping x " << count << " : avg time=" << (tsum / count) << " min=" << tmin << " max=" << tmax << " ms" << endl;

	return 0;
}

#pragma comment(lib, "ws2_32.lib")
#pragma comment(lib, "gpr.lib")
#pragma comment(lib, "grpc.lib")
#pragma comment(lib, "grpc++.lib")
#pragma comment(lib, "ssleay32.lib")
#pragma comment(lib, "libeay32.lib")

#if defined(DEBUG) || defined(_DEBUG)
#pragma comment(lib, "libprotobufd.lib")
#pragma comment(lib, "libprotocd.lib")
#pragma comment(lib, "zlibd.lib")
#else
#pragma comment(lib, "libprotobuf.lib")
#pragma comment(lib, "libprotoc.lib")
#pragma comment(lib, "zlib.lib")
#endif

#pragma comment(lib, "User32.lib")
#pragma comment(lib, "GDI32.lib")
#pragma comment(lib, "Advapi32.lib")

