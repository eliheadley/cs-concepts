from abc import ABC, abstractmethod

class Shape(ABC):

    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def perimeter(self):
        pass


class Circle(Shape):
    PI = 3.14159

    def __init__(self, radius = 1.0):
        self._radius = radius

    def area(self):
        return Circle.PI * self._radius * self._radius
    
    def perimeter(self):
        return Circle.PI * 2 * self._radius
    

class Rectangle(Shape):

    def __init__(self, width:float=1.0, height:float=1.0):
        self._width = width
        self._height = height

    def area(self) -> float:
        return self._width * self._height
    
    def perimeter(self) -> float:
        return (2*self._height + 2*self._width)

# Polymorphic function
def print_shape_info(shape:Shape):
    """Function that works polymorphically with any Shape object."""
    print(f"Shape Area: {shape.area()}")
    print(f"Shape Perimeter: {shape.perimeter()}")

if __name__ == "__main__":
    circle = Circle(radius=5)
    rectangle = Rectangle(width=4, height=6)

    print_shape_info(circle)
    print_shape_info(rectangle)