import bluetooth
import subprocess
import os

#connection setup
server_sock = bluetooth.BluetoothSocket(bluetooth.RFCOMM)
port = 5
server_sock.bind(("",port))
server_sock.listen(1)
print("LISTENING ON PORT:", port)

#main loop
while True:
    #wait for client
    print("Waiting for client...")
    client_sock,client_info = server_sock.accept()
    print("ACCEPTED CONNECTION:",client_info)

    #server check
    if client_sock.recv(1024) == b'server0':
        client_sock.send("server closed")
        break
    client_sock.send("server ready!")

    #file name info recv
    fname = client_sock.recv(1024).decode('utf-8') #xxx.jpg
    client_sock.send("server: ok i got that file info")

    #write recv bytes of img
    print("Writing recv bytes...(img.xxx)")
    with open("in/"+fname, 'wb') as f:
        while True:
            data = client_sock.recv(1000)
            if data == b'end':
                break
            f.write(data)
            client_sock.send("server: ok got that chunk")

    #write recv bytes of config.json
    print("Writing recv bytes...(config.json)")
    with open("conin/config.json", 'wb') as f:
        while True:
            data = client_sock.recv(1000)
            if data == b'end':
                break
            f.write(data)
            client_sock.send("server: ok got that chunk")

    #convert image --------------#
    client_sock.send("converting...")
    subprocess.run(['go','run','pixgen.go']) #with this it auto display the subprocess stdout to this stdout
    client_sock.send("finish converting")
    #----------------------------#

    #send file back
    print("Sending converted file back...")
    with open("out/"+fname+".html", 'rb') as f:
        while True:
            data = f.read(1000)
            if data == b'':
                client_sock.send("end")
                break
            client_sock.send(data)
            client_sock.recv(1024)
    #delete sended file
    os.remove("out/"+fname+".html")
    #-----------------------------#

    #finsih this cliend so close connection
    print("Complete this client...")
    client_sock.close()

#close server
server_sock.close()
print("SERVER CLOSED")


