from turtle import *
from enum import Enum
import random
import math

# Screen dimensions
W = 800
H = 480

setup(W, H)

# Calculate tile size using the GCD of W and H
def calculate_tile_size(w, h, recursion_levels):
    gcd = math.gcd(w, h)
    return gcd / (2 ** recursion_levels)

# Determine initial tile size based on recursion depth
RECURSION_LEVELS = 3
TILE_SIZE = calculate_tile_size(W, H, RECURSION_LEVELS)

class TileMode(Enum):
    STRAIGHT = 'straight'
    DIAGONAL = 'diagonal'

def tiling(x, y, w, h, l, mode=TileMode.STRAIGHT): 
    # If we reached the final level of recursion, draw
    if l == 0: 
        penup()

        if mode == TileMode.STRAIGHT: 
            if random.random() < 0.5: 
                # Vertical Line
                goto(x + w / 2, y)
                pendown()
                goto(x - w / 2, y)
            else:
                # Horizontal Line
                goto(x, y + h / 2)
                pendown()
                goto(x, y - h / 2)

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

        tiling(x - w, y + h, w, h, l, mode)
        tiling(x + w, y + h, w, h, l, mode)
        tiling(x - w, y - h, w, h, l, mode)
        tiling(x + w, y - h, w, h, l, mode)


# Set turtle properties
hideturtle()
tracer(False)  # Disable animation for faster drawing

# Start tiling
tiling(0, 0, W, H, RECURSION_LEVELS, mode=TileMode.STRAIGHT)

# Exit on click
exitonclick()

