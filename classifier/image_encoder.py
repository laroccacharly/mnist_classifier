import base64
import numpy as np
import png
import os

class ImageEncoder:
    image_name = "image.png"
    width = 28
    height = 28

    @classmethod
    def base64string_to_np(cls, string):
        encodedbase64 = string.split(",")[-1].encode()  # get the string
        decoded = base64.decodebytes(encodedbase64)  # Decode it
        f = open(cls.image_name, 'wb')  # make a png from it
        f.write(decoded)
        f.close()

        map = png.Reader(filename=cls.image_name).read()[2]  # read the file and get the iterator
        arr = np.array([], dtype=np.uint8)
        for row in map:
            arr = np.append(arr, np.array(row))  # Build the np array

        os.remove(cls.image_name) # clean up
        return arr.reshape([cls.height, cls.width])

    @classmethod
    def np_to_base64string(cls, array):
        f = open(cls.image_name, 'wb')
        w = png.Writer(cls.height, cls.width, greyscale=True)
        w.write(f, array) # Write the array to a png
        f.close()

        # Open the created png and convert its content to base64
        f = open(cls.image_name, "rb")
        encoded = base64.b64encode(f.read()).decode()
        f.close()
        os.remove(cls.image_name) # clean up
        return 'data:image/png;base64,{}'.format(encoded)
