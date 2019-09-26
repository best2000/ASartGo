import tkinter as tk
from tkinter import*
from tkinter import ttk
from PIL import ImageTk, Image
from tkinter import filedialog
import os

def Openfile():
    global Img
    global Name
    global Format
    global Size
    global Bits
    picfile = filedialog.askopenfilename(title='open',filetypes = (("jpeg files","*.jpg"),("png files","*.png"),("all files","*.*")))
    Img = Image.open(picfile)
    Label(root, text = os.path.basename(picfile)).grid(row = 0, column = 2, sticky = 'w')
    Label(root, text = Img.size).grid(row = 1, column = 2, sticky = 'w')
    Label(root, text = Img.format).grid(row = 2, column = 2, sticky = 'w')
    print(picfile)

def ShowPic():
    Pic = tk.Toplevel()
    img = ImageTk.PhotoImage(Img)
    panel = tk.Label(Pic, image=img)
    panel.image = img
    panel.pack()
    
def Setting():
    global a
    global b
    global c
    global d
    global e
    global f
    a =  StringVar()
    b =  StringVar()
    c =  StringVar()
    d =  StringVar()
    e =  StringVar()
    f =  DoubleVar()
    Set = tk.Toplevel()
    Set.title('Setting')
    Set.geometry('210x70')
    Label(Set, text='First Name').grid(row=0) 
    Label(Set, text='Last Name').grid(row=1) 
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
    t1 = a.get()
    t2 = b.get()
    t3 = c.get()
    t4 = d.get()
    t5 = e.get()
    resize = f.get()
    print(t1,t2,t3,t4,t5,resize)

Name = None
Size = None
Format = None
root = tk.Tk()
root.title('ASCII')
root.geometry('400x104')
Button(root, text="SELECT FILE", width=25,command = Openfile).grid(row=0)
Button(root, text="SHOW PICTURE", width=25,command = ShowPic).grid(row=1)
Button(root, text="SETTING", width=25,command = Setting).grid(row=2)
Button(root, text="CONVERT", width=25).grid(row=3)
Label(root, text = 'Name = ').grid(row = 0, column=1)
Label(root, text = 'Size = ').grid(row = 1, column = 1)
Label(root, text = 'Format = ').grid(row = 2, column = 1)

root.mainloop()