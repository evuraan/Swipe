#!/usr/bin/python3 -u

# /* Copyright (C) 2022  Evuraan, <evuraan@gmail.com> */
import threading
import sys
import os
import time
from datetime import datetime
import warnings
import gi
gi.require_versions({'Gtk': '3.0','XApp': '1.0'})
from gi.repository import Gtk, XApp
from gi.repository.GdkPixbuf import Pixbuf
from queue import Queue


warnings.filterwarnings("ignore")

NAME = "Swipe"
DESC =  NAME + " Linux Gestures"
VERSION = "6.0d"
WEBSITE = "https://github.com/evuraan/Swipe"

class SwipeIcon:
    def __init__(self):
        self.state = False 
        self.left_Menu()
        self.rightMenu()
        self.status_icon = XApp.StatusIcon()
        self.status_icon.set_icon_name(ICON)
        self.status_icon.props.icon_size = 16
        self.status_icon.set_primary_menu (self.left_menu)
        self.status_icon.set_secondary_menu (self.right_menu)
        self.status_icon.set_tooltip_text(DESC)
        self.state = True 
        self.q = Queue()
        t3 = threading.Thread(target=self.icon_changer_fifo)
        t3.start()
        t4 = threading.Thread(target=self.blink)
        t4.start()


    def icon_changer_fifo(self):
        if len(sys.argv) >= 5:
            fifo = sys.argv[4]
            try:
                with open(fifo) as fi:
                    while True:
                        line = fi.readline()
                        self.q.put(1)
            except Exception as e:
                print("pipe err", e)
                os._exit(1)

    def blink(self):
        while True:
            x = self.q.get()
            self.status_icon.set_icon_name(sys.argv[3])
            time.sleep(0.3)
            self.status_icon.set_icon_name(ICON)
            self.q.task_done()

    def left_Menu(self):
        self.left_menu = Gtk.Menu()

        about = Gtk.ImageMenuItem(label="About", image=Gtk.Image.new_from_icon_name("help-about", 16))
        about.connect("activate", self.aboutDialog)
        self.left_menu.append(about)

        quit = Gtk.ImageMenuItem(label="Quit", image=Gtk.Image.new_from_icon_name("application-exit", 16))
        quit.connect("activate", self.quitter)
        self.left_menu.append(quit)
        self.left_menu.show_all()

    def rightMenu(self):
        self.right_menu = Gtk.Menu()

        about = Gtk.ImageMenuItem(label="About", image=Gtk.Image.new_from_icon_name("help-about", 16))
        about.connect("activate", self.aboutDialog)
        self.right_menu.append(about)

        self.right_menu.show_all()

    def getState(self):
        return self.state

    def quitter(self, x):
        try:
            Gtk.main_quit()
            os.remove(sys.argv[0])
            os.remove(sys.argv[1])
            os.remove(sys.argv[3])
            os.remove(sys.argv[4])
        except:
            pass
        os._exit(0)

    def aboutDialog(self, widget):
        about_dialog = Gtk.AboutDialog()

        about_dialog.set_destroy_with_parent(True)
        about_dialog.set_name(DESC)
        about_dialog.set_program_name(NAME)
        about_dialog.set_comments(NAME)
        about_dialog.set_version(VERSION)
        about_dialog.set_copyright("Copyright © 2021 - {} evuraan".format(datetime.now().year))
        about_dialog.set_authors(["evuraan"])
        about_dialog.set_website(WEBSITE)
        about_dialog.set_website_label(WEBSITE)
        about_dialog.set_logo(PIXBUF)

        about_dialog.run()
        about_dialog.destroy()

def if_caller_gone(pid):
    pidFo = "/proc/" + pid
    while True:
        if not os.path.isdir(pidFo):
            os._exit(1)
        time.sleep(1.1)

def cleanup():
    while True:
        if app.getState():
            break
        time.sleep(3)
    try:
        #os.remove(sys.argv[1])
        os.remove(sys.argv[0])
    except:
        pass 

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Invalid usage")
        sys.exit(1)
    ICON = sys.argv[1]
    PIXBUF = Pixbuf.new_from_file(sys.argv[1])
    app = SwipeIcon()
    t1 = threading.Thread(target=if_caller_gone, args=(sys.argv[2],))
    t1.start()
    t2 = threading.Thread(target=cleanup)
    t2.start()
    Gtk.main()
