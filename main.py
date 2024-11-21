from enum import Enum
import random
import math
import time
from turtle import *

# Screen dimensions
W = 800
H = 480
SLEEP_TIME = 90

TILE_MIN_LEVEL = 3
TILE_MAX_LEVEL = 6

colors = ['#B1FF00','#642dd2', 'green', 'yellow', '#C73838']

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
    new_color = random.choice(colors)
    color(new_color)

    recursion_level = random.randint(TILE_MIN_LEVEL, TILE_MAX_LEVEL)
    if recursion_level > 4: 
        speed(2)
    else: 
        speed(1)
        delay(20)

    tiling(0, 0, W, H, recursion_level, mode)
    time.sleep(SLEEP_TIME)


def draw_recursive_shape(sides):
    lens = {3: 200, 4: 170, 5: 130, 6: 110}
    speed(2)
    delay(0)

    new_color = random.choice(colors)
    color(new_color)

    rotation_angle = random.randint(5, 180)

    side_len = lens[sides]

    shape_draw(side_len, 100, sides, rotation_angle)

    time.sleep(SLEEP_TIME)

def shape_draw(len, count, sides, rotation_angle):
    turn_angle = 360 / sides

    for _ in range(count): 
        right(rotation_angle)
        for _ in range(sides):
            forward(len)
            right(turn_angle)

    
def draw_loop():
    while True: 
        clear()

        choice = random.randint(0,1)

        if choice == 0:
            straight_or_diag = random.randint(0,1)
            if straight_or_diag == 0: 
                drawTiling(TileMode.STRAIGHT)
            else:
                drawTiling(TileMode.DIAGONAL)

        elif choice == 1: 
            num_sides = random.randint(3,6)
            draw_recursive_shape(num_sides)
        
        else:
            exit('not here yet')

        time.sleep(SLEEP_TIME)




# Set turtle properties
random.seed()
hideturtle()
bgcolor('black')

draw_loop()


exitonclick()

