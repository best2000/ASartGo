import bluetooth
import os

def intdivup(n1, n2):
    remin = n1%n2
    if remin != 0:
        return int((n1/n2)+1)
    else:
        return n1/n2

bd_addr = "00:1A:7D:DA:71:11"

port = 5

sock=bluetooth.BluetoothSocket( bluetooth.RFCOMM )
sock.connect((bd_addr, port))


sock.send("client: im going to send bytes")
print(sock.recv(1024).decode('utf-8'))
while True:
    try:
        fname = input("Filepath: ")
        fsize = os.stat(fname).st_size
        break
    except Exception as err:
        print(err)
        print("ctrl+c to exit")

sock.send(fname) #send file name
sock.send(str(fsize)) #send file size
print(sock.recv(1024).decode('utf-8'))
print("File Info")
print(" name:",fname)
print(" size:", str(fsize))
max = intdivup(fsize, 1000)
i = 0
#send byte to server(img)
with open(fname, "rb") as f:
    print("sending bytes...")
    while True:
        percent = int((i/max)*100)
        print(percent, "%", end="\r", flush=True)
        data = f.read(1000)
        if data == b'':
            sock.send("end")
            break

        sock.send(data)
        sock.recv(1024)

        i+=1

#send byte to server(config.json)
with open("config.json", "rb") as f:
    print("sending bytes...")
    while True:
        data = f.read(1000)
        if data == b'':
            sock.send("end")
            break

        sock.send(data)
        sock.recv(1024)

print(sock.recv(1024).decode('utf-8'))

#write recv bytes
print(sock.recv(1024).decode('utf-8'))
print("writing recving file...")
with open("re/"+fname+".html", 'wb') as f:
    while True:
        data = sock.recv(1000)
        if data == b'end':
            break
        f.write(data)
        sock.send("client: ok got that chunk")

sock.close()
print("complete")

