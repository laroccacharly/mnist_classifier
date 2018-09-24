from image_encoder import ImageEncoder


def test_decode_encode():
    string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA0klEQVR4nGNgGMyAWUhIqK5jvdSy/9/rUSTkVOJmrfoLAg/X/P102AFZzvDdXyj4HRsUZKGOolHoNljm2LbvH7HYFzAn++/fs9wM2rOwuYaPcdbfKNyO7f67jwmnJPe+v264tSp/fLgghxGXbOCHv3/LJXHJ6u76+3eaNC5Zgdg/f3fjtvjn358O2GX0mrb//Xseq4fUpzwFBuGvbVikJIrugoL3pB+mlLjTVXDIB2KaKbQaHCuHAzgxpMzXPAJJfWnlxmJbB1DmSnuLAHYfUB0AAPtta4Z9bfBAAAAAAElFTkSuQmCC"
    array = ImageEncoder.base64string_to_np(string)
    result_string = ImageEncoder.np_to_base64string(array)
    assert(string == result_string)
