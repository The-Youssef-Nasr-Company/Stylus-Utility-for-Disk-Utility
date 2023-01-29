import pygame
from termcolor import colored

def draw():
	print(colored("ERROR: Cannot find a way to draw anything", "red"))

screen = {
	"draw": draw()
}

screen["draw"]
