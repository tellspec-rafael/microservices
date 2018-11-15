#Worker based on http://zguide.zeromq.org/py:mtserver

import time
import threading
import zmq
import sys
import argparse
import uuid

def worker(worker_url, context=None):
    """Worker routine"""
    context = context or zmq.Context.instance()
    # Socket to talk to dispatcher
    socket = context.socket(zmq.REP)
    socket.connect(worker_url)

    while True:

        string  = socket.recv()

        print("Received request: [ %s ]" % (string))

        # do some 'work'
        time.sleep(1)

        #send reply back to client
        socket.send(string)

def main(argv):

	parser = argparse.ArgumentParser(description='Arguments of python worker')
	parser.add_argument ("--threads", type=int, default=1,   help='Number of threads for this worker.') 
	args = parser.parse_args()

	workerInternalIdentifier =  str(uuid.uuid4())

	worker_url = "ipc:///tmp/workers" + workerInternalIdentifier + ".ipc"

	# Start threads
	for i in range(args.threads):
		thread = threading.Thread(target=worker, args=(worker_url,))
		thread.start()
		print("Thread %d created" % (i))

	context = zmq.Context()
	frontend = context.socket(zmq.ROUTER)
	frontend.connect("tcp://router:5558")

	backend = context.socket(zmq.DEALER)
	backend.bind(worker_url)

	zmq.proxy(frontend, backend)

	frontend.close()
	backend.close()
	context.term()


if __name__ == "__main__":
    main(sys.argv)

