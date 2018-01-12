import time
import argparse
import grpc
import ping_pb2_grpc as ping
import ping_pb2 as pb

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--addr', help='endpoint')
    args = parser.parse_args()

    channel = grpc.insecure_channel(args.addr)
    sub = ping.PingStub(channel)

    n = 20
    count = 0
    tsum = 0.0
    tmin = 9999.0
    tmax = 0.0

    for i in range(n):
        t1 = time.time()
        res = sub.Ping(pb.PingReq(msg='dsff'))
        t2 = time.time()
        d = (t2 - t1) * 1000
        print(res.msg, '%.3f' % d, 'ms')

        if i > 0:
            count += 1
            tsum += d
            tmin = min(tmin, d)
            tmax = max(tmax, d)

    print('ping x', count, ': avg time=', '%.3f' % (tsum/count), ' min=', '%.3f' % tmin, ' max=', '%.3f' % tmax, 'ms')