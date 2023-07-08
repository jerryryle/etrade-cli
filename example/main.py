import queue
import os
import requests
import signal
import subprocess
import sys
import threading
import time


class ETradeServer:
    def __init__(self, command):
        self._command = command
        self._process = None
        self._output_queue = queue.Queue()
        self._reader_thread = None

    def _reader_thread(self):
        # Read each line until EOF (b'')
        for line in iter(self._process.stdout.readline, b''):
            self._output_queue.put(line.decode().strip())

    def start(self):
        # Open process, combining stdout and stderr for capture via pipe
        self._process = subprocess.Popen(self._command, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        # Start thread to read any command output
        self._reader_thread = threading.Thread(target=self._reader_thread)
        self._reader_thread.start()

    def stop(self):
        if self._process:
            # Send sigint to tell the server to shut itself down
            self._process.send_signal(signal.SIGINT)
            self._process.wait()
            self._process = None
        if self._reader_thread:
            self._reader_thread.join()
            self._reader_thread = None

    def is_running(self):
        return self._process is not None and self._process.poll() is None

    def get_output(self):
        while not self._output_queue.empty():
            yield self._output_queue.get()


if __name__ == "__main__":
    customerId = os.getenv("ETRADE_CUSTOMER")
    if customerId is None:
        print("error: you must specify a customer ID with the ETRADE_CUSTOMER environment variable", file=sys.stderr)
    server = ETradeServer(["etrade", "server", "--addr", ":8888"])
    server.start()
    time.sleep(1)

    try:
        r = requests.post(f"http://127.0.0.1:8888/customers/{customerId}/auth")
        r.raise_for_status()
        response = r.json()
        if response["status"] == "authorize":
            print(f"visit this url to get a verification code: {response['authorizationUrl']}")
            authCode = input("Enter verification code: ")
            r = requests.post(f"http://127.0.0.1:8888/customers/{customerId}/auth", data={'verifyCode': authCode})
            r.raise_for_status()

        r = requests.get(f"http://127.0.0.1:8888/customers/{customerId}/accounts")
        r.raise_for_status()
        print(r.text)

    finally:
        server.stop()

