import bluetooth
import subprocess

server_sock = bluetooth.BluetoothSocket(bluetooth.RFCOMM)

port = 5
server_sock.bind(("",port))
server_sock.listen(1)
print("LISTENING ON PORT:", port)

client_sock,client_info = server_sock.accept()

print("ACCEPTED CONNECTION:",client_info)

print(client_sock.recv(1024).decode('utf-8'))
client_sock.send("server: okay im ready gimme file")
fext = client_sock.recv(1024).decode('utf-8')
print(fext)
client_sock.send("server: ok i got that file info")
with open("in/tempin"+fext, 'wb') as f:
    while True:
        data = client_sock.recv(1000)
        print(data)
        if data == b'end':
            client_sock.send("server: files wrote complete")
            break
        f.write(data)
        print("written")
        client_sock.send(data)

subprocess.run(['pixgen.exe']) #with this it auto display the subprocess stdout to this stdout

server_sock.close()
print("SERVER CLOSED")


