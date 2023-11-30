from input import format_input
import sys


def get_total_fuel(modules):
    total = 0
    for module in modules:
        print("module mass:", module)
        get_module_fuel(module)
        total += get_module_fuel(module)
    return total


def get_module_fuel(mass):
    module_fuel = 0
    fuel_needed = get_fuel(mass)
    while fuel_needed > 0:
        module_fuel += fuel_needed
        fuel_needed = get_fuel(fuel_needed)
    return module_fuel

def get_fuel(mass):
    return int(mass / 3) - 2

def run():
    my_input = format_input("day1")
    print("input size:", len(my_input))
    print(get_total_fuel(my_input))
