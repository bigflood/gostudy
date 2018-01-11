
#include <grpc++/grpc++.h>
#include "ping.grpc.pb.h"
#include <iostream>
#include <chrono>

using namespace std;

int main() {
    auto stub = pb::Ping::NewStub(
        grpc::CreateChannel("localhost:8082", grpc::InsecureChannelCredentials())
    );

    
    const int N = 20;
    double tsum = 0;
    double tmin = 9999;
    double tmax = 0;
    int count = 0;

    for (int i = 0 ; i < N ; i++) {
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

    cout << "ping x " << count << " : avg time=" << (tsum/count) << " min=" << tmin << " max=" << tmax << " ms" << endl;

    return 0;
}