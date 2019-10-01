import tkinter as tk
from tkinter import*
from tkinter import ttk
from PIL import ImageTk, Image
from tkinter import filedialog
import os
import json
import subprocess

def SelectFile():
    global Img
    global Name
    global Format
    global Size
    global picpath
    try:
        picpath = filedialog.askopenfilename(title='open',filetypes = (("jpeg files","*.jpg"),("png files","*.png"),("all files","*.*")))
        Img = Image.open(picpath)
        imsize = str(Img.size[0])+" x "+str(Img.size[1])
        Label(root, text = os.path.basename(picpath)).grid(row = 0, column = 2, sticky = 'w')
        Label(root, text = imsize).grid(row = 1, column = 2, sticky = 'w')
        Label(root, text = Img.format).grid(row = 2, column = 2, sticky = 'w')
    except Exception:
        pass
    
    pathbyt = bytes(picpath, encoding='utf-8')
    with open("picpath",'wb') as picpath:
        picpath.write(pathbyt)

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
    global g
    a =  StringVar(value=jdic['Tone'][0])
    b =  StringVar(value=jdic['Tone'][1])
    c =  StringVar(value=jdic['Tone'][2])
    d =  StringVar(value=jdic['Tone'][3])
    e =  StringVar(value=jdic['Tone'][4])
    f =  StringVar(value=jdic['ResizeMul'])
    g =  StringVar(value=jdic['SavePath'])
    global Set
    Set = tk.Toplevel()
    Set.title('Setting')
    Set.geometry('210x90')
    Label(Set, text='Tone').grid(row=0) 
    Label(Set, text='ResizeMul').grid(row=1) 
    Label(Set, text = '%').grid(row=1, column=6)
    Label(Set, text='SavePath').grid(row=2) 
    e1 = Entry(Set, width = 2, textvariable = a) 
    e2 = Entry(Set, width = 2, textvariable = b) 
    e3 = Entry(Set, width = 2, textvariable = c) 
    e4 = Entry(Set, width = 2, textvariable = d)
    e5 = Entry(Set, width = 2, textvariable = e)
    e6 = Entry(Set, width = 21, textvariable = f)
    e7 = Entry(Set, width = 21, textvariable = g)
    b1 = Button(Set, text = 'Confirm', command = Confirm)
    e1.grid(row=0, column=1) 
    e2.grid(row=0, column=2)
    e3.grid(row=0, column=3) 
    e4.grid(row=0, column=4)
    e5.grid(row=0, column=5)
    e6.grid(row=1, column=1, columnspan = 5)
    e7.grid(row=2, column=1, columnspan = 5)
    b1.grid(row = 3, column = 0, columnspan = 6)

def Confirm():
    tone = [a.get(), b.get(), c.get(), d.get(), e.get()]
    jdic = {"Tone":["", "", "", "", ""], "ResizeMul":"", "SavePath":""}
    for i in range(5):
        jdic["Tone"][i] = tone[i] 
    jdic["ResizeMul"] = f.get()
    jdic["SavePath"] = g.get()
    jstr = json.dumps(jdic)
    with open("config.json",'wb') as cj:
        jbyt = bytes(jstr, encoding='utf-8')
        cj.write(jbyt)
    Set.destroy()

def Convert():
    with open("proc",'rb') as f:
        stat = f.read()
        statstr = stat.decode('utf-8')

    if statstr == "0":
        p = subprocess.Popen(['pixgen.exe'])
    

def ShowOut():
    subprocess.run(['openout.bat'])

with open("proc",'wb') as f:
    f.write(b'0')
Name = None
Size = None
Format = None
root = tk.Tk()
root.title('ASCII')
root.geometry('400x130')
Button(root, text="Select File", width=25, command = SelectFile).grid(row=0)
Button(root, text="Show Pic", width=25, command = ShowPic).grid(row=1)
Button(root, text="Setting", width=25, command = Setting).grid(row=2)
cbut = Button(root, text="Convert", width=25, command = Convert, state=NORMAL)
cbut.grid(row=3)
Button(root, text="Show Output Folder", width=25, command = ShowOut).grid(row=4)
Label(root, text = 'Name = ').grid(row = 0, column=1)
Label(root, text = 'Size  = ').grid(row = 1, column = 1)
Label(root, text = 'Format = ').grid(row = 2, column = 1)

root.mainloop()
