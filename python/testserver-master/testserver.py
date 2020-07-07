import sys
import asynchat
import asyncore
import socket
import time
import datetime

run_mode = sys.argv[1]  # added to run the same module on Master and Slave

if run_mode == 'master':

    task_file_name = sys.argv[2]  # file containing tasks to allocate
    task_file = open(task_file_name, 'r')
    task_list = task_file.read().strip('\n').split('\n')
    task_file.close()

    res_file_name = sys.argv[3]  # result file

    slave_map = {}

    class TestHandler(asynchat.async_chat):
        def __init__(self, sock):
            asynchat.async_chat.__init__(self, sock=sock, map=slave_map)
            self.set_terminator('\n')
            self.buffer = []

        def collect_incoming_data(self, data):
            self.buffer.append(data)

        def found_terminator(self):
            response = ''.join(self.buffer)
            print(response)
            self.buffer = []
            if "Slave ready" in response:
                if hasattr(self, 'push'):
                    if len(task_list) > 0:
                        self.push(task_list.pop(0) + '\n')
                    else:
                        self.close()
            else:
                res_file = open(res_file_name, 'a')
                res_file.write(response + '\n')
                res_file.close()

    class TestMaster(asyncore.dispatcher):
        def __init__(self, hostnm, portnm):
            asyncore.dispatcher.__init__(self, map=slave_map)
            self.create_socket(socket.AF_INET, socket.SOCK_STREAM)
            self.bind((hostnm, portnm))
            self.listen(5)

        def handle_accept(self):
            pair = self.accept()
            if pair is not None:
                sock, address = pair
                print 'Slave found at %s' % repr(address)
                handler = TestHandler(sock)
                self.add_channel(slave_map)

        def handle_close(self):
            self.close()

    host = socket.gethostbyname(socket.gethostname())
    port = 5050
    master = TestMaster(host, port)
    print 'Serving on ', host, ':', port
    asyncore.loop(map=slave_map)

elif run_mode == 'slave':

    master_host = sys.argv[2]
    port = 5050

    class TestSlave(asynchat.async_chat):

        def __init__(self, hostnm, portnm):
            asynchat.async_chat.__init__(self)
            self.terminator = '\n'
            self.buffer = []
            self.create_socket(socket.AF_INET, socket.SOCK_STREAM)
            #self.connect((hostnm, portnm))
            #print(self.connected)
            while not self.connected:
                try:
                    self.connect((hostnm, portnm))
                except socket.error:
                    time.sleep(1)
            if self.connected:
                response_1 = socket.gethostname() + ',' + datetime.datetime.now().strftime('%Y-%m-%d:%H:%M:%S') + ','
                self.push(response_1 + 'Slave ready' + '\n')

        def collect_incoming_data(self, data):
            self.buffer.append(data)

        def found_terminator(self):
            task_assigned = ''.join(self.buffer)
            result, timetaken = self.execute_task(task_assigned)
            response_part_1 = socket.gethostname() + ',' + datetime.datetime.now().strftime('%Y-%m-%d:%H:%M:%S') + ','
            response_part_2 = result + ',' + timetaken
            self.push(response_part_1 + task_assigned + ':' + response_part_2 + '\n')
            self.buffer = []
            self.push(response_part_1 + 'Slave ready' + '\n')

        def handle_connect_event(self):
            self.connected = True

        def execute_task(self, task):
            strtime = datetime.datetime.now()
            print 'Received task:', task
            time.sleep(5)  # added to see behaviour on multiple slaves
            endtime = datetime.datetime.now()
            return ('Pass', str(endtime-strtime))

    slave = TestSlave(master_host, port)
    print 'Waiting on ', socket.gethostbyname(socket.gethostname()), ':', port
    asyncore.loop()
