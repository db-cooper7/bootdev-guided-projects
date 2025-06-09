from graphics import Window, Line, Point

def main():
    win = Window(800, 600)
    l = Line(Point(69,69), Point(420,420))
    win.draw_line(l)
    win.wait_for_close()


if __name__ == "__main__":
    main()