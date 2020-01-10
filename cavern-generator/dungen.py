# dungen.py
# -------------------------
# Winter 2013; Alex Safatli
# -------------------------
# Software API and interface
# for scalable constructing
# of random dungeon maps.

import random
import cairo as ca

class dir:
    def __init__(self,r):
        self.le = r
        self.arr = self.__rng__()
        self.cursor = -1
    def __rng__(self):
        return [random.randint(-1,1) for x in range(self.le)]
    def get(self):
        if (self.cursor == len(self.arr) - 1):
            self.arr = self.__rng__()
            random.shuffle(self.arr)
            self.cursor = -1
        self.cursor += 1
        return self.arr[self.cursor]

class mapboard:
    def __init__(self,w,h):
        self.width = w
        self.height = h
        self.board = [[0 for i in range(h)] for y in range(w)]
    def get(self,x,y):
        return self.board[x][y]
    def set(self,x,y,i):
        self.board[x][y] = i

class generate:
    def __init__(self,w,h):
        self.width = w
        self.height = h
        self.m = mapboard(w,h)
        self.root = (w/2,h/2)
    def outofbounds(self,x,y):
        w,h = self.width,self.height
        return (x<0 or x>=w or \
                y<0 or y>=h)
    def setwidth(self,x,y,ty):
        m = self.m
        for i in [x,x+1,x-1]:
            for j in [y,y+1,y-1]:
                if not self.outofbounds(i,j):
                    g = m.get(i,j)
                    if g == ty or g == 0 \
                       or g == 2:
                        m.set(i,j,ty)
    def corridors(self):
        # start from root
        x,y = self.root
        m = self.m
        wi = self.setwidth
        oob = self.outofbounds
        wi(x,y,1)
        # travel randomly
        rng = dir(12)
        cx,cy = (x,y)
        while (not oob(cx,cy)):
            wi(cx,cy,1)
            cx += rng.get()
            cy += rng.get()
    def draw(self,fname):
        surf = ca.PDFSurface(fname,self.width,self.height)
        ctx = ca.Context(surf)
        ctx.set_line_width(0.1)
        for x in range(self.width-1):
            for y in range(self.height-1):
                if self.m.get(x,y) == 1:
                    ctx.rectangle(x,y,1,1)
                    ctx.stroke()
        ctx.show_page()

g = generate(200,200)
g.corridors()
g.draw("d.pdf")