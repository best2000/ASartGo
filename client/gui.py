import tkinter as tk
from tkinter import*
from tkinter import ttk
from PIL import ImageTk, Image
from tkinter import filedialog
import os
import json
import bluetooth

def SelectFile():
    global Img
    global Name
    global Format
    global Size
    global picpath
    picpath = filedialog.askopenfilename(title='open',filetypes = (("jpeg files","*.jpg"),("png files","*.png"),("all files","*.*")))
    Img = Image.open(picpath)
    imsize = str(Img.size[0])+" x "+str(Img.size[1])
    Label(root, text = os.path.basename(picpath)).grid(row = 0, column = 2, sticky = 'w')
    Label(root, text = imsize).grid(row = 1, column = 2, sticky = 'w')
    Label(root, text = Img.format).grid(row = 2, column = 2, sticky = 'w')

def ShowPic():
    Pic = tk.Toplevel()
    img = ImageTk.PhotoImage(Img)
    panel = tk.Label(Pic, image=img)
    panel.image = img
    panel.pack()
    
def Setting():
    with open("config.json",'rb') as cj:
        jbyt = cj.read()
        jstr = jbyt.decode('utf-8')
        jdic = json.loads(jstr)
    global a
    global b
    global c
    global d
    global e
    global f
    a =  StringVar(value=jdic['Tone'][0])
    b =  StringVar(value=jdic['Tone'][1])
    c =  StringVar(value=jdic['Tone'][2])
    d =  StringVar(value=jdic['Tone'][3])
    e =  StringVar(value=jdic['Tone'][4])
    f =  StringVar(value=jdic['ResizeMul'])
    global Set
    Set = tk.Toplevel()
    Set.title('Setting')
    Set.geometry('210x70')
    Label(Set, text='Tone').grid(row=0) 
    Label(Set, text='ResizeMul').grid(row=1) 
    Label(Set, text = '%').grid(row=1, column=6)
    e1 = Entry(Set, width = 2, textvariable = a) 
    e2 = Entry(Set, width = 2, textvariable = b) 
    e3 = Entry(Set, width = 2, textvariable = c) 
    e4 = Entry(Set, width = 2, textvariable = d)
    e5 = Entry(Set, width = 2, textvariable = e)
    e6 = Entry(Set, width = 21, textvariable = f)
    b1 = Button(Set, text = 'Confirm', command = Confirm)
    e1.grid(row=0, column=1) 
    e2.grid(row=0, column=2)
    e3.grid(row=0, column=3) 
    e4.grid(row=0, column=4)
    e5.grid(row=0, column=5)
    e6.grid(row=1, column=1, columnspan = 5)
    b1.grid(row = 3, column = 0, columnspan = 6)

def Confirm():
    tone = [a.get(), b.get(), c.get(), d.get(), e.get()]
    jdic = {"Tone":["", "", "", "", ""],"ResizeMul":""}
    for i in range(5):
        jdic["Tone"][i] = tone[i] 
    jdic["ResizeMul"] = f.get()
    jstr = json.dumps(jdic)
    with open("config.json",'wb') as cj:
        jbyt = bytes(jstr, encoding='utf-8')
        cj.write(jbyt)
    Set.destroy()
    
def Convert():
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

    fname = os.path.basename(picpath)
    fsize = os.stat(fname).st_size
    sock.send(fname) #send file name
    sock.send(str(fsize)) #send file size
    print(sock.recv(1024).decode('utf-8'))
    print("File Info")
    print(" name:",fname)
    print(" size:", str(fsize))
    max = intdivup(fsize, 1000)
    i = 0
    #send byte to server(img.xxx)
    with open(picpath, "rb") as f:
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

Name = None
Size = None
Format = None
root = tk.Tk()
root.title('ASCII')
root.geometry('400x104')
Button(root, text="SELECT FILE", width=25,command = SelectFile).grid(row=0)
Button(root, text="SHOW PICTURE", width=25,command = ShowPic).grid(row=1)
Button(root, text="SETTING", width=25,command = Setting).grid(row=2)
Button(root, text="CONVERT", width=25).grid(row=3)
Label(root, text = 'Name = ').grid(row = 0, column=1)
Label(root, text = 'Size = ').grid(row = 1, column = 1)
Label(root, text = 'Format = ').grid(row = 2, column = 1)

root.mainloop()