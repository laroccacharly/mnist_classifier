import tensorflow as tf
from random import randint
from image_encoder import ImageEncoder


class SamplesGenerator:
    mnist = tf.keras.datasets.mnist

    @classmethod
    def get_samples(cls, number_samples):
        # loading samples
        (x_train, y_train), (_, _) = cls.mnist.load_data()
        samples = []
        for _ in range(number_samples):
            samples.append(cls.make_one_sample(x_train, y_train))

        return samples

    @classmethod
    def make_one_sample(cls, features, labels):
        index = randint(0, len(features))
        encodedImage = ImageEncoder.np_to_base64string(features[index])
        label = int(labels[index])
        return {"encodedImage": encodedImage, "label": label}
