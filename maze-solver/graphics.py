from tkinter import Tk, BOTH, Canvas

class Point:
    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y

class Line:
    def __init__(self, p1: Point, p2: Point) -> None:
        self.p1 = p1
        self.p2 = p2
    
    def draw(self, canvas: Canvas, fill_color: str = "black") -> None:
        canvas.create_line(
            self.p1.x, self.p1.y, self.p2.x, self.p2.y, fill=fill_color, width=2
        )

class Window:
    def __init__(self, width: int = 800, height: int = 600) -> None:
        self.__root = Tk()
        self.__root.title("maze-solver")
        self.__canvas = Canvas(self.__root, bg="white", height=height, width=width)
        self.__canvas.pack(fill=BOTH, expand=1)
        self.__running = False
        self.__root.protocol("WM_DELETE_WINDOW", self.close)

    def redraw(self) -> None:
        self.__root.update_idletasks()
        self.__root.update()

    def wait_for_close(self) -> None:
        self.__running = True
        while self.__running:
            self.redraw()
    
    def close(self) -> None:
        self.__running = False

    def draw_line(self, line: Line, fill_color: str = "black") -> None:
        line.draw(self.__canvas, fill_color)