class Rectangle:
    def __init__(self, width:float=1.0, height:float=1.0):
        self._width = width
        self._height = height

    def get_width(self) -> float:
        return self._width
    
    def get_height(self) -> float:
        return self._height
    
    def get_area(self) -> float:
        return self._width * self._height
    
    def get_perimeter(self) -> float:
        return (2*self._height + 2*self._width)

    def print_rectangle(self, name:str) -> None:
        print(f"Rectangle {name} is {self._width} units wide and {self._height} units high.")


if __name__ == "__main__":
    # Input width and height from the user
    width = float(input("Enter the width of the rectangle (decimal): "))
    height = float(input("Enter the height of the rectangle (decimal): "))

    my_rectangle = Rectangle()
    your_rectangle = Rectangle(width=width, height=height)

    my_rectangle.print_rectangle("my_rectangle")
    your_rectangle.print_rectangle("your_rectangle")