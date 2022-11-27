#!/usr/bin/python3 -u

# /* Copyright (C) 2022  Evuraan, <evuraan@gmail.com> */
import sys
import warnings
import gi
gi.require_versions({'Gtk': '3.0','XApp': '1.0'})

from gi.repository import Gtk, XApp
from gi.repository.GdkPixbuf import Pixbuf

warnings.filterwarnings("ignore")

NAME = "Swipe"
desc =  NAME + " Linux Gestures"
VERSION = "5.0"
WEBSITE = "https://github.com/evuraan/Swipe"

class SwipeIcon:
    def __init__(self):
        self.left_Menu()
        self.rightMenu()
        self.status_icon = XApp.StatusIcon()
        self.status_icon.set_icon_name(ICON)
        self.status_icon.props.icon_size = 16
        self.status_icon.set_primary_menu (self.left_menu)
        self.status_icon.set_secondary_menu (self.right_menu)
        self.status_icon.set_tooltip_text(desc)
        
    def left_Menu(self):
        self.left_menu = Gtk.Menu()

        about = Gtk.ImageMenuItem(label="About", image=Gtk.Image.new_from_icon_name("help-about", 16))
        about.connect("activate", self.aboutDialog)
        self.left_menu.append(about)

        quit = Gtk.ImageMenuItem(label="Quit", image=Gtk.Image.new_from_icon_name("application-exit", 16))
        quit.connect("activate", Gtk.main_quit)
        self.left_menu.append(quit)

        self.left_menu.show_all()
        
    def rightMenu(self):
        self.right_menu = Gtk.Menu()

        about = Gtk.ImageMenuItem(label="About", image=Gtk.Image.new_from_icon_name("help-about", 16))
        about.connect("activate", self.aboutDialog)
        self.right_menu.append(about)

        self.right_menu.show_all()


    def aboutDialog(self, widget):
        about_dialog = Gtk.AboutDialog()

        about_dialog.set_destroy_with_parent(True)
        about_dialog.set_name(desc)
        about_dialog.set_program_name(NAME)
        about_dialog.set_comments(NAME)
        about_dialog.set_version(VERSION)
        about_dialog.set_copyright("Copyright © 2021 evuraan")
        about_dialog.set_authors(["evuraan"])
        about_dialog.set_website(WEBSITE)
        about_dialog.set_website_label(WEBSITE)
        about_dialog.set_logo(PIXBUF)

        about_dialog.run()
        about_dialog.destroy()

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Invalid usage")
        sys.exit(1)
    ICON = sys.argv[1]
    PIXBUF = Pixbuf.new_from_file(sys.argv[1])
    app = SwipeIcon()
    Gtk.main()