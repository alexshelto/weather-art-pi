from enum import Enum
import random
import time
from turtle import *

# Screen dimensions
W = 800
H = 480
SLEEP_TIME = 90

TILE_MIN_LEVEL = 3
TILE_MAX_LEVEL = 6

setup(W, H)

class TileMode(Enum):
    STRAIGHT = 'straight'
    DIAGONAL = 'diagonal'

def tiling(x, y, w, h, l, mode=TileMode.STRAIGHT): 
    # If we reached the final level of recursion, draw
    if l == 0: 
        penup()

        if mode == TileMode.STRAIGHT: 
            if random.random() < 0.5: 
                # Vertical Line (spans full tile height)
                goto(x, y + h / 2)
                pendown()
                goto(x, y - h / 2)
            else:
                # Horizontal Line (spans full tile width)
                goto(x - w / 2, y)
                pendown()
                goto(x + w / 2, y)

        elif mode == TileMode.DIAGONAL: 
            if random.random() < 0.5: 
                # Top-left to bottom-right
                goto(x - w / 2, y + h / 2)
                pendown()
                goto(x + w / 2, y - h / 2)
            else:
                # Bottom-left to top-right
                goto(x - w / 2, y - h / 2)
                pendown()
                goto(x + w / 2, y + h / 2)

        penup()
    else: 
        # Reduce tile size and go deeper into recursion
        w /= 2
        h /= 2
        l -= 1

        # Recursively tile each quadrant
        tiling(x - w / 2, y + h / 2, w, h, l, mode)
        tiling(x + w / 2, y + h / 2, w, h, l, mode)
        tiling(x - w / 2, y - h / 2, w, h, l, mode)
        tiling(x + w / 2, y - h / 2, w, h, l, mode)


def drawTiling(mode): 
        clear()

        tiling(0, 0, W, H, random.randint(TILE_MIN_LEVEL, TILE_MAX_LEVEL), mode)
        time.sleep(SLEEP_TIME)

        clear()




def draw_loop():
    while True: 
        drawTiling(TileMode.STRAIGHT)
        drawTiling(TileMode.DIAGONAL)
        

# Set turtle properties
hideturtle()

speed(1)
delay(20)

draw_loop()


exitonclick()

