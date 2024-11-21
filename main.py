from turtle import *
from enum import Enum
import random 

W = 800
H = 480

setup(W, H)


class TileMode(Enum):
    STRAIGHT = 'straight'
    DIAGONAL = 'diagonal'


def tiling(x, y, s, l, mode=TileMode.STRAIGHT): 
    # if we reached final level of recursion draw
    if l == 0: 
        penup()

        if mode == TileMode.STRAIGHT: 
            if random.random() < 0.5: 
                # Vertical Line 
                goto(x,y-s)
                pendown()
                goto(x,y+s)
            else:
                # HZ Line 
                goto(x-s,y)
                pendown()
                goto(x+s,y)

        elif mode == TileMode.DIAGONAL: 
            # Top left to bottom right
            if random.random() < 0.5: 
                # Vertical Line 
                goto(x-s,y+s)
                pendown()
                goto(x+s,y-s)
            # Bottom left to top right 
            else:
                # HZ Line 
                goto(x-s,y-s)
                pendown()
                goto(x+s,y+s)


        penup()
    # Split screen and go down another level of recursion
    else: 
        s /= 2
        l -= 1

        tiling(x-s,y+s,s,l, mode)
        tiling(x+s,y+s,s,l, mode)
        tiling(x-s,y-s,s,l, mode)
        tiling(x+s,y-s,s,l, mode)

hideturtle()

#speed(0.5)
tracer(False) #Disable when done dev

tiling(0,0,240,3, mode=TileMode.STRAIGHT)

exitonclick()
